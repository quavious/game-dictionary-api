package main

import (
	"log"

	"github.com/go-chi/chi/v5"
)

type App struct {
	infoLog      *log.Logger
	errorLog     *log.Logger
	clientID     string
	clientSecret string
	accessToken  string
	mux          *chi.Mux
}
