package controllers

import "net/http"

// here w is the http.ResponseWriter and r is the *http.Request
// here r is returning as reference type because if we create a new request each time it will take a lot of memory
// here w is not pass a reference because it's a Responwriter object and we each time write a different response, supppose if we have 5 endpoints and for all response we are using the same http.ResponseWriter object then it will be a problem because we are writing to the same object and it will overwrite the previous response.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// if content type is not set in the header the Write method will set the content type based on the initial 512 bytes of data that you are going to wirte.
	w.Write([]byte("pong"))
}
