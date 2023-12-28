package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (as *answerService) Reply(feedback_id string, answer entity.Answer) (entity.Answer, *rest.RestErr) {
	return as.repository.Reply(feedback_id, answer)
}
