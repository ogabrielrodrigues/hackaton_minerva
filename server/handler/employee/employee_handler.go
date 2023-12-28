package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/employee"
)

type employeeHandler struct {
	service service.EmployeeService
}

type EmployeeHandler interface {
	Create(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	FindByRegistry(http.ResponseWriter, *http.Request)
	Authenticate(http.ResponseWriter, *http.Request)
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}
