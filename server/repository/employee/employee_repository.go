package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

type employeeRepository struct {
	database database.DB
}

type EmployeeRepository interface {
	Create(employee entity.Employee) (entity.Employee, *rest.RestErr)
	List() ([]entity.Employee, *rest.RestErr)
	FindByRegistry(string) (entity.Employee, *rest.RestErr)
}

func NewEmployeeRepository(database database.DB) EmployeeRepository {
	return &employeeRepository{database}
}
