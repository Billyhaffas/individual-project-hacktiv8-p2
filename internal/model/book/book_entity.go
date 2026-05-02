package book

type Book struct {
	BookId            int
	Name              string
	StockAvailability int
	RentalCost        string
	Category          string
}
type BookRent struct {
	BookId            int
	RentalCost        float32
	StockAvailability int
}
type GetBook struct {
	BookId            int
	RentalCost        float32
	StockAvailability int
	Category          string
}
