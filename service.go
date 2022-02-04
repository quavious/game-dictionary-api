package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"nwlee.app/go/go-igdb-api/model/model"
	"nwlee.app/go/go-igdb-api/model/response"
)

type Request struct {
	*http.Request
}

const (
	size = 20
)

func (app *App) injectToken(req *http.Request) {
	req.Header["Client-ID"] = []string{app.clientID}
	req.Header["Authorization"] = []string{"Bearer " + app.accessToken}
	req.Header["Accept"] = []string{"*/*"}
	req.Header["Content-Type"] = []string{"text/plain"}
}

func stringify(response []response.GameItemResponse) []model.GameListModel {
	gameModel := []model.GameListModel{}
	for _, game := range response {
		genres := []model.GenreModel{}
		themes := []model.ThemeModel{}
		for _, genre := range game.Genres {
			genres = append(genres, model.GenreModel{
				ID:   genre.ID,
				Name: genre.Name,
			})
		}
		for _, theme := range game.Themes {
			themes = append(themes, model.ThemeModel{
				ID:   theme.ID,
				Name: theme.Name,
			})
		}
		model := &model.GameListModel{
			ID:        game.ID,
			Name:      game.Name,
			CreatedAt: game.CreatedAt,
			UpdatedAt: game.UpdatedAt,
			Rating:    game.Rating,
			Count:     game.Count,
			Cover:     game.Cover.ImageID,
			Genres:    genres,
			Themes:    themes,
		}
		gameModel = append(gameModel, *model)
	}
	return gameModel
}

func (app *App) GetGameByID(id int, rw *http.ResponseWriter) []byte {
	buf := bytes.NewBuffer([]byte(fmt.Sprintf("fields artworks.image_id,cover.image_id,id,name,genres.slug,genres.id,created_at,updated_at,platforms.name,platforms.platform_logo.image_id,rating,rating_count,summary,storyline,total_rating,total_rating_count,themes.id,themes.slug,url,videos.name,videos.video_id,websites.category,websites.url; where id = %d;", id)))
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", buf)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Bad Request"))
		return nil
	}
	app.injectToken(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	defer resp.Body.Close()
	gameResponse := []response.GameItemResponse{}
	json.NewDecoder(resp.Body).Decode(&gameResponse)
	gameModel := []model.GameItemModel{}

	for _, game := range gameResponse {
		genres := []model.GenreModel{}
		for _, genre := range game.Genres {
			genres = append(genres, model.GenreModel{
				ID:   genre.ID,
				Name: genre.Name,
			})
		}
		artworks := []string{}
		themes := []model.ThemeModel{}
		videos := []model.VideoModel{}
		websites := []model.WebsiteModel{}
		platforms := []model.PlatformModel{}
		for _, artwork := range game.Artworks {
			artworks = append(artworks, *artwork.ImageID)
		}
		for _, theme := range game.Themes {
			themes = append(themes, model.ThemeModel{ID: theme.ID, Name: theme.Name})
		}
		for _, video := range game.Videos {
			videos = append(videos, model.VideoModel{Name: video.Name, VideoID: video.VideoID})
		}
		for _, website := range game.Websites {
			category := categoryName(*website.Category)
			websites = append(websites, model.WebsiteModel{Category: &category, URL: website.URL})
		}
		for _, platform := range game.Platforms {
			platforms = append(platforms, model.PlatformModel{Name: platform.Name, Logo: platform.Logo.ImageID})
		}
		model := model.GameItemModel{
			ID:        game.ID,
			Name:      game.Name,
			CreatedAt: game.CreatedAt,
			UpdatedAt: game.UpdatedAt,
			Rating:    game.Rating,
			Count:     game.Count,
			Summary:   game.Summary,
			StoryLine: game.StoryLine,
			URL:       game.URL,
			Cover:     game.Cover.ImageID,
			Artworks:  artworks,
			Genres:    genres,
			Themes:    themes,
			Videos:    videos,
			Websites:  websites,
			Platforms: platforms,
		}
		gameModel = append(gameModel, model)
	}
	msg, _ := json.Marshal(gameModel)
	return msg
}

