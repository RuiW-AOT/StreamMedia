package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RuiW-AOT/StreamMedia/server/api/defs"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &middleWareHandler{}
	m.r = r
	return m
}

func skipSessionValidation(method, url string) bool {
	fmt.Println(method, url)
	if method == "POST" && strings.HasPrefix(url, "/user") {

		return true
	}

	return false
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session

	if !skipSessionValidation(r.Method, r.URL.Path) && !validateUserSession(r) {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:username", Login)

	router.GET("/user/:username", GetUserInfo)

	router.POST("/user/:username/videos", AddNewVideo)

	router.GET("/user/:username/videos", ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	router.POST("/videos/:vid-id/comments", PostComment)

	router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
