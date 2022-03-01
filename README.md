# Game Dictionary API

The backend api deployed in Heroku, written in Go

## Libraries Used

- IGDB api, for fetching external data
- Go Chi Router
- Heroku

This api fetches game data from [IGDB api](https://api-docs.igdb.com/) from Twitch. IGDB api uses APICalypse.

```
APICalypse is a new language used for this api which greatly simplifies how you can query your requests compared to the url parameters used in API V2.
```

The thing I do know is that body data format in http post request should be text, not json.

Generating a access token is required client id and client secret. In api, the access token is refreshed every 24 hours. To persist access tokens, I deployed this in Heroku. 

To ensure game data is sent in a more consistent format without null properties.

The game list must have these properties.

```
id: number;
name: string;
created_at: number;
updated_at: number;
total_rating: number;
total_rating_count: number;
cover: string;
genres: GameGenre[];
themes: GameTheme[];
```

The game infomation must have these properties.

```
id: number;
name: string;
created_at: number;
updated_at: number;
total_rating: number;
total_rating_count: number;
summary: string;
storyline: string;
url: string;
cover: string;
artworks: string[];
genres: GameGenre[];
themes: GameTheme[];
videos: GameVideo[];
websites: GameWebsite[];
platforms: GamePlatform[];
```

Genre
```
id: number;
name: string;
```

Theme
```
id: number;
name: string;
```

Video
```
name: string;
video_id: string;
```

Website
```
category: string;
url: string;
```

Platform
```
name: string;
logo: string;
```

## Features

- Query games by popularity, release date, review scores, alphabetical order.
- Query game infomation by game id
  - title, summary, storyline, images, videos, themes, genres

- Filter games by themes or genres
- Search games by keywords
- Generate and refresh access tokens every 24 hours


