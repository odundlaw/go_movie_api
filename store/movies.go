package store

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Store struct {
	movies []*Movie
	mu     sync.Mutex
}

type Movie struct {
	Director *Director `json:"director"`
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewStore() *Store {
	return &Store{
		movies: []*Movie{},
	}
}

func (s *Store) Create(data *Movie) (*Movie, error) {
	if data == nil {
		return nil, errors.New("movie data is empty")
	}

	if data.Title == "" {
		return nil, errors.New("movie title is empty")
	}

	if data.Director == nil {
		return nil, errors.New("movie director details is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.generateUniqueID()
	data.ID = id

	s.movies = append(s.movies, data)

	return data, nil
}

func (s *Store) GetOne(id string) (*Movie, error) {
	if id == "" {
		return nil, errors.New("params id is empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, movie := range s.movies {
		if movie.ID == id {
			return movie, nil
		}
	}

	return nil, errors.New("cannot find movie with such Id")
}

func (s *Store) GetAll() []*Movie {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.movies
}

func (s *Store) generateUniqueID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), len(s.movies))
}

func (s *Store) UpdateOne(id string, data *Movie) (*Movie, error) {
	if data == nil {
		return nil, errors.New("movie data is empty")
	}

	if data.Title == "" {
		return nil, errors.New("movie title is empty")
	}

	if data.Director == nil {
		return nil, errors.New("movie director details is required")
	}

	if data.ID == "" {
		data.ID = id
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for idx, movie := range s.movies {
		if movie.ID == id {
			s.movies[idx] = data
			return data, nil
		}
	}

	return nil, errors.New("cannot find movies with id")
}

func (s *Store) DeleteOne(id string) (bool, error) {
	if id == "" {
		return false, errors.New("params id is empty")
	}

	for idx, movie := range s.movies {
		if movie.ID == id {
			s.movies = append(s.movies[:idx], s.movies[idx+1:]...)
			return true, nil
		}
	}

	return false, errors.New("cannot find movie with id")
}
