package view

import (
	"encoding/json"
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

var time_layout = "2006-01-02T15:04:05.999999"

func FeedbackToView(feedback entity.Feedback) *response.FeedbackResponse {
	return &response.FeedbackResponse{
		ID:                feedback.GetID(),
		Employee_registry: feedback.GetEmployeeRegistry(),
		Content:           feedback.GetContent(),
		Answered:          feedback.GetAnswered(),
		Answer_id:         feedback.GetAnswerID(),
		Sent_at:           feedback.GetSentAt(),
	}
}

func ViewToFeedback(json_feedback []byte) []entity.Feedback {
	var feedbacks_json []response.FeedbackJSON
	var feedbacks []entity.Feedback

	json.Unmarshal(json_feedback, &feedbacks_json)

	for _, feedback := range feedbacks_json {
		t, _ := time.Parse(time_layout, feedback.Sent_at)

		f := entity.NewFeedback(feedback.Employee_registry, feedback.Content)
		f.SetID(feedback.ID)
		f.SetAnswered(feedback.Answered)
		f.SetAnswerID(feedback.Answer_id)
		f.SetSentAt(t)

		feedbacks = append(feedbacks, f)
	}

	return feedbacks
}
