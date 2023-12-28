package response

import "time"

type AnswerResponse struct {
	ID          string    `json:"id"`
	Answered_by string    `json:"answered_by"`
	Content     string    `json:"content"`
	Answered_at time.Time `json:"answered_at"`
}
