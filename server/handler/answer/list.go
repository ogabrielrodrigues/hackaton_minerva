package handler

import (
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// List answers godoc
// @Summary      List answers
// @Description  List answers of feedbacks
// @Tags         Answer
// @Produce      json
// @Param Authorization header string true "Insert your auth token" default(Bearer <Add access token here>)
// @Success      200  {object}  []response.AnswerResponse
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/answer [get]
func (ah *answerHandler) List(w http.ResponseWriter, r *http.Request) {
	answers, err := ah.service.List()
	if err != nil {
		logger.Err("error on list answers", err)
		rest.JSON(w, err.Code, err)
		return
	}

	var answers_response []response.AnswerResponse
	for _, answer := range answers {
		answers_response = append(answers_response, *view.AnswerToView(answer))
	}

	rest.JSON(w, http.StatusOK, answers_response)
}
