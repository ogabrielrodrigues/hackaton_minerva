package entity

import "time"

type feedback struct {
	id                string
	employee_registry string
	content           string
	answered          bool
	answer_id         string
	sent_at           time.Time
	answer            Answer
}

type Feedback interface {
	SetID(id string)
	GetID() string
	SetEmployeeRegistry(employee_registry string)
	GetEmployeeRegistry() string
	SetContent(content string)
	GetContent() string
	SetAnswered(answered bool)
	GetAnswered() bool
	SetAnswerID(answer_id string)
	GetAnswerID() string
	SetSentAt(sent_at time.Time)
	GetSentAt() time.Time
	SetAnswer(answer Answer)
	GetAnswer() Answer
}

func NewFeedback(employee_registry, content string) Feedback {
	var id, answer_id = "", ""
	var answered = false
	var sent_at = time.Now()
	var answer Answer

	return &feedback{
		id,
		employee_registry,
		content,
		answered,
		answer_id,
		sent_at,
		answer,
	}
}

func (f *feedback) SetID(id string) {
	f.id = id
}

func (f *feedback) GetID() string {
	return f.id
}

func (f *feedback) SetEmployeeRegistry(employee_registry string) {
	f.employee_registry = employee_registry
}

func (f *feedback) GetEmployeeRegistry() string {
	return f.employee_registry
}

func (f *feedback) SetContent(content string) {
	f.content = content
}

func (f *feedback) GetContent() string {
	return f.content
}

func (f *feedback) SetAnswered(answered bool) {
	f.answered = answered
}

func (f *feedback) GetAnswered() bool {
	return f.answered
}

func (f *feedback) SetAnswerID(answer_id string) {
	f.answer_id = answer_id
}

func (f *feedback) GetAnswerID() string {
	return f.answer_id
}

func (f *feedback) SetSentAt(sent_at time.Time) {
	f.sent_at = sent_at
}

func (f *feedback) GetSentAt() time.Time {
	return f.sent_at
}

func (f *feedback) SetAnswer(answer Answer) {
	f.answer = answer
}

func (f *feedback) GetAnswer() Answer {
	return f.answer
}
