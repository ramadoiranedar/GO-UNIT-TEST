package service

import (
	"errors"
	"fmt"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil { // **FAILED
		// return category, errors.New("Category Not Found")
		fmt.Println("@@FindById RETURN NIL")
		return nil, errors.New("Category Not Found") // can use nil
	} else { // **SUCCESS
		return category, nil
	}
}
