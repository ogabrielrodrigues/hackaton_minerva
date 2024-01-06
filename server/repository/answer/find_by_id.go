package repository

import (
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (ar *answerRepository) FindByID(id string) (entity.Answer, *rest.RestErr) {
	conn, err := ar.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`SELECT
			a.*
		FROM "tb_feedbacks" f
		INNER JOIN "tb_answers" a
		ON f.answer_id = a.id
		WHERE a.id = $1`,
		id,
	)

	var answered_by, content string
	var answered_at time.Time

	if err := row.Scan(
		&id,
		&answered_by,
		&content,
		&answered_at,
	); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	answer := entity.NewAnswer(answered_by, content)
	answer.SetID(id)
	answer.SetAnsweredAt(answered_at)

	return answer, nil
}
