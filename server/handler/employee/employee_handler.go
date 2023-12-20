package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/employee"
)

type employeeHandler struct {
	service service.EmployeeService
}

type EmployeeHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	FindByRegistry(w http.ResponseWriter, r *http.Request)
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}
