package model

type Product struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}
