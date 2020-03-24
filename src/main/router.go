package main

import "net/http"

func RoutesLayout() {
	http.HandleFunc("/",LoggingMiddleware(CheckAuthMiddleware(Index))  )
	http.HandleFunc("/home", addMiddlewares(Home, CheckAuthMiddleware, LoggingMiddleware) )
}
