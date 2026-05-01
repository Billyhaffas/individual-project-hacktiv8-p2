package book

type BookRespon struct {
	BookId      int     `json:"book_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
type Getbook struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    *[]BookRespon `json:"data"`
}

type GetbookById struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    *BookRespon `json:"data"`
}
