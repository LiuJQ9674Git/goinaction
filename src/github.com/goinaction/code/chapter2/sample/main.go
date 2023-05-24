package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"golang.org/x/time/rate"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
	//rateLimiter *rate.Limiter
}

type APIConnection struct {
	rateLimiter *rate.Limiter
}
