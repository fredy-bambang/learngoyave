# Personal Go Project with goyave and sqlboiler

I'm using goyave as for this project and sqlboiler for replace GORM <br>
[Goyave](https://github.com/System-Glitch/goyave) framework. <br>
I use it for learning purpose, you can use it if it suits your needs with your own responsibility.

## Getting Started
get goose db migration 
go get -u github.com/pressly/goose/cmd/goose

sqlboiler v4
go get -u -t github.com/volatiletech/sqlboiler/v4
go get -u -t github.com/volatiletech/null/v8
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql

<!-- goose -dir ./database/sqlboiler/migrations sqlite3 "user=postgres password=password dbname=todo sslmode=disable" up -->
running the migration 
goose -dir ./database/sqlboiler/migrations sqlite3 ./foo.db up

create new migration file 
goose -dir ./database/sqlboiler/migrations create create_table_task sql

generate model 
sqlboiler --add-soft-deletes psql

### Requirements

- Go 1.17+
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
