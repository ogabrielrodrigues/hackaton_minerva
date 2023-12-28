package service

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fs *feedbackService) List() ([]entity.Feedback, *rest.RestErr) {
	return fs.repository.List()
}
