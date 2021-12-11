//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	frontend := http.Dir("compelo/dist")
	err := vfsgen.Generate(frontend, vfsgen.Options{
		PackageName:  "frontend",
		VariableName: "Frontend",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
