package repository

import (
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (ar *answerRepository) Reply(feedback_id string, answer entity.Answer) (entity.Answer, *rest.RestErr) {
	conn, err := ar.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	var answer_id string
	var answered_at time.Time

	row := conn.QueryRow(
		`INSERT INTO "tb_answers" (answered_by, content)
			(SELECT registry, $1
			FROM "tb_employees"
			WHERE registry = $2 AND administrator = TRUE)
		RETURNING id, answered_at;`,
		answer.GetContent(), answer.GetAnsweredBy())

	if err := row.Scan(&answer_id, &answered_at); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	_, err = conn.Exec(
		`UPDATE "tb_feedbacks"
			SET answered = TRUE,
			answer_id = $1
			WHERE id = $2`, answer_id, feedback_id)

	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	answer.SetID(answer_id)
	answer.SetAnsweredAt(answered_at)

	return answer, nil
}
