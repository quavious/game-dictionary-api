package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *App) Register() {
	app.mux.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(rw, r)
		}
		rw.Write([]byte("Index Page"))
	})
	app.mux.Get("/games/{page}/{option}", func(rw http.ResponseWriter, r *http.Request) {
		_page := chi.URLParam(r, "page")

		option := chi.URLParam(r, "option")
		page, err := strconv.Atoi(_page)
		if err != nil || r.Method != "GET" {
			app.LogError(err)
			rw.WriteHeader(http.StatusBadRequest)
			http.NotFound(rw, r)
			return
		}
		msg := app.GetGameList(page, option, &rw)
		rw.Write(msg)
	})
	app.mux.Get("/games/search/{term}/{page}", func(rw http.ResponseWriter, r *http.Request) {
		_page := chi.URLParam(r, "page")

		term := chi.URLParam(r, "term")
		page, err := strconv.Atoi(_page)
		if err != nil || r.Method != "GET" {
			app.LogError(err)
			rw.WriteHeader(http.StatusBadRequest)
			http.NotFound(rw, r)
			return
		}
		msg := app.GetGameSearchList(page, term, &rw)
		rw.Write(msg)
	})
	app.mux.Get("/games/id/{id}", func(rw http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.Atoi(_id)
		if err != nil || r.Method != "GET" {
			app.LogError(err)
			rw.WriteHeader(http.StatusBadRequest)
			http.NotFound(rw, r)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		msg := app.GetGameByID(id, &rw)
		rw.Write(msg)
	})
	app.mux.Get("/games/genre/{genre}/{page}/{option}", func(rw http.ResponseWriter, r *http.Request) {
		genre := chi.URLParam(r, "genre")
		_page := chi.URLParam(r, "page")

		option := chi.URLParam(r, "option")
		page, err := strconv.Atoi(_page)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			http.NotFound(rw, r)
			return
		}
		msg := app.GetGameGenreList(genre, page, option, &rw)
		rw.Write(msg)
	})

	app.mux.Get("/games/theme/{theme}/{page}/{option}", func(rw http.ResponseWriter, r *http.Request) {
		theme := chi.URLParam(r, "theme")
		_page := chi.URLParam(r, "page")

		option := chi.URLParam(r, "option")
		page, err := strconv.Atoi(_page)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			http.NotFound(rw, r)
			return
		}
		msg := app.GetGameThemeList(theme, page, option, &rw)
		rw.Write(msg)
	})

	app.mux.Get("/check", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			rw.Write([]byte("GET Method Only"))
			return
		}
		rw.Write([]byte("OK"))
	})
}
