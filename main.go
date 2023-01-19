package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/derricw/soulmsg/msg"
)

func response(w http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("%s\n", msg.RandomMessage(false))
	fmt.Fprintf(w, msg)
}

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.HandleFunc("/", response)
	err := http.ListenAndServe(":"+httpPort, nil)
	fmt.Printf("%s", err)
}
