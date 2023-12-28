package handler

import (
	"net/http"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

// Filter Feedbacks godoc
// @Summary      Filter feedbacks
// @Description  Filter feedback by unit, sector or both
// @Tags         Feedback
// @Produce      json
// @Param        unit    query     string  false  "unit search by query param unit"
// @Param        sector    query     string  false  "sector search by query param sector"
// @Param Authorization header string true "Insert your auth token" default(Bearer <Add access token here>)
// @Success      200  {object}  []response.FeedbackResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      401  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /api/v1/feedback/filter [get]
func (fh *feedbackHandler) Filter(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()
	if !filters.Has("unit") && !filters.Has("sector") {
		err := rest.NewBadRequestErr("unit or sector not provided", nil)
		logger.Err("unit or sector not provided", err)
		rest.JSON(w, err.Code, err)
		return
	}

	feedbacks, err := fh.service.Filter(filters)
	if err != nil {
		logger.Err("error on filter feedbacks", err)
		rest.JSON(w, err.Code, err)
		return
	}

	var feedbacks_response []response.FeedbackResponse
	for _, feedback := range feedbacks {
		feedbacks_response = append(feedbacks_response, *view.FeedbackToView(feedback))
	}

	rest.JSON(w, http.StatusOK, feedbacks_response)
}
