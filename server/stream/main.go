package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	RateLimit = 2
)

type middleWareHandler struct {
	router      *httprouter.Router
	connLimiter *ConnectionLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := &middleWareHandler{
		router:      r,
		connLimiter: NewConnectionLimiter(cc),
	}
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.connLimiter.GetConnection() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	m.router.ServeHTTP(w, r)
	defer m.connLimiter.ReleaseConnection()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/videos/:vid-id", uploadHandler)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, RateLimit)
	http.ListenAndServe(":9000", mh)
}
