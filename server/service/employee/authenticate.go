package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (es *employeeService) Authenticate(employee entity.Employee) (string, *rest.RestErr) {
	employee, err := es.repository.Authenticate(employee)
	if err != nil {
		return "", err
	}

	token, err := employee.Authenticate()
	if err != nil {
		return "", nil
	}

	return token, nil
}
