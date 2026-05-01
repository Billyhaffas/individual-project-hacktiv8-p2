package book

type Book struct {
	bookId            int
	Name              string
	StockAvailability int
	RentalCost        string
	Cartegory         string
}
type BookRent struct {
	BookId            int
	RentalCost        float32
	StockAvailability int
}
