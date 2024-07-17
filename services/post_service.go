package services

import (
	"encoding/json"
	"errors"
	"github.com/natealcedo/go-goose/models"
	"github.com/natealcedo/go-goose/repository"
	"gorm.io/gorm"
)

type PostService struct {
	PostService   repository.Repository[models.Post]
	db            *gorm.DB
	relationships []string
}

func NewPostService(postRepository repository.Repository[models.Post], db *gorm.DB) *PostService {
	return &PostService{PostService: postRepository, db: db, relationships: []string{"Comments"}}
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
	query := s.db

	// Dynamically preload relationships
	for _, relation := range s.relationships {
		query = query.Preload(relation)
	}

	var post models.Post

	// Fetch the model by ID, preloading the specified relationships
	result := query.First(&post, "id = ?", id) // model is already a pointer to a struct
	if result.Error != nil {
		// Return nil and the error if any
		return nil, result.Error
	}

	// Return the fetched model and nil error if fetch was successful
	return post, nil
}

func (s *PostService) DeleteByID(id string) error {
	err := s.PostService.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Relationships() []string {
	return s.relationships
}
