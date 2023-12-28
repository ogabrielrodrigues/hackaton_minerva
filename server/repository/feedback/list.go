package repository

import (
	"database/sql"
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fr *feedbackRepository) List() ([]entity.Feedback, *rest.RestErr) {
	conn, err := fr.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM "tb_feedbacks"`)
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var id, employee_registry, content string
	var answered bool
	var answer_id sql.NullString
	var sent_at time.Time

	var feedbacks []entity.Feedback

	for rows.Next() {
		defer rows.Close()

		if err := rows.Scan(
			&id,
			&employee_registry,
			&content,
			&answered,
			&answer_id,
			&sent_at,
		); err != nil {
			return nil, rest.NewInternalServerErr(err.Error())
		}

		f := entity.NewFeedback(employee_registry, content)
		f.SetID(id)
		f.SetAnswered(answered)
		f.SetAnswerID(answer_id.String)
		f.SetSentAt(sent_at)

		feedbacks = append(feedbacks, f)
	}

	return feedbacks, nil
}
