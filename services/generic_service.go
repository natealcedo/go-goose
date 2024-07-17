package services

type GenericService interface {
	Create(body interface{}) (interface{}, error)
	GetAll() ([]interface{}, error)
	GetByID(id string) (interface{}, error)
	DeleteByID(id string) error
}
