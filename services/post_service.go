package services

import (
	"encoding/json"
	"errors"
	"github.com/natealcedo/go-goose/models"
	"github.com/natealcedo/go-goose/repository"
)

type PostService struct {
	PostService repository.Repository[models.Post]
}

func NewPostService(postRepository repository.Repository[models.Post]) *PostService {
	return &PostService{PostService: postRepository}
}

func (s *PostService) Create(body interface{}) (interface{}, error) {
	// Convert body to JSON bytes
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Decode JSON bytes into models.Post, ignoring ID
	post := &models.Post{}
	if err := json.Unmarshal(bodyBytes, &post); err != nil {
		return nil, errors.New("invalid type")
	}

	// Assuming CreateTestTable is a method that validates and creates a TestTable entry.
	// This might need to be replaced with the correct method call to create a TestTable entry.
	return s.PostService.Create(*post)
}

func (s *PostService) GetAll() ([]interface{}, error) {
	posts, err := s.PostService.GetAll()
	if err != nil {
		return nil, err
	}
	var interfaceSlice []interface{} = make([]interface{}, len(posts))
	for i, d := range posts {
		interfaceSlice[i] = d
	}
	return interfaceSlice, nil
}

func (s *PostService) GetByID(id string) (interface{}, error) {
	row, err := s.PostService.GetByID(id)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (s *PostService) DeleteByID(id string) error {
	err := s.PostService.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
