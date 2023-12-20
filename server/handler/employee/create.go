package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/validation"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/request"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// Create Employee godoc
// @Summary      Create new employee
// @Description  Receive user request body to create a new employee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        request body request.CreateEmployeeRequest true "Request Body"
// @Success      200  {object}  response.EmployeeResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /employee [post]
func (eh *employeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := request.CreateEmployeeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Err("error decoding request body", err)
		rest_err := rest.NewInternalServerErr("error decoding request body")
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	if err := validation.Validate.Struct(body); err != nil {
		logger.Err("error on request validation", err)
		rest_err := validation.ValidateErr(err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	employee := entity.NewEmployee(
		body.Name,
		body.Email,
		body.Sector,
		body.Unit,
		body.Administrator,
	)

	employee, err := eh.service.Create(employee)
	if err != nil {
		logger.Err("error on create employee", err)
		rest.JSON(w, err.Code, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/employee/%s", employee.GetRegistry()))
	rest.JSON(w, http.StatusCreated, view.EmployeeToView(employee))
}
