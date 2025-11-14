**Answer why you choose Golang Instead of Typescript to build the ApiGateway.**

Visit your notion notes for this answer : https://www.notion.so/Why-you-choose-Golang-for-ApiGateway-instead-of-Typescript-288ce7b4693280299e20c1c97dcda39c


# Features of API Gateway we build: 

### This API Gateway is built using Golang and provides core functionalities such as Authentication, Authorization, Reverse Proxy, and Rate Limiting. It follows microservice principles and enforces secure access through JWT and RBAC (Role-Based Access Control).

1. Authentication 
    i. User signedUp using name email and password before save user to db hash it's password using bcrypt.GenerateFromPassword method also assign a default role of "user".

    ii. User signIn using email and password Before Generating jwt token match hashed password(using bcrypt.CompareHashAndPassword) then generate jwt using jwt.NewWithClaims method.

    iii. When anyone wants to check their profile they must have either user or admin role other wise it will say unauthorised user.

---

## 🔐 JWT Signing Algorithms Overview

* // don't worrry about the header it's automatically generated when you add diff algo in the options field like below :

// jwt.sign(payload, secret, { algorithm: "RS256", expiresIn: "1h" });


When working with **JWT (JSON Web Tokens)**, the way we **sign and verify** tokens depends on the type of algorithm used.
There are mainly **two types of algorithms** for signing JWTs:

---

### 🧩 1. Symmetric Algorithms — **HMAC (HS256, HS384, HS512)**

* Hash-based Message Authentication Code(HMAC)
* Uses **one single shared secret key** for both **signing** and **verification**.
* The **same key** is used by whoever creates the token and whoever verifies it.
* ✅ **Fast** and efficient
* ⚠️ **Less secure** in distributed environments because the same secret must be shared across services.

**Example Use Case:**

> Best for **Monolithic applications**, where the same server handles both token generation and verification.

---

### 🔑 2. Asymmetric Algorithms — **RSA (RS256, RS384, RS512)** or **ECDSA (ES256)**

* The acronym RSA stands for Rivest, Shamir, Adleman, the three inventors who first publicly described the algorithm in 1977.

* Uses **two keys**:

  * **Private key** → used for signing the token
  * **Public key** → used for verifying the token
* ✅ **More secure**, since only the private key needs to be protected
* ⚠️ **Slightly slower** and **computationally more expensive**

**Example Use Case:**

> Best for **Microservices or Distributed systems**, where one service (e.g., API Gateway) signs the token, and other services verify it using the public key.

---

## 🏗️ Recommended Algorithms by Architecture

| **Architecture**                    | **Recommended Algorithm**            | **Why**                                                                                                                                                                                                                        |
| ----------------------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Monolith (Single Server)**        | **HMAC (HS256)**                     | Simple, fast, and sufficient. Since only one server signs and verifies, a shared secret key is perfectly acceptable.                                                                                                           |
| **Microservices / Distributed API** | **RSA (RS256)** or **ECDSA (ES256)** | Essential for distributed systems. The API Gateway signs the token with a private key, and all downstream microservices verify it with the public key. This allows secure verification without sharing the secret signing key. |

---

### 🧠 In Simple Terms

* **HMAC (HS256)** → One key for everything → fast but less secure if shared.
* **RSA / ECDSA (RS256 / ES256)** → Two keys (private + public) → safer but slower.
* Choose **HMAC** for simple single-server setups, and **RSA/ECDSA** for distributed systems or microservices.

---



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

* In golang there is a package named "rate package" available on github which take the ip, time and the limit as per the given time.

* we can extract ip from the req body in node.js ver simple (req.ip) and in golang ----> ip, _, err := net.SplitHostPort(r.RemoteAddr)

* In node js there is npm package named express-rate-limit which take the time in ms(windowMs) and limit, and as a middleware we can put it before any handler.

Absolutely ✅ — here’s a **clean 50-line summary** version of your full write-up.
It keeps all **core ideas, technical reasoning, and interview-useful points** while removing long explanations.
You can easily revise and recall this before interviews 👇

