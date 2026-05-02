package domain

import "individual-project-hacktiv8-p2/internal/model/googlebooks"

type GoogleBooksRepository interface {
	GetBookByCategory(category string) (*googlebooks.GoogleBooksResponse, error)
}
