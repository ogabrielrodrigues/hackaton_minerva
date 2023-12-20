package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/validation"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// Find Employee godoc
// @Summary      Find one employee
// @Description  Receive registry on url params to find one employee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        registry path string true "Request Param"
// @Success      200  {object}  response.EmployeeResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      404  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /employee/{registry} [get]
func (eh *employeeHandler) FindByRegistry(w http.ResponseWriter, r *http.Request) {
	registry := chi.URLParam(r, "registry")

	if err := validation.Validate.Var(registry, "uuid4"); err != nil {
		logger.Err("param registry not provided", err)
		rest_err := validation.ValidateVar("registry", err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	employee, err := eh.service.FindByRegistry(registry)
	if err != nil {
		logger.Err("error on found employee", err)
		rest.JSON(w, err.Code, err)
		return
	}

	rest.JSON(w, http.StatusOK, view.EmployeeToView(employee))
}
