package book

import "individual-project-hacktiv8-p2/internal/model/googlebooks"

type BookRespon struct {
	BookId      int     `json:"book_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
type Getbook struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    *GetBookRespon `json:"data"`
}

type GetbookById struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    *BookRespon `json:"data"`
}

type RentBookRespon struct {
	Name       string  `json:"name"`
	Cartegory  string  `json:"category"`
	RentalCost float32 `json:"rental_cost"`
}

type GetBookRespon struct {
	BookId            int                              `json:"book_id"`
	RentalCost        float32                          `json:"rental_cost"`
	StockAvailability int                              `json:"stock_availability"`
	Category          string                           `json:"category"`
	RekomendationBook *googlebooks.GoogleBooksResponse `json:"book_recomendation"`
}
