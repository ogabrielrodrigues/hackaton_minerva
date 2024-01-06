package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/answer"
)

type answerHandler struct {
	service service.AnswerService
}

type AnswerHandler interface {
	Reply(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	FindByID(http.ResponseWriter, *http.Request)
}

func NewAnswerHandler(service service.AnswerService) AnswerHandler {
	return &answerHandler{service}
}
