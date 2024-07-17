package services

import (
	"encoding/json"
	"errors"
	"github.com/natealcedo/go-goose/models"
	"github.com/natealcedo/go-goose/repository"
)

type TestTableService struct {
	testTableRepository repository.Repository[models.TestTable]
}

func NewTestTableService(testTableRepository repository.Repository[models.TestTable]) *TestTableService {
	return &TestTableService{testTableRepository: testTableRepository}
}

func (s *TestTableService) Create(body interface{}) error {
	// Convert body to JSON bytes
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Decode JSON bytes into models.TestTable, ignoring ID
	var testTable models.TestTable
	if err := json.Unmarshal(bodyBytes, &testTable); err != nil {
		return errors.New("invalid type")
	}

	// Assuming CreateTestTable is a method that validates and creates a TestTable entry.
	// This might need to be replaced with the correct method call to create a TestTable entry.
	return s.testTableRepository.Create(testTable)
}

func (s *TestTableService) GetAll() ([]interface{}, error) {
	testTables, err := s.testTableRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var interfaceSlice []interface{} = make([]interface{}, len(testTables))
	for i, d := range testTables {
		interfaceSlice[i] = d
	}
	return interfaceSlice, nil
}

func (s *TestTableService) GetByID(id string) (interface{}, error) {
	row, err := s.testTableRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return row, nil
}
