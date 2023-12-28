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

// Create Feedback godoc
// @Summary      Create new feedback
// @Description  Receive user request body to create a new feedback
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Insert your auth token" default(Bearer <Add access token here>)
// @Param        request body request.CreateFeedbackRequest true "Request Body"
// @Success      200  {object}  response.FeedbackResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/feedback [post]
func (fh *feedbackHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := request.CreateFeedbackRequest{}

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

	feedback := entity.NewFeedback(
		body.Employee_registry,
		body.Content,
	)

	feedback, err := fh.service.Create(feedback)
	if err != nil {
		logger.Err("error on create feedback", err)
		rest.JSON(w, err.Code, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/feedback/%s", feedback.GetID()))
	rest.JSON(w, http.StatusCreated, view.FeedbackToView(feedback))
}
