package main

import (
	"flag"
	"log"
	"os"
)

var forceFlag = flag.Bool("f", false, "Force output generation")

// main is the entry point for the application.
func main() {
	log.SetOutput(os.Stderr)

	flag.Parse()
	if *forceFlag {
		log.Printf("Error loading cluster configuration:")
		return
	}

}
