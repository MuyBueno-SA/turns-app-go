package main

import (
	"net/http"
	"turns-app-go/server"
)

func main() {
	s := server.NewAPPServer(nil)
	http.ListenAndServe(":5000", s)
}
