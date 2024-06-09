package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/odundlaw/go_movies/config"
	"github.com/odundlaw/go_movies/store"
)

type MovieController struct{}

func (rs MovieController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetAll)
	r.Post("/", rs.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Delete("/", rs.Delete)
		r.Patch("/", rs.Update)
	})

	return r
}

func (rs MovieController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movieStore := config.MovieStore

	var err error
	var movie *store.Movie

	if movieId := chi.URLParam(r, "id"); movieId != "" {
		movie, err = movieStore.GetOne(movieId)
	} else {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs MovieController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movieStore := config.MovieStore

	movies := movieStore.GetAll()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs MovieController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movieStore := config.MovieStore
	var err error
	var movie *store.Movie

	if err = json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if movieId := chi.URLParam(r, "id"); movieId != "" {
		movie, err = movieStore.UpdateOne(movieId, movie)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs MovieController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movieStore := config.MovieStore
	var err error

	if movieId := chi.URLParam(r, "id"); movieId != "" {
		_, err = movieStore.DeleteOne(movieId)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rs MovieController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movieStore := config.MovieStore

	var err error
	var movie *store.Movie

	if err = json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if movie, err = movieStore.Create(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
