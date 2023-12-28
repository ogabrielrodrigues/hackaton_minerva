package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/validation"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/request"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

// Authenticate godoc
// @Summary      Authenticate employee
// @Description  Authenticate employee to access private routes
// @Tags         Employee
// @Accept       json
// @Header       200 {string} Authentication "token"
// @Param        request body request.AuthenticateEmployeeRequest true "Request Body"
// @Success      200
// @Failure      400  {object}  rest.RestErr
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/employee/auth [post]
func (eh *employeeHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	body := request.AuthenticateEmployeeRequest{}

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
		"",
		body.Email,
		"",
		"",
		body.Password,
		false,
	)

	token, err := eh.service.Authenticate(employee)
	if err != nil {
		logger.Err("error on authenticate employee", err)
		rest.JSON(w, err.Code, err)
		return
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
