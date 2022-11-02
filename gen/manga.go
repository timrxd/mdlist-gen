package gen

type GetMangaResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     Manga  `json:"data"`
}

type Manga struct {
	Id            string          `json:"id"`
	Type          string          `json:"type"`
	Attributes    MangaAttributes `json:"attributes"`
	Relationships []Relationship  `json:"relationships"`
}

type MangaAttributes struct {
	Title         ListLangs `json:"title"`
	IsLocked      bool      `json:"isLocked"`
	Status        string    `json:"status"`
	Year          int       `json:"year"`
	ContentRating string    `json:"contentRating"`
	Tags          []Tag     `json:"tags"`
}

// Hardcoded list of languages that could be listed
type ListLangs struct {
	En string `json:"en"`
}

type Tag struct {
	Id            string        `json:"id"`
	Type          string        `json:"type"`
	TagAttributes TagAttributes `json:"attributes"`
}

type TagAttributes struct {
	Name    ListLangs `json:"name"`
	Group   string    `json:"group"`
	Version int       `json:"version"`
}

type Relationship struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Related string `json:"related"`
}
