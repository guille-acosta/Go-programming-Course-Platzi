package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		page404(w, r)
		return
	}
	fmt.Fprint(w, "Welcome to Index")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home")
}

func page404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "Web site not Found 404")
}
