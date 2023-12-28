package service

import (
	"net/url"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/feedback"
)

type feedbackService struct {
	repository repository.FeedbackRepository
}

type FeedbackService interface {
	Create(entity.Feedback) (entity.Feedback, *rest.RestErr)
	List() ([]entity.Feedback, *rest.RestErr)
	Filter(url.Values) ([]entity.Feedback, *rest.RestErr)
}

func NewFeedbackService(repository repository.FeedbackRepository) FeedbackService {
	return &feedbackService{repository}
}
