
# Setting up : API Gateway in Golang

To start the project first run
1. go mod init <moduleName> or go mod init AuthInGo

* To setup dotenv liberary we will use joho/godotenv for this run below command
```go
go get github.com/joho/godotenv
```

* Here go.sum is similir to package.lock.json file in then node js project which maintin the exact versions of all installed packages, including their sub-dependencies and their respective versions, which also ensuring that when npm install is run, the same specific versions of packages are installed, even if newer versions are available on the npm registry.

* Install chi router for routing by following command : 
go get -u github.com/go-chi/chi/v5

* Now i will create a main.go file to run the server
* Now we have to run our server so to run the server we need some configuration, that we did in code files :

* Run the server : 
```bash
go run main.go
```



# **Go Server Architecture — Flow & Dependency Injection**

This document explains the code flow of our Go application, starting from `main.go` and moving through server setup, routing, controllers, services, and repositories. It also covers our use of **constructor-based dependency injection** for flexibility and maintainability.

---

## **1. Application Entry Point — `main.go`**

When the application starts, it executes `main.go`:

```go
func main() {
	config.Load()
	cfg := app.NewConfig(":3005")
	app := app.NewApplication(cfg)
	app.Run()
}
```

**Step-by-step:**

1. **Load environment variables** using `config.Load()`.
2. **Create the configuration object** with `app.NewConfig(":3005")` — this holds server details like the address (`:3005`).
3. **Create the `Application` object** using `app.NewApplication(cfg)`.
4. **Run the application** with `app.Run()` — this method starts the HTTP server.

---

## **2. The `Run` Method in `Application`**

Inside the `Run` method, we:

1. Create dependencies (repository → service → controller → router).
2. Configure and start the HTTP server.

```go
func (app *Application) Run() error {
	// Dependency creation
	ur := db.NewUserRepository()
	us := service.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(*uc)

	// Server configuration
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on http://localhost", app.Config.Addr)
	return server.ListenAndServe()
}
```

---

## **3. Router Setup — `SetUpRouter`**

