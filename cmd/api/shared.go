package main

import "net/http"

func makeHandler(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(next)
}
