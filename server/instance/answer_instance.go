package instance

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database/postgres"
	handler "github.com/ogabrielrodrigues/hackaton-minerva/server/handler/answer"
	repository "github.com/ogabrielrodrigues/hackaton-minerva/server/repository/answer"
	service "github.com/ogabrielrodrigues/hackaton-minerva/server/service/answer"
)

func NewAnswerInstance() handler.AnswerHandler {
	db := postgres.GetDB()
	ar := repository.NewAnswerRepository(db)
	as := service.NewAnswerService(ar)
	ah := handler.NewAnswerHandler(as)

	return ah
}
