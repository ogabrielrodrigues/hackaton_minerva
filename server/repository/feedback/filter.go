package repository

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fr *feedbackRepository) Filter(filters url.Values) ([]entity.Feedback, *rest.RestErr) {
	conn, err := fr.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	query :=
		`SELECT 
		f.*
	FROM "tb_feedbacks" f
	JOIN "tb_employees" e
	ON e.registry = f.employee_registry`

	var rows *sql.Rows

	if filters.Has("unit") && filters.Has("sector") {
		query += ` WHERE e.unit = $1 AND e.sector = $2`
		rows, err = conn.Query(query, filters.Get("unit"), filters.Get("sector"))
	} else if filters.Has("unit") {
		query += ` WHERE e.unit = $1`
		rows, err = conn.Query(query, filters.Get("unit"))
	} else {
		query += ` WHERE e.sector = $1`
		rows, err = conn.Query(query, filters.Get("sector"))
	}

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
