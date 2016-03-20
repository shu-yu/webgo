package webgo

import (
	Er "errors"
	"log"
	"os"
)

type Errors struct {
	Log *log.Logger

	// Errors in JSON response is a array of string containing all error messages
	Errs []string

	// This is used to add custom error types
	AppErr map[string]error

	// C for `Code`
	C001 error
	C002 error
	C003 error
	C004 error
	C005 error
	C006 error
}

func (e *Errors) Init(errTypes map[string]error) {

	// Error codes which are to be used through out the app.
	// App configuration errors
	e.C001 = Er.New("Invalid number of arguments provided")
	e.C002 = Er.New("Could not unmarshal JSON config file")
	e.C003 = Er.New("App environment not provided in config file. Accepted values are `production` or `development`")
	e.C004 = Er.New("App port not provided in config file")
	e.C005 = Er.New("Invalid JSON")

	Err.Log = log.New(os.Stderr, "app: ", log.LstdFlags|log.Llongfile)
	e.AppErr = errTypes
}

// Global variable to access Error logging structure.
var Err Errors
