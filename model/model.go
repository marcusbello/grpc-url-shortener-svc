package model

import "time"

type AddRequest struct {
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	GoerUrl   string    `json:"goerUrl"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type ShowRequest struct {
	ID string `json:"ID"`
}

type ShortUrl struct {
	ID        string    `json:"ID"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	GoerUrl   string    `json:"goerUrl"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
