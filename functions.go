package main

import "os"

var funcMap map[string]any

func init() {
	funcMap = make(map[string]any)
	funcMap["file"] = File
	funcMap["filep"] = Filep
	funcMap["env"] = Env
	funcMap["envp"] = Envp
}

// File reads a file and returns its content as a string
func File(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Filep checks if a file exists
func Filep(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// Env expands environment variables in a string
func Env(text string) string {
	return os.ExpandEnv(text)
}

// Envp checks if an environment variable is set
func Envp(name string) bool {
	_, ok := os.LookupEnv(name)
	return ok
}
