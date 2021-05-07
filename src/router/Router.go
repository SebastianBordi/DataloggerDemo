package router

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}
