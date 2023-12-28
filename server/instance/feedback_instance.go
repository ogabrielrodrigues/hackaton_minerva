package instance

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database/postgres"
	handler "github.com/ogabrielrodrigues/hackaton-minerva/server/handler/feedback"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/feedback"
	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/feedback"
)

func NewFeedbackInstance() handler.FeedbackHandler {
	db := postgres.GetDB()
	fr := repository.NewFeedbackRepository(db)
	fs := service.NewFeedbackService(fr)
	fh := handler.NewFeedbackHandler(fs)

	return fh
}
