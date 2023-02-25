package models

type Model struct {
	ID       int64  `json:"id"`
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortURl"`
}
