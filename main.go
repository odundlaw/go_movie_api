package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/odundlaw/go_movies/config"
	controller "github.com/odundlaw/go_movies/controllers"
	"github.com/odundlaw/go_movies/store"
)

func main() {
	config.MovieStore = store.NewStore()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Mount("/movie", controller.MovieController{}.Routes())

	fmt.Println("server connected...")
	http.ListenAndServe(":5000", router)
}
