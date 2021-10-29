package lib

import "fmt"

var languages map[string]string

// This is a standard function which gets invoked when the package is imported in other packages.
// function name "init" starts with a small letter to signify that this function is not exported.
func init() {
	languages = make(map[string]string)
	languages["cs"] = "C#"
	languages["js"] = "JavaScript"
	languages["ts"] = "TypeScript"
	languages["go"] = "Golang"
	languages["rs"] = "Rust"
	languages["rb"] = "Ruby"
	fmt.Println("lib.init(): success")
}

// function name starts with capital letter ("G") to make it exported function
func Get(key string) string {
	return languages[key]
}

func Add(key, value string) {
	languages[key] = value
}

func GetAll() map[string]string {
	return languages
}
