# Micro service test
This is a simple golang micro service test

## Run tests locally
```
$ go test -v ./...
?       github.com/karl-johan-grahn/microtest   [no test files]
=== RUN   TestRouter
--- PASS: TestRouter (0.01s)
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      github.com/karl-johan-grahn/microtest/handlers  0.525s
```

## Run service
See the Makefile targets for how to run the service as a docker container,
or run it via golang:
```
$ PORT=8090 go run main.go
```

## Access server locally
Try access server locally via localhost, for example:
`$ curl -i localhost:8090/hello`
