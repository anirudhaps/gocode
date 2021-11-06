package server

import (
	"fmt"
	"net/http"
	"testserver/input"
	"time"
)

func TestServer(w http.ResponseWriter, r *http.Request) {
	params := input.GetParams()
	headers := params.ResponseHeaders()
	for key, val := range headers {
		w.Header().Add(key, val)
	}

	if params.SleepTime() != 0 {
		fmt.Println("Going to sleep for:", params.SleepTime())
		time.Sleep(params.SleepTime())
	}

	if params.Verbose() {
		fmt.Println("\nPath:", r.URL.Path[1:])
		fmt.Println("\nRequest Headers:")
		fmt.Println("Host:", r.Host)
		for key, val := range r.Header {
			fmt.Println(key, ":", val)
		}

		if headers != nil {
			fmt.Println("\nResponse Headers (sent):")
			for key, val := range headers {
				fmt.Println(key, ":", val)
			}
		}
	}

	// TODO: add debug logs and remove fmt.Println()
	// TODO: add logic to read and serve the file whose name received in the path
	// TODO: Add readme and unit tests, test it manually with envoy
	w.WriteHeader(params.StatusCode())
}
