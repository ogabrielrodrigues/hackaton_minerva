package repository

import (
	"net/url"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

type feedbackRepository struct {
	database database.DB
}

type FeedbackRepository interface {
	Create(entity.Feedback) (entity.Feedback, *rest.RestErr)
	List() ([]entity.Feedback, *rest.RestErr)
	Filter(url.Values) ([]entity.Feedback, *rest.RestErr)
}

func NewFeedbackRepository(database database.DB) FeedbackRepository {
	return &feedbackRepository{database}
}
