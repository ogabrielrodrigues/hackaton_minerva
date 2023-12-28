package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/answer"
)

type answerService struct {
	repository repository.AnswerRepository
}

type AnswerService interface {
	Reply(string, entity.Answer) (entity.Answer, *rest.RestErr)
	List() ([]entity.Answer, *rest.RestErr)
}

func NewAnswerService(repository repository.AnswerRepository) AnswerService {
	return &answerService{repository}
}
