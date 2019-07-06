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

## Run server locally
`$ go run main.go`

## Access server locally
Try access server locally via localhost, for example:
`$ curl -i localhost:8090/hello`
