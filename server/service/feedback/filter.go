package service

import (
	"net/url"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fs *feedbackService) Filter(filters url.Values) ([]entity.Feedback, *rest.RestErr) {
	return fs.repository.Filter(filters)
}
