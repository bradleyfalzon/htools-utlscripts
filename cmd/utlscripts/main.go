package main

import (
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

// var forceFlag = flag.Bool("f", false, "Force output generation")
var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

// main is the entry point for the application.
func main() {
	t := T{}
	log.SetOutput(os.Stderr)

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
