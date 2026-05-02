package book

import (
	"context"
	"fmt"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/book"
)

type bookUseCase struct {
	bookRepository   domain.BookRepository
	googleRepository domain.GoogleBooksRepository
}

func NewbookUseCase(bookRepo domain.BookRepository, googleRepo domain.GoogleBooksRepository) domain.BookUseCase {
	return &bookUseCase{
		bookRepository:   bookRepo,
		googleRepository: googleRepo,
	}
}
func (bookUC *bookUseCase) GetbookByName(ctx context.Context, bookName string) (*book.GetBookRespon, error) {
	books, category, err := bookUC.bookRepository.GetbookByName(bookName)
	if err != nil {
		fmt.Println(category)
		return nil, err
	}
	googleBook, err := bookUC.googleRepository.GetBookByCategory("business")
	fmt.Println(googleBook)
	if err != nil {
		return nil, err
	}

	responses := book.GetBookRespon{
		BookId:            books.BookId,
		RentalCost:        books.RentalCost,
		StockAvailability: books.StockAvailability,
		Category:          books.Category,
		RekomendationBook: googleBook,
	}
	return &responses, nil
}
