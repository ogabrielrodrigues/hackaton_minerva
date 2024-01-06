package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (as *answerService) FindByID(id string) (entity.Answer, *rest.RestErr) {
	return as.repository.FindByID(id)
}
