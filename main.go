package main

import (
	"log"
	"net/http"

	"github.com/biFebriansyah/demo_flyio/libs/router"
)

func main() {
	route := router.NewApp()

	log.Println("app running on http://127.0.0.1:3001")
	err := http.ListenAndServe(":3001", route)
	if err != nil {
		log.Fatal(err)
	}
}
