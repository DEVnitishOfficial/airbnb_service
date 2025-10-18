**Answer why you choose Golang Instead of Typescript to build the ApiGateway.**

Visit your notion notes for this answer : https://www.notion.so/Why-you-choose-Golang-for-ApiGateway-instead-of-Typescript-288ce7b4693280299e20c1c97dcda39c


# Features of API Gateway we build: 

### This API Gateway is built using Golang and provides core functionalities such as Authentication, Authorization, Reverse Proxy, and Rate Limiting. It follows microservice principles and enforces secure access through JWT and RBAC (Role-Based Access Control).

1. Authentication 
    i. User signedUp using name email and password before save user to db hash it's password using bcrypt.GenerateFromPassword method also assign a default role of "user".

    ii. User signIn using email and password Before Generating jwt token match hashed password(using bcrypt.CompareHashAndPassword) then generate jwt using jwt.NewWithClaims method.

    iii. When anyone wants to check their profile they must have either user or admin role other wise it will say unauthorised user.

    iv. Here we have implemented RBAC(Role based access control) and total 5 tables which is as follows :
        a. user
        b. role
        c. permission
        d. user_role
        e. role_permission

        * user : user table have only user related info like name, email, password.

        * role : role table have name of role and their description.

        * permission : this table have name(name of permission) description(des of permission) resource(to whom provided this permission) action(what action they can perform)

        * role_permission : roleId and permissionId it's a through table wich role have whcih type of permission

        * user_role : userId and roleId it's also a throught table store which user have which role.

## RBAC(Role Based Access Contorl)
 * In this complete Auth microservices if any user want to do any operation like 
        User_specific : view user profile, getUserById, getAlluser,
        Role_specific : createRole, readRole, updateRole, deleteRole 
        Permission_specific : create, read, update, delete Permisssion
        Role_permission : create, read, update, delete Role_Permisssion
        User_Role : create, read, update, delete User_Role

    All the above required jwt authentication and a permission of role either ["user" or "admin"] except signup and login require only name, email and password.

* Implemented rate limiting middleware on the basis of ip address in one minute there will be max 5 req and then it will be blocked until time completes one minutes.

### Added Reverse proxy 

* to check reverseProxy paste the below url in the postman
	// http://localhost:3002/fakestoreService/api/products/category

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

* Gateway--> gateway.go

```go
func NewGatewayRouter() http.Handler {
	r := chi.NewRouter()

	// Forward requests to HotelService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/hotelService/*", utils.ProxyToService(
		"http://localhost:3001", // Target service
		"/hotelService",         // Prefix to strip
	))

	// Forward requests to BookingService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/bookingService/*", utils.ProxyToService(
		"http://localhost:3005",
		"/bookingService",
	))

	// Forward requests to ReviewService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/reviewService/*", utils.ProxyToService(
		"http://localhost:8081",
		"/reviewService",
	))

	return r
}

```