We use the [**Chi router**](https://github.com/go-chi/chi), a lightweight and idiomatic HTTP router for Go.

The `SetUpRouter` function:

1. Creates a **new Chi router**.
2. Accepts custom routers like `UserRouter`, `PostRouter`, `TweetRouter`, etc.
3. Registers routes defined in each router’s `Register` method.
4. Returns the fully configured `chi.Router`.

---

## **4. The `Router` Interface**

To ensure all routers follow the same pattern, we define a **Router interface**:

```go
type Router interface {
	Register(r chi.Router)
}
```

Any router (e.g., `UserRouter`) that implements this interface **must** provide a `Register` method.

---

## **5. Example: `UserRouter`**

```go
type UserRouter struct {
	UserController controllers.UserController
}

// Constructor
func NewUserRouter(_userController controllers.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

// Register routes
func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/signup", ur.UserController.RegisterUser)
}
```

**Flow:**

* The `UserRouter` receives a `UserController`.
* In `Register`, it adds a POST `/signup` route linked to the `RegisterUser` method in the controller.
* `SetUpRouter` calls this `Register` function, passing the shared Chi router.

---

## **6. Dependency Injection Flow**

We follow **constructor-based dependency injection**, meaning:

* Dependencies are created outside the struct itself.
* Each constructor receives its dependencies as parameters, instead of creating them internally.

**Flow in `Run` method:**

```go
// 1. Repository Layer
ur := db.NewUserRepository()

// 2. Service Layer (depends on repository interface)
us := service.NewUserService(ur)

// 3. Controller Layer (depends on service interface)
uc := controllers.NewUserController(us)

// 4. Router Layer (depends on controller)
uRouter := router.NewUserRouter(*uc)
```

---

## **7. Why Constructor-Based Dependency Injection?**

* **Loose coupling:** Concrete classes depend on **interfaces**, not other concrete classes.
* **Flexibility:** We can easily switch implementations without changing higher layers.
  Example: Switching `MySQLUserRepository` to `MongoUserRepository` is a matter of passing a different implementation to `NewUserService()`.
* **Testability:** We can inject mock repositories/services for unit testing.
* **Clarity:** All dependencies are explicit — no hidden creation of objects inside constructors.

---

## **8. Example: Switching Repository Implementation**

If `UserService` depends on a `UserRepository` interface, we can swap implementations easily:

```go
// Using MySQL
ur := db.NewMySQLUserRepository()

// Using MongoDB
// ur := db.NewMongoUserRepository()

us := service.NewUserService(ur)
```

No changes are required in the service or controller layer — only in the wiring.

---

## **9. Key Differences from Java Spring Boot**

In frameworks like **Spring Boot**:

* Dependency injection is handled automatically via annotations (`@Autowired`, `@Service`, etc.).
* The framework manages the object lifecycle and wiring.

In Go:

* We **manually create and wire dependencies** (unless using a DI framework like `wire` or `fx`).
* This manual approach gives us more control, but can become verbose in large applications — that's when DI frameworks for Go become useful.

---

## **10. Summary Diagram**

```
main.go
  ↓
Application.Run()
  ↓
Repository → Service → Controller → Router
  ↓
SetUpRouter() with chi
  ↓
http.Server.ListenAndServe()
```

---

✅ **Benefits of this architecture**:

* Clean separation of concerns.
* Easy to extend and modify.
* Test-friendly structure.
* Independent layers communicating via interfaces.
---

# Next Goal : Database connection ---> done
# how to query database : 
* short answer : Raw Queries
* Long answer : 
	1. Prepare a query(string)
	2. Execute the query using inbuild query methods, most of these methods going to return some row data to us.
	3. We need to process this row data and create an output object
	4. Return the object.



# Next Goal : Add migration(goose)

1. COMMAND FOR CREATING MIGRATION : 
goose -dir "db/migrations" create create_user_table sql

2. How to run the migration
goose -dir "db/migrations" mysql "root:mysql@1234&?@tcp(127.0.0.1:3306)/Airbnb_auth_dev" up


* Install air form github for hot reload like nodemon
```go 
go install github.com/air-verse/air@latest
```
and update your Makefile by folliwing command
dev:  # make dev
	air
Now it will see your all file and if any changes found it will recompile and run you server

# Install jwt token
```go
go get -u github.com/golang-jwt/jwt/v5
```

# Writing cusotm JSON response writer using "gin".

**JSON marshalling**
* In Go, JSON marshalling refers to the process of converting Go data structures (like structs, maps, slices, and basic types) into a JSON-formatted byte slice. This process is handled by the encoding/json package, specifically using the json.Marshal function.

## How to Marshal JSON in Go
* Define your Go data structure: This can be a struct, a map, or a simple variable. For structs, ensure that the fields you want to be marshaled have uppercase first letters (exported fields) and optionally include json struct tags to control the JSON field names.

```go
    type Person struct {
        Name    string `json:"full_name"` // `json:"full_name"` maps "Name" to "full_name" in JSON
        Age     int    `json:"age"`
        IsAdult bool   `json:"is_adult,omitempty"` // `omitempty` omits the field if its zero value
    }
```

# Full code example of Marshalling :

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}

func main() {
	user := User{ID: 1, Username: "johndoe", Email: "john.doe@example.com"}

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("Marshaled JSON:", string(jsonData))

	// Example with omitempty
	userWithoutEmail := User{ID: 2, Username: "janedoe"}
	jsonData2, err := json.Marshal(userWithoutEmail)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("Marshaled JSON (without email):", string(jsonData2))
}
```

# Adding go validator:
```go 
go get github.com/go-playground/validator/v10
```

# Homework
 * connect all the remaining apis : Signup and get user with service using json marshalling
 * write a middleware to validate every incoming request body
 * setup a good error handling mechanism


 # setting up the rate limiting

 # Homework
 * Explore ip address and jwt based rate limiting

 # working on reverse proxy
  * let's understand how it works.

  * suppose we want to make request on fakestore api  on url : https://fakestoreapi.in/api/products/category

  * Suppose from postman you are making request on url : http://localhost:3002/fakestoreService/api/products/category

  * Behind the scene working of apiGateway

  * to achieve this functionality we have defined a ProxyToService function which take two parameter 
			1. The(targetBaseUrl) base URL on which api gateway will going to make call
			2. Second one is pathPrefix, on the basis of which we remove previous of this url including this as well 

			For example suppose you have complete url : http://localhost:3002/fakestoreService/api/products/category

			and "pathPrefix" is fakestoreService then there is a method TrimPrefix on go inbuilt strings, so using this method it remove the path prefix and return a stripped path like this one is stripped path : /api/products/category

			* Finally here we combine the targetBaseUrl + stripped path which becomes something like this one : http://fakestoreapi.in/ + /api/products/category
				= http://fakestoreapi.in/api/products/category


```go  used in router.go
	chiRouter.HandleFunc("/fakestoreService/*", utils.ProxyToService("http://fakestoreapi.in/", "/fakestoreService"))
