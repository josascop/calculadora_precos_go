package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func alou() {
	c := chi.NewRouter()

	http.ListenAndServe(":8000", c)
}
