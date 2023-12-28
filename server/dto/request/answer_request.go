package request

type AnswerRequest struct {
	Feedback_id string `json:"feedback_id" validate:"required,uuid4"`
	Content     string `json:"content" validate:"required,min=5"`
}
