package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (es *employeeService) FindByRegistry(registry string) (entity.Employee, *rest.RestErr) {
	return es.repository.FindByRegistry(registry)
}
