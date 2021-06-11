package main

import (
	"entryTask/main/cgo"
	"entryTask/main/controller"
	"log"
	"net/http"
	"time"
)

func main() {

	cgo.InitDB()

	server := &http.Server{
		Addr:        ":8080",
		Handler:     cgo.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(cgo.Router)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func RegiterRouter(handler *cgo.RouterHandler) {
	new(controller.UserConterller).Router(handler)
}
