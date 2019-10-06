package main


import(
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/RuiW-AOT/StreamMedia/server/api/handlers"
)
func RegisterHandlers()  *httprouter.Router{
	router := httprouter.New()

	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:user_name", handlers.Login)
	return router
}
func main(){
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}