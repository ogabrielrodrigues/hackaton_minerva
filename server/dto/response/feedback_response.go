package response

import "time"

type FeedbackResponse struct {
	ID                string    `json:"id"`
	Employee_registry string    `json:"employee_registry"`
	Content           string    `json:"content"`
	Answered          bool      `json:"answered"`
	Answer_id         string    `json:"answer_id,omitempty"`
	Sent_at           time.Time `json:"sent_at"`
}

type FeedbackJSON struct {
	ID                string `json:"id"`
	Employee_registry string `json:"employee_registry"`
	Content           string `json:"content"`
	Answered          bool   `json:"answered"`
	Answer_id         string `json:"answer_id,omitempty"`
	Sent_at           string `json:"sent_at"`
}
