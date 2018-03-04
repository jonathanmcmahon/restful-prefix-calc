# restful-prefix-calc
A RESTful normal Polish notation calculator.

This project is an example of a simple RESTful API written in Go.

## Instructions ##

Get the code: 

```
go get github.com/jonathanmcmahon/restful-prefix-calc
```

Start the http server:

```
cd `go env GOPATH`/src/github.com/jonathanmcmahon/restful-prefix-calc/
go run calculator.go
```

Now the HTTP server is listening on *localhost:8000*. You can query it as below:

**Example query:**

```
curl -X GET http://localhost:8000/mul/9/3/3
```

**Response:**
```
81.000000
```