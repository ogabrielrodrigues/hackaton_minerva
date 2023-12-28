package view

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func AnswerToView(answer entity.Answer) *response.AnswerResponse {
	return &response.AnswerResponse{
		ID:          answer.GetID(),
		Answered_by: answer.GetAnsweredBy(),
		Content:     answer.GetContent(),
		Answered_at: answer.GetAnsweredAt(),
	}
}