```

# Next Assignment : start writing the api gateway with other microservices

# RBAC(Role based access control implementation started;)

* completed RBAC migration here got some issue and that issue were related to 
data type incompitable

* Be cautious when you use SERIEL in id column because is serieal is equivalent to
"BIGINT UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE" and in reference if you provide 
int then you will get incompitable error.


# Detailed features of API-Gateway and how to use in nodejs 



# 🔁 Reverse Proxy in API Gateway (Golang)

In our **API Gateway**, we act as an intermediary between the client and the internal microservices (like hotel, booking, and review services).
Whenever a request hits the gateway, it forwards that request to the correct underlying service — this behavior is implemented using a **Reverse Proxy**.

---

## ⚙️ Gateway Routing Setup

We have defined routing rules in the gateway to forward incoming requests:

| Incoming Request                       | Routed To (Target Service) |
| -------------------------------------- | -------------------------- |
| `http://localhost:3002/hotelService`   | `http://localhost:3001`    |
| `http://localhost:3002/bookingService` | `http://localhost:3005`    |
| `http://localhost:3002/reviewService`  | `http://localhost:8081`    |

So for example, a request from a client to

```
http://localhost:3002/hotelService/api/v1/allHotels
```

will be automatically forwarded by the gateway to

```
http://localhost:3001/api/v1/allHotels
```

---

## 🧩 The `ProxyToService` Utility Function

We created a helper function named `ProxyToService` inside the **utils** folder.
This function takes two parameters:

1. **`targetBaseUrl`** → The actual backend service URL that the gateway should forward the request to.
2. **`pathPrefix`** → The prefix to strip from the incoming URL path before forwarding the request.

It then returns an `http.HandlerFunc` that acts as a **reverse proxy**.

---

## 🧠 How the Reverse Proxy Works Internally

### 1️⃣ Parsing the Target URL

The provided `targetBaseUrl` (like `"http://localhost:3001"`) is first parsed using:

```go
url.Parse(targetBaseUrl)
```

This ensures the target URL conforms to proper URL syntax (as per **RFC 3986**).
Invalid URLs like `"http:///bad-format"` would fail here.

---

### 2️⃣ Using `http/httputil` – Built-in Reverse Proxy

The Go standard library provides a powerful utility in the `net/http/httputil` package called:

```go
httputil.NewSingleHostReverseProxy(target *url.URL)
```

This function takes a **parsed target URL** and returns a ready-to-use reverse proxy that automatically rewrites requests to the target host.

---

### 3️⃣ Example of How It Works

Let’s say the incoming request is:

```
http://localhost:3001/api/v1/allHotels?city=paris&limit=10
```

Internally, Go parses it like this:

```text
target.Scheme   → "http"
target.Host     → "localhost:3001"
target.Path     → "/api/v1/allHotels"
target.RawQuery → "city=paris&limit=10"
```

The proxy then rewrites and forwards the request to this target.
Under the hood, it behaves like this:

```go
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy {
	director := func(req *http.Request) {
		rewriteRequestURL(req, target)
	}
	return &ReverseProxy{Director: director}
}

func rewriteRequestURL(req *http.Request, target *url.URL) {
	targetQuery := target.RawQuery
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)

	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
}
```

So essentially, the **director** function modifies the incoming request so that:

* The **scheme** and **host** point to the target service.
* The **path** and **query** are correctly merged.

---

## 🧱 When You Need More Control

The `NewSingleHostReverseProxy` method doesn’t modify request headers by default.
If you want to control or rewrite headers, you can create a custom `ReverseProxy` directly and use the `Rewrite` hook:

```go
proxy := &httputil.ReverseProxy{
    Rewrite: func(r *httputil.ProxyRequest) {
        r.SetURL(target)      // Set target URL
        r.Out.Host = r.In.Host // Optional: preserve original host header
    },
}
```

This allows you to modify request headers, cookies, or any other fields before forwarding.

---

## 🔄 What is a Director Function?

In the Go `httputil.ReverseProxy`, the **Director** is a function that modifies the incoming request before it is sent to the target backend.
Think of it as a middleware step that **rewrites** the request.

---

## 🪄 Implementation Steps in Our Gateway

Inside our handler, here’s what happens step by step:

1. Extract the original path:

   ```go
   originalPath := r.URL.Path
   ```

