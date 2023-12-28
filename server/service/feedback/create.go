package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fs *feedbackService) Create(feedback entity.Feedback) (entity.Feedback, *rest.RestErr) {
	return fs.repository.Create(feedback)
}
