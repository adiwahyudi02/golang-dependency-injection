package cases

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isErorr bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isErorr,
	}
}


type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed to create service")
	}
	return &SimpleService{repository}, nil
}