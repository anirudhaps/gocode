package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", EchoServer)
	http.ListenAndServe(":8080", nil)
}

func EchoServer(w http.ResponseWriter, r *http.Request) {
	// writes back the path string and its length
	if len(r.URL.Path[1:]) == 0 {
		fmt.Fprintf(w, "{value: nil, len: 0}")
	} else {
		fmt.Fprintf(w, "{value: %s, len: %d}", r.URL.Path[1:], len(r.URL.Path[1:]))
	}
}
