package services

import (
	"encoding/json"
	"errors"
	"github.com/natealcedo/go-goose/models"
	"github.com/natealcedo/go-goose/repository"
	"gorm.io/gorm"
)

type CommentService struct {
	CommentService repository.Repository[models.Comment]
	db             *gorm.DB
	relationships  []string
}

func NewCommentService(postRepository repository.Repository[models.Comment], db *gorm.DB) *CommentService {
	return &CommentService{CommentService: postRepository, db: db, relationships: []string{"Comments"}}
}

func (s *CommentService) Create(body interface{}) (interface{}, error) {
	// Convert body to JSON bytes
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Decode JSON bytes into models.Post, ignoring ID
	comment := &models.Comment{}
	if err := json.Unmarshal(bodyBytes, &comment); err != nil {
		return nil, errors.New("invalid type")
	}

	// Assuming CreateTestTable is a method that validates and creates a TestTable entry.
	// This might need to be replaced with the correct method call to create a TestTable entry.
	return s.CommentService.Create(*comment)
}

func (s *CommentService) GetAll() ([]interface{}, error) {
	comments, err := s.CommentService.GetAll()
	if err != nil {
		return nil, err
	}

	var interfaceSlice []interface{} = make([]interface{}, len(comments))
	for i, comment := range comments {
		// Assuming the need to set IncludePostID to true for each comment
		comment.IncludePostID = true
		interfaceSlice[i] = comment
	}

	return interfaceSlice, nil
}

func (s *CommentService) GetByID(id string) (interface{}, error) {
	query := s.db

	// Dynamically preload relationships
	query = query.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, post_id, content, created_at, updated_at")
	})

	comment := &models.Comment{
		IncludePostID: true,
	}

	// Fetch the model by ID, preloading the specified relationships
	result := query.First(&comment, "id = ?", id) // model is already a pointer to a struct
	if result.Error != nil {
		// Return nil and the error if any
		return nil, result.Error
	}

	return comment, nil
}

func (s *CommentService) DeleteByID(id string) error {
	var comment models.Comment
	result := s.db.Where("id = ?", id).Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
