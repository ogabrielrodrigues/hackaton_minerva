package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/feedback"
)

type feedbackHandler struct {
	service service.FeedbackService
}

type FeedbackHandler interface {
	Create(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	Filter(http.ResponseWriter, *http.Request)
}

func NewFeedbackHandler(service service.FeedbackService) FeedbackHandler {
	return &feedbackHandler{service}
}
