package model

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	PriceIn  float64   `json:"price_in"`
	PriceOut float64   `json:"price_out"`
	Quantity int       `json:"quantity"`
}
