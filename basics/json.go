package main

import (
	"encoding/json"
	"fmt"
)

type host struct {
	ip   string `json: "ip"`
	port uint16 `json: "port"`
	Host string `json: "host"`
}

func main() {
	// Note: the json package only considers exported members while marshaling.
	// i.e. the members whose name start with a capital letter. All other
	// members are ignored. Tip: run this code and see.
	myhost := host{ip: "1.2.3.4", port: 80, Host: "localhost"}
	b, err := json.Marshal(myhost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
