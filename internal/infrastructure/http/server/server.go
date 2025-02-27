package server

import (
	"net/http"
	"test/internal/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	addressController *controllers.AddressController
}

func New(addressController *controllers.AddressController) *Server {
	return &Server{
		addressController: addressController,
	}
}

func (s *Server) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Добавляем healthcheck эндпоинт
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/address", func(r chi.Router) {
			r.Post("/search", s.addressController.Search)
			r.Post("/geocode", s.addressController.Geocode)
		})
	})

	return r
}
