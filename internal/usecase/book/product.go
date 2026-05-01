package book

import (
	"context"
)

type bookUseCase struct {
	bookRepository domain.bookRepository
}

func NewbookUseCase(bookRepo domain.bookRepository) domain.bookUseCase {
	return &bookUseCase{bookRepository: bookRepo}
}
func (bookUC *bookUseCase) GetAllbook(ctx context.Context) (*[]book.bookRespon, error) {
	books, err := bookUC.bookRepository.Getbook()
	if err != nil {
		return nil, err
	}

	var responses []book.bookRespon

	for _, p := range *books {
		responses = append(responses, book.bookRespon{
			bookId:      p.bookId,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return &responses, nil
}

func (bookUC *bookUseCase) GetbookById(ctx context.Context, bookID int) (*book.bookRespon, error) {
	getbook, err := bookUC.bookRepository.GetbookById(bookID)
	if err != nil {
		return nil, err
	}

	responses := book.bookRespon{
		bookId:      getbook.bookId,
		Name:        getbook.Name,
		Description: getbook.Description,
		Price:       getbook.Price,
	}
	return &responses, nil
}
