package repository

import "go-unit-test/entity"

type CategoryRepository interface { // **CONTRACT
	FindById(id string) *entity.Category
}
