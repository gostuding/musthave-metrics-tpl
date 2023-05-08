package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gostuding/musthave-metrics-tpl/cmd/server/handlers"
	"github.com/gostuding/musthave-metrics-tpl/cmd/server/storage"
)

var memory = storage.MemStorage{}

func GetRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllMetrics(w, r, memory)
	})
	router.Post("/update/{m_type}/{m_name}/{m_value}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Update(w, r, &memory)
	})
	router.Get("/value/{m_type}/{m_name}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetMetric(w, r, memory)
	})

	return router
}
