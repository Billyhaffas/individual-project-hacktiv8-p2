package externalapi

import (
	"encoding/json"
	"fmt"
	"individual-project-hacktiv8-p2/internal/domain"
	"individual-project-hacktiv8-p2/internal/model/googlebooks"
	"io"
	"log"
	"net/http"
)

type googleBooksRepo struct {
	client *http.Client
}

func NewGoogleBooksRepo(client *http.Client) domain.GoogleBooksRepository {
	return &googleBooksRepo{client: client}
}
func (googleRepo *googleBooksRepo) GetBookByCategory(category string) (*googlebooks.GoogleBooksResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://www.googleapis.com/books/v1/volumes",
		nil,
	)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", "subject:"+category)
	q.Set("maxResults", "10")
	req.URL.RawQuery = q.Encode()

	resp, err := googleRepo.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google books api error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Println(string(body)) // ✅ debug OK

	var result googlebooks.GoogleBooksResponse
	// log.Println("Total items:", len(result))

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	log.Println(result)
	return &result, nil
}
