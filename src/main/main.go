package main

import (
	"log"
	"net/http"
)

func main() {
	RoutesLayout()
	log.Println(http.ListenAndServe(":5000", nil))
}
