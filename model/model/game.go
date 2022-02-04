package model

type GameListModel struct {
	ID        *int         `json:"id"`
	Name      *string      `json:"name"`
	CreatedAt *int         `json:"created_at"`
	UpdatedAt *int         `json:"updated_at"`
	Rating    *float64     `json:"total_rating"`
	Count     *int         `json:"total_rating_count"`
	Cover     *string      `json:"cover"`
	Genres    []GenreModel `json:"genres"`
	Themes    []ThemeModel `json:"themes"`
}

type GameItemModel struct {
	ID        *int            `json:"id"`
	Name      *string         `json:"name"`
	CreatedAt *int            `json:"created_at"`
	UpdatedAt *int            `json:"updated_at"`
	Rating    *float64        `json:"total_rating"`
	Count     *int            `json:"total_rating_count"`
	Summary   *string         `json:"summary"`
	StoryLine *string         `json:"storyline"`
	URL       *string         `json:"url"`
	Cover     *string         `json:"cover"`
	Artworks  []string        `json:"artworks"`
	Genres    []GenreModel    `json:"genres"`
	Themes    []ThemeModel    `json:"themes"`
	Videos    []VideoModel    `json:"videos"`
	Websites  []WebsiteModel  `json:"websites"`
	Platforms []PlatformModel `json:"platforms"`
}

type GenreModel struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type VideoModel struct {
	Name    *string `json:"name"`
	VideoID *string `json:"video_id"`
}

type WebsiteModel struct {
	Category *string `json:"category"`
	URL      *string `json:"url"`
}

type PlatformModel struct {
	Name *string `json:"name"`
	Logo *string `json:"logo"`
}

type ThemeModel struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}
