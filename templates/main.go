package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

var (
	templateString = "{{range .}}\n{{.}}{{end}}"
)

type A struct {
	TableName string
	Count     int
}

func main() {
	t := template.New("t")
	t, err := t.Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	data := []A{
		{TableName: "blahblah", Count: 10},
		{TableName: "loopy", Count: 40},
		{TableName: "Proppy", Count: 25},
	}

	var out bytes.Buffer
	t.Execute(&out, data)

	fmt.Println(out.String())
}
