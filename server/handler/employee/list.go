package handler

import (
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// List Employees godoc
// @Summary      List all employees
// @Description  List all employees of all sectors and all units
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Success      200  {object}  []response.EmployeeResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /employee/list [get]
func (eh *employeeHandler) List(w http.ResponseWriter, r *http.Request) {
	employees, err := eh.service.List()
	if err != nil {
		logger.Err("error listing employees", err)
		rest.JSON(w, err.Code, err)
		return
	}

	var employees_response []*response.EmployeeResponse

	for _, employee := range employees {
		employees_response = append(employees_response, view.EmployeeToView(employee))
	}

	rest.JSON(w, http.StatusOK, employees_response)
}
