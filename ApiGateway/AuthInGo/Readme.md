
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

