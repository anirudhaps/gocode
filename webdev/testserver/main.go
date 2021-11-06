package main

import (
	"net/http"
	_ "testserver/input"
	"testserver/server"
)

func main() {
	http.HandleFunc("/", server.TestServer)
	http.ListenAndServe(":8080", nil)
}
