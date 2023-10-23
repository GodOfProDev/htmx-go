package main

import "htmx-go/internal/router"

func main() {
	router := router.NewRouter()
	router.Listen()
}
