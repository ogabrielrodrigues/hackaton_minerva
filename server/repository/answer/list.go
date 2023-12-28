package repository

import (
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (ar *answerRepository) List() ([]entity.Answer, *rest.RestErr) {
	conn, err := ar.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(
		`SELECT
			a.*
		FROM "tb_feedbacks" f
		INNER JOIN "tb_answers" a
		ON f.answer_id = a.id`,
	)
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var id, answered_by, content string
	var answered_at time.Time
	var answers []entity.Answer

	for rows.Next() {
		defer rows.Close()

		if err := rows.Scan(
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

		answers = append(answers, answer)
	}

	return answers, nil
}
