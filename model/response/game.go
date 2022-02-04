package response

type GameItemResponse struct {
	ID        *int               `json:"id"`
	Name      *string            `json:"name"`
	CreatedAt *int               `json:"created_at"`
	UpdatedAt *int               `json:"updated_at"`
	Rating    *float64           `json:"total_rating"`
	Count     *int               `json:"total_rating_count"`
	Summary   *string            `json:"summary"`
	StoryLine *string            `json:"storyline"`
	URL       *string            `json:"url"`
	Cover     *CoverResponse     `json:"cover"`
	Artworks  []ArtworkResponse  `json:"artworks"`
	Genres    []GenreResponse    `json:"genres"`
	Platforms []PlatformResponse `json:"platforms"`
	Themes    []ThemeResponse    `json:"themes"`
	Videos    []VideoResponse    `json:"videos"`
	Websites  []WebsiteResponse  `json:"websites"`
}

type SearchItemResponse struct {
	Game struct {
		ID        *int            `json:"id"`
		Name      *string         `json:"name"`
		CreatedAt *int            `json:"created_at"`
		UpdatedAt *int            `json:"updated_at"`
		Rating    *float64        `json:"total_rating"`
		Count     *int            `json:"total_rating_count"`
		Cover     *CoverResponse  `json:"cover"`
		Genres    []GenreResponse `json:"genres"`
		Themes    []ThemeResponse `json:"themes"`
	} `json:"game"`
}

type ArtworkResponse struct {
	ID      *int    `json:"id"`
	ImageID *string `json:"image_id"`
}

type CoverResponse struct {
	ID      *int    `json:"id"`
	ImageID *string `json:"image_id"`
}

type GenreResponse struct {
	ID   *int    `json:"id"`
	Name *string `json:"slug"`
}

type PlatformResponse struct {
	Logo *struct {
		ImageID *string `json:"image_id"`
	} `json:"platform_logo"`
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type ThemeResponse struct {
	ID   *int    `json:"id"`
	Name *string `json:"slug"`
}

type VideoResponse struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	VideoID *string `json:"video_id"`
}

type WebsiteResponse struct {
	Category *int    `json:"category"`
	ID       *int    `json:"id"`
	URL      *string `json:"url"`
}
