package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {

	target, err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		originalPath := r.URL.Path
		fmt.Println("Original Path:", originalPath)

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)
		fmt.Println("Stripped Path:", strippedPath)

		r.URL.Host = target.Host
		r.URL.Path = target.Path + strippedPath // here target.Path is usually empty unless targetBaseUrl has some path

		r.Host = target.Host

		if userId, ok := r.Context().Value("userID").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}

	}

	return proxy.ServeHTTP

}
