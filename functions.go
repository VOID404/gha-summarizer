package main

import "os"

var funcMap map[string]any

func init() {
	funcMap = make(map[string]any)
	funcMap["file"] = File
	funcMap["env"] = Env
}

// File reads a file and returns its content as a string
func File(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Env expands environment variables in a string
func Env(text string) string {
	return os.ExpandEnv(text)
}