---

# 🔁 **Reverse Proxy in API Gateway – Summary**

1. The **API Gateway** acts as an intermediary between clients and internal microservices.

2. When a request hits the gateway, it’s forwarded to the correct target service.

3. Example routing:

   * `/hotelService` → `http://localhost:3001`
   * `/bookingService` → `http://localhost:3005`
   * `/reviewService` → `http://localhost:8081`

4. Implemented via a helper function **`ProxyToService`**.

5. `ProxyToService(targetBaseUrl, pathPrefix)` returns an `http.HandlerFunc`.

6. It performs **reverse proxying** — forwarding and rewriting requests.

7. The target URL is parsed using `url.Parse(targetBaseUrl)` (validates RFC 3986).

8. Go’s standard package `net/http/httputil` provides `NewSingleHostReverseProxy`.

9. This proxy rewrites the incoming request → sets target **scheme**, **host**, **path**, and **query**.

10. Example: Incoming `/hotelService/api/v1/allHotels` → routed to `/api/v1/allHotels` on target.

11. Internally, a **Director function** modifies the request before forwarding.

12. Director updates the request’s `URL`, merges query params, and preserves structure.

13. The proxy then calls `ServeHTTP(w, r)` to forward the modified request.

14. When more control is needed (e.g., modify headers), a custom `ReverseProxy` is created.

15. The `Rewrite` hook (`ProxyRequest.SetURL(target)`) allows header/cookie manipulation.

16. Example: Add `X-User-ID` header from request context before forwarding.

17. The proxy chain:
    **Client → Gateway → Reverse Proxy → Target Microservice.**

18. Benefits: central routing, authentication, logging, header injection, and API version control.

19. In short steps:

    1. Parse target URL
    2. Create reverse proxy
    3. Rewrite request (scheme, host, path, query)
    4. Add headers (userID, correlationID)
    5. Forward to service

---

## ⚙️ **Equivalent Implementation in Node.js**

20. In Node.js, use **Express.js** + **http-proxy-middleware** package.
21. It performs the same as Go’s `httputil.NewSingleHostReverseProxy`.
22. Example route:

```ts
app.use('/hotelService', createProxyMiddleware({
  target: 'http://localhost:3001',
  pathRewrite: { '^/hotelService': '' },
  changeOrigin: true,
}));
```

23. `pathRewrite` → removes prefix like `/hotelService`.

24. `onProxyReq` → modify request headers before sending to backend.

25. `onError` → handle backend service or connection failures.

26. Node’s proxy middleware supports logging, rate limiting, and middleware chaining easily.

27. Go gives lower-level control, Node.js offers flexibility and faster iteration.

---

## 🧩 **Go vs Node.js Summary Table**

| Concept        | Golang                               | Node.js                               |
| -------------- | ------------------------------------ | ------------------------------------- |
| Framework      | `net/http`                           | `Express.js`                          |
| Proxy Library  | `httputil.NewSingleHostReverseProxy` | `http-proxy-middleware`               |
| Path Rewrite   | Manual                               | `pathRewrite` option                  |
| Header Edit    | Director function                    | `onProxyReq` hook                     |
| Error Handling | Custom ServeHTTP                     | Built-in callbacks                    |
| Performance    | Faster                               | More flexible                         |
| Example Tools  | Built-in                             | `express-gateway`, `morgan`, `helmet` |

---

## 🧠 **Interview Summary Answer (20-sec Pitch)**

> “In Go, I implemented reverse proxying using `httputil.NewSingleHostReverseProxy`.
> It rewrites incoming requests, strips path prefixes, and forwards them to internal services, adding headers like userID from context.
> If I had to build the same in Node.js, I’d use `Express` with `http-proxy-middleware`, which offers similar path rewriting, header injection, and error handling features.
> Both follow the same reverse proxy pattern — Go gives fine control, Node provides middleware simplicity.”

---