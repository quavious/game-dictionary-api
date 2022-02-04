package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"nwlee.app/go/go-igdb-api/model"
)

func (app *App) createAccessToken() {
	apiURL := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", app.clientID, app.clientSecret)
	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		app.LogError(err)
		return
	}
	defer resp.Body.Close()
	token := new(model.Token)
	json.NewDecoder(resp.Body).Decode(token)
	app.accessToken = token.AccessToken
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}
	app := &App{
		infoLog:      log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:     log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		mux:          chi.NewMux(),
		clientID:     "",
		clientSecret: "",
	}
	app.clientID = os.Getenv("CLIENT_ID")
	app.clientSecret = os.Getenv("CLIENT_SECRET")
	app.accessToken = "5kuyhvva1rwsi0xv4gwpeeen0ff6iv"

	port := os.Getenv("PORT")
	host := func(port string) string {
		if len(port) > 0 {
			return ":" + port
		} else {
			return "localhost:8000"
		}
	}(port)

	app.mux.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			rw.Header().Set("Access-Control-Allow-Methods", "GET")

			app.LogInfo(r.URL.Path, r.Method, r.ContentLength)
			h.ServeHTTP(rw, r)
		})
	})
	app.Register()
	// app.createAccessToken()

	go func() {
		for {
			time.Sleep(time.Minute * 5)
			app.LogInfo("Check")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour * 24)
			app.LogInfo("Access Token Changed")
			app.createAccessToken()
		}
	}()

	app.LogInfo("Server running on the port ", port)
	http.ListenAndServe(host, app.mux)
}
