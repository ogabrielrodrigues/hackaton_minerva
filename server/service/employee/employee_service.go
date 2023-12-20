package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/employee"
)

type employeeService struct {
	repository repository.EmployeeRepository
}

type EmployeeService interface {
	Create(employee entity.Employee) (entity.Employee, *rest.RestErr)
	List() ([]entity.Employee, *rest.RestErr)
	FindByRegistry(string) (entity.Employee, *rest.RestErr)
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &employeeService{repository}
}
