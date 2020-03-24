package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		flag := true
		if flag {
			next.ServeHTTP(w, request)
		} else {
			fmt.Fprint(w, "No authenticated")
			return
		}
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		now := time.Now()
		defer func() {
			log.Println(request.URL.Path, now)
		}()
		next.ServeHTTP(w, request)
	}
}

func addMiddlewares(next http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		next = m(next)
	}

	return next
}
