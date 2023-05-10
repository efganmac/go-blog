package main

import (
	"blogproject/config"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", config.Routes())
}
