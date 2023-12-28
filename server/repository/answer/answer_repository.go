package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

type answerRepository struct {
	database database.DB
}

type AnswerRepository interface {
	Reply(string, entity.Answer) (entity.Answer, *rest.RestErr)
	List() ([]entity.Answer, *rest.RestErr)
}

func NewAnswerRepository(database database.DB) AnswerRepository {
	return &answerRepository{database}
}
