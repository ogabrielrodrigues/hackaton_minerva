package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/validation"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/request"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// Reply Feedback godoc
// @Summary      Reply feedback
// @Description  Reply one open feedback
// @Tags         Answer
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Insert your auth token" default(Bearer <Add access token here>)
// @Param        request body request.AnswerRequest true "Request Body"
// @Success      200  {object}  response.AnswerResponse
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/answer/reply [post]
func (ah *answerHandler) Reply(w http.ResponseWriter, r *http.Request) {
	body := request.AnswerRequest{}

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

	answer := entity.NewAnswer(r.Header.Get("X-Authenticated-Employee"), body.Content)
	answer, err := ah.service.Reply(body.Feedback_id, answer)
	if err != nil {
		logger.Err("error on reply feedback", err)
		rest.JSON(w, err.Code, err)
		return
	}

	rest.JSON(w, http.StatusOK, view.AnswerToView(answer))
}
