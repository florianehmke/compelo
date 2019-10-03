// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	scripts := http.Dir(".")
	err := vfsgen.Generate(scripts, vfsgen.Options{
		PackageName:  "scripts",
		VariableName: "Scripts",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
