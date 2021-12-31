package routes

import (
	"project1/controllers"
	"project1/middleware"

	"github.com/gorilla/mux"
)

var ApiRouter *mux.Router

func InitApi() {
	ApiRouter = BaseRouter.PathPrefix("/api").Subrouter()

	ApiRouter.Use(middleware.CorsMiddlware)
	ApiRouter.Use(middleware.ApiMiddlware)

	InitTodo()
	InitPost()
}

func InitTodo() {
	ApiRouter.Methods("GET").Path("/todos").HandlerFunc(controllers.ListTodo)
	ApiRouter.Methods("POST").Path("/todos").HandlerFunc(controllers.StoreTodo)
}

func InitPost() {
	ApiRouter.Methods("GET").Path("/posts").HandlerFunc(controllers.ListPost)
	ApiRouter.Methods("POST").Path("/posts").HandlerFunc(controllers.StorePost)
}
