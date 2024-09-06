package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <template file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	tplFile := os.Args[1]
	outFile, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	funcMap := make(map[string]any)
	funcMap["file"] = func(name string) string {
		data, err := os.ReadFile(name)
		if err != nil {
			panic(err)
		}
		return string(data)
	}

	funcMap["env"] = func(text string) string {
		return os.ExpandEnv(text)
	}

	tpl := template.Must(
		template.New(path.Base(tplFile)). //
							Funcs(funcMap).
							ParseFiles(tplFile),
	)

	err = tpl.Execute(outFile, nil)
	if err != nil {
		panic(err)
	}
}
