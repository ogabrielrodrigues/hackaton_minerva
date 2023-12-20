package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (es *employeeService) Create(employee entity.Employee) (entity.Employee, *rest.RestErr) {
	return es.repository.Create(employee)
}
