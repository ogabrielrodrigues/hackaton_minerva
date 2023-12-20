package instance

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database/postgres"
	handler "github.com/ogabrielrodrigues/hackaton-minerva/server/handler/employee"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/employee"
	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/employee"
)

func NewEmployeeInstance() handler.EmployeeHandler {
	db := postgres.GetDB()
	er := repository.NewEmployeeRepository(db)
	es := service.NewEmployeeService(er)
	eh := handler.NewEmployeeHandler(es)

	return eh
}
