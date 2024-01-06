package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/validation"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// Find Answer godoc
// @Summary      Find one answer
// @Description  Receive id on url params to find one answer
// @Tags         Answer
// @Produce      json
// @Param        id path string true "Request Param"
// @Success      200  {object}  response.AnswerResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/answer/{id} [get]
func (ah *answerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := validation.Validate.Var(id, "uuid4"); err != nil {
		logger.Err("param id not provided", err)
		rest_err := validation.ValidateVar("id", err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	answer, err := ah.service.FindByID(id)
	if err != nil {
		logger.Err("error on found answer", err)
		rest.JSON(w, err.Code, err)
		return
	}

	rest.JSON(w, http.StatusOK, view.AnswerToView(answer))
}
