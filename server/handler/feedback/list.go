package handler

import (
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// List Feedback godoc
// @Summary      List feedbacks
// @Description  List all feedbacks of all employees
// @Tags         Feedback
// @Produce      json
// @Param Authorization header string true "Insert your auth token" default(Bearer <Add access token here>)
// @Success      200  {object}  []response.FeedbackResponse
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/feedback [get]
func (fh *feedbackHandler) List(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := fh.service.List()
	if err != nil {
		logger.Err("error on list feedbacks", err)
		rest.JSON(w, err.Code, err)
		return
	}

	var feedbacks_response []response.FeedbackResponse
	for _, feedback := range feedbacks {
		feedbacks_response = append(feedbacks_response, *view.FeedbackToView(feedback))
	}

	rest.JSON(w, http.StatusOK, feedbacks_response)
}
