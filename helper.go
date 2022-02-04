package main

import "fmt"

var (
	dataRequired string = "fields cover.image_id,id,name,genres.slug,genres.id,created_at,updated_at,platforms.name,platforms.platform_logo.image_id,rating,rating_count,total_rating,total_rating_count,themes.id,themes.slug;"
	notNull      string = "cover.image_id != null & id != null & name != null & genres != null & created_at != null & updated_at != null & platforms != null & rating != null & rating_count != null & total_rating != null & total_rating_count != null & release_dates != null & themes != null"
)

func categoryName(id int) string {
	switch id {
	case 1:
		return "official"
	case 2:
		return "wikia"
	case 3:
		return "wikipedia"
	case 4:
		return "facebook"
	case 5:
		return "twitter"
	case 6:
		return "twitch"
	case 8:
		return "instagram"
	case 9:
		return "youtube"
	case 10:
		return "iphone"
	case 11:
		return "ipad"
	case 12:
		return "android"
	case 13:
		return "steam"
	case 14:
		return "reddit"
	case 15:
		return "itch"
	case 16:
		return "epicgames"
	case 17:
		return "gog"
	case 18:
		return "discord"
	default:
		return "unknown"
	}
}

func toSortOption(option string) string {
	switch option {
	case "popular":
		return fmt.Sprintf(`
			sort total_rating_count desc; 
			where total_rating_count > 0 & %s 
		`, notNull)
	case "rating":
		return fmt.Sprintf(`
			sort total_rating desc; 
			where rating > 0 & %s 
		`, notNull)
	case "latest":
		return fmt.Sprintf(`
			sort created_at desc; 
			where %s 
		`, notNull)
	case "title":
		return fmt.Sprintf(`
			sort name asc; 
			where %s 
		`, notNull)
	default:
		return fmt.Sprintf(`
			sort total_rating_count desc; 
			where total_rating_count > 0 & %s 
		`, notNull)
	}
}