func (app *App) GetGameList(page int, option string, rw *http.ResponseWriter) []byte {
	(*rw).Header().Set("Content-Type", "application/json")
	option = toSortOption(option)
	buf := bytes.NewBuffer([]byte(fmt.Sprintf(`%s limit %d; offset %d; %s;`, dataRequired, size, size*(page-1), option)))
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", buf)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Bad Request"))
		return nil
	}
	app.injectToken(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	defer resp.Body.Close()
	gameResponse := []response.GameItemResponse{}
	json.NewDecoder(resp.Body).Decode(&gameResponse)

	gameModel := stringify(gameResponse)
	msg, _ := json.Marshal(gameModel)
	return msg
}

func (app *App) GetGameSearchList(page int, term string, rw *http.ResponseWriter) []byte {
	(*rw).Header().Set("Content-Type", "application/json")
	buf := bytes.NewReader([]byte(fmt.Sprintf(`
		fields game.cover.image_id,game.id,game.name,game.genres.slug,game.genres.id,game.created_at,game.updated_at,game.platforms.name,game.platforms.platform_logo.image_id,game.rating,game.rating_count,game.total_rating,game.total_rating_count,game.themes.id,game.themes.slug; 
		search "%s"; 
		limit %d; 
		offset %d;
		where game != null & game.cover.image_id != null & game.id != null & game.name != null & game.genres.name != null & game.created_at != null & game.updated_at != null & game.platforms.name != null & game.platforms.platform_logo.image_id != null & game.rating != null & game.rating_count != null & game.total_rating != null & game.total_rating_count != null; 
	`, term, size, (page-1)*size)))
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/search", buf)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	app.injectToken(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	defer resp.Body.Close()
	gameResponse := []response.SearchItemResponse{}
	json.NewDecoder(resp.Body).Decode(&gameResponse)

	gameModel := []model.GameListModel{}
	for _, resp := range gameResponse {
		game := resp.Game
		genres := []model.GenreModel{}
		themes := []model.ThemeModel{}
		for _, genre := range game.Genres {
			genres = append(genres, model.GenreModel{
				ID:   genre.ID,
				Name: genre.Name,
			})
		}
		for _, theme := range game.Themes {
			themes = append(themes, model.ThemeModel{
				ID:   theme.ID,
				Name: theme.Name,
			})
		}

		model := &model.GameListModel{
			ID:        game.ID,
			Name:      game.Name,
			CreatedAt: game.CreatedAt,
			UpdatedAt: game.UpdatedAt,
			Rating:    game.Rating,
			Count:     game.Count,
			Cover:     game.Cover.ImageID,
			Genres:    genres,
			Themes:    themes,
		}
		gameModel = append(gameModel, *model)
	}

	msg, _ := json.Marshal(gameModel)
	(*rw).Write(msg)
	return nil
}

func (app *App) GetGameGenreList(genre string, page int, option string, rw *http.ResponseWriter) []byte {
	(*rw).Header().Set("Content-Type", "application/json")
	genre = url.PathEscape(genre)
	option = toSortOption(option)
	buf := bytes.NewBuffer([]byte(fmt.Sprintf(`%s limit %d; offset %d; %s & genres.slug = "%s";`, dataRequired, size, size*(page-1), option, genre)))
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", buf)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Bad Request"))
		return nil
	}
	app.injectToken(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	defer resp.Body.Close()
	gameResponse := []response.GameItemResponse{}
	json.NewDecoder(resp.Body).Decode(&gameResponse)

	gameModel := stringify(gameResponse)
	msg, _ := json.Marshal(gameModel)
	return msg
}

func (app *App) GetGameThemeList(theme string, page int, option string, rw *http.ResponseWriter) []byte {
	(*rw).Header().Set("Content-Type", "application/json")
	theme = url.PathEscape(theme)
	option = toSortOption(option)
	buf := bytes.NewBuffer([]byte(fmt.Sprintf(`%s limit %d; offset %d; %s & themes.slug = "%s";`, dataRequired, size, size*(page-1), option, theme)))
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", buf)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Bad Request"))
		return nil
	}
	app.injectToken(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.LogError(err)
		(*rw).Write([]byte("Server Error"))
		return nil
	}
	defer resp.Body.Close()
	gameResponse := []response.GameItemResponse{}
	json.NewDecoder(resp.Body).Decode(&gameResponse)

	gameModel := stringify(gameResponse)
	msg, _ := json.Marshal(gameModel)
	return msg
}
