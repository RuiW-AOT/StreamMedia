package main

import (
	"net/http"

	"github.com/RuiW-AOT/StreamMedia/server/scheduler/tasksrunner"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	go tasksrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