2. Strip the predefined prefix (`pathPrefix`) from the path.

3. Set the **target host** and **modified path** in the forwarded request.

4. Add any custom headers — for example, attaching the `userId` from request context:

   ```go
   if userId, ok := r.Context().Value("userID").(string); ok {
       r.Header.Set("X-User-ID", userId)
   }
   ```

5. Finally, the proxy handles the request:

   ```go
   proxy.ServeHTTP(w, r)
   ```

---

## 🧠 In Short

| Step                           | Description                                      |
| ------------------------------ | ------------------------------------------------ |
| **1. Parse target URL**        | Validate and convert string URL → `*url.URL`     |
| **2. Create Reverse Proxy**    | Use `httputil.NewSingleHostReverseProxy(target)` |
| **3. Rewrite Request**         | Director modifies scheme, host, path, and query  |
| **4. Add Custom Headers**      | Example: Add user ID or trace ID                 |
| **5. Forward to Microservice** | Proxy sends modified request to target service   |



### 🧭 Summary

* **Reverse Proxy** allows your API Gateway to forward requests to internal microservices.
* **`httputil.NewSingleHostReverseProxy()`** handles most of the rewriting automatically.
* You can customize headers or modify logic using the **`Rewrite`** or **`Director`** functions.
* Prefix stripping and context propagation (like `userID`) make the request routing flexible and secure.


## 🧩 **Reverse Proxy: Go vs Node.js – Quick Comparison**

| Concept                   | **Golang Implementation**                                     | **Node.js Implementation**                                      |
| ------------------------- | ------------------------------------------------------------- | --------------------------------------------------------------- |
| **Framework / Core Tool** | `net/http` package                                            | `Express.js` (most common web framework)                        |
| **Reverse Proxy Library** | `httputil.NewSingleHostReverseProxy`                          | `http-proxy-middleware` (built on `http-proxy`)                 |
| **Purpose**               | Forwards incoming requests to target microservices            | Same — acts as gateway forwarding requests to target services   |
| **Path Rewrite**          | Manual logic (`strings.TrimPrefix` etc.)                      | Built-in option → `pathRewrite`                                 |
| **Header Manipulation**   | Done in Director function (e.g., `r.Header.Set("X-User-ID")`) | Done in `onProxyReq` hook (`proxyReq.setHeader()`)              |
| **Target URL Handling**   | `url.Parse()` + validation                                    | Direct string URL in config                                     |
| **Error Handling**        | Custom `ServeHTTP` implementation                             | `onError` callback in middleware                                |
| **Middleware Support**    | Need to build manually                                        | Express middlewares available (auth, rate limit, logging, etc.) |
| **Performance**           | Faster (compiled language)                                    | Slightly slower but more flexible for rapid development         |
| **Example Package Names** | Built-in (`net/http`, `httputil`)                             | `http-proxy-middleware`, `http-proxy`, `express-gateway`        |

---

## 🧠 **How to Explain This in an Interview (Sample Answer)**

> “In my Go-based API Gateway, I used the `httputil.NewSingleHostReverseProxy` from the standard library. It rewrites incoming requests and routes them to target microservices after parsing and validating the target URL.
>
> If I had to build the same in Node.js with TypeScript, I’d use **Express.js** along with the **http-proxy-middleware** package — which provides similar functionality to Go’s `ReverseProxy`.
>
> It allows path rewriting, custom header injection, and error handling out of the box. I can define routes like `/hotelService → http://localhost:3001` and `/bookingService → http://localhost:3005` easily with only a few lines of code.
>
> Essentially, Go’s `httputil` reverse proxy and Node’s `http-proxy-middleware` solve the same problem — Go gives you more control at a lower level, while Node.js provides faster iteration with middleware flexibility.”

---

## ⚙️ **Node.js Key Packages to Mention**

| Package                   | Use Case                                  |
| ------------------------- | ----------------------------------------- |
| **http-proxy-middleware** | Most common reverse proxy for Express     |
| **http-proxy**            | Low-level proxy library used by the above |
| **express-gateway**       | Full-featured API Gateway framework       |
| **morgan / winston**      | Logging incoming and outgoing requests    |
| **express-rate-limit**    | To add rate limiting before proxying      |
| **helmet / cors**         | For securing gateway routes               |

---




* Task to implement
1. Try to implement an email confirmation mechanism for new user signup
2. On a new User signup, people should get automatically role of 'user'
3. Try to integrate any MFA mechanism like EMAIL OTP || SMS OTP || MFA APP
