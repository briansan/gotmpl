package main

// Usage: gotmpl <file-name> <args-yaml-file>

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: gotmpl <file-name> <args-yaml-file>\n")
		os.Exit(1)
	}

	// Read template file.
	fname := os.Args[1]
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	t := template.Must(template.New("tmpl").Parse(string(b)))

	// Read object file.
	argsFile := os.Args[2]
	if b, err = ioutil.ReadFile(argsFile); err != nil {
		panic(err)
	}

	// Parse object as yaml.
	var args map[string]interface{}
	if err = yaml.Unmarshal(b, &args); err != nil {
		panic(err)
	}
	if err = t.Execute(os.Stdout, args); err != nil {
		panic(err)
	}
}
