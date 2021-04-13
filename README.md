# Personal Go Project with goyave and sqlboiler

I'm using goyave as for this project and sqlboiler for replace GORM <br>
[Goyave](https://github.com/System-Glitch/goyave) framework. <br>
I use it for learning purpose, you can use it if it suits your needs with your own responsibility.

## Getting Started

### Requirements

- Go 1.13+
- Go modules

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then try to request the `hello` route.
```
$ curl http://localhost:8080/hello
Hi!
```

There is also an `echo` route, with a basic body validation.
```
$ curl -H "Content-Type: application/json" -X POST -d '{"text":"abc 123"}' http://localhost:8080/echo
abc 123
```

## License

for goyave license, can refer to https://github.com/go-goyave/goyave/blob/master/LICENSE
