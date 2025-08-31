
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

