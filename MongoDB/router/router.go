package router

import (
	"github.com/gorilla/mux"
	"github.com/kashyapt/MongoDB/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")
	return router
}

//In Go, "mux" typically refers to a multiplexer, which is a component used in web development to handle
//HTTP requests by matching the incoming request URL pattern to a specific handler function.
