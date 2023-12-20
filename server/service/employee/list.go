package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (es *employeeService) List() ([]entity.Employee, *rest.RestErr) {
	return es.repository.List()
}
