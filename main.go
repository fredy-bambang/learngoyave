package main

import (
	"os"

	"github.com/fredy-bambang/learngoyave/http/route"
	_ "github.com/fredy-bambang/learngoyave/http/validation"

	"goyave.dev/goyave/v3"
	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v3/database/dialect/mysql"
	// _ "goyave.dev/goyave/v3/database/dialect/postgres"
	// _ "goyave.dev/goyave/v3/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v3/database/dialect/mssql"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
