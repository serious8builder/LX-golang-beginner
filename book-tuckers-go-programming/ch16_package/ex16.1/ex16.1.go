package main

import (
	"fmt"
	htmltemplate "html/template"
	"os"

	"math/rand"
	"text/template"
)

type Member struct {
	Name string
}

func main() {
	fmt.Println(rand.Int())

	tmpl, err := template.New("foo").Parse(`{{define "T"}}Hello{{end}}{{template "T"}}, {{.Name}}`)
	if err != nil {
		panic(err)
	}

	member := Member{Name: "RK"}

	tmpl.Execute(os.Stdout, member)

	tmpl2, err2 := htmltemplate.New("foo").Parse(`{{define "T"}}Hello{{end}}{{template "T"}}`)
	if err2 != nil {
		panic(err)
	}
	fmt.Println()
	tmpl2.Execute(os.Stdout, "")
}
