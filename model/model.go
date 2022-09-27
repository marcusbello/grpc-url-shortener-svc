package model

import "time"

type AddRequest struct {
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type ShortUrl struct {
	ID        string    `json:"ID"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
