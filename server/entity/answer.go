package entity

import "time"

type answer struct {
	id          string
	answered_by string
	content     string
	answered_at time.Time
}

type Answer interface {
	SetID(string)
	GetID() string
	SetAnsweredBy(answered_by string)
	GetAnsweredBy() string
	SetContent(content string)
	GetContent() string
	SetAnsweredAt(answered_at time.Time)
	GetAnsweredAt() time.Time
}

func NewAnswer(answered_by, content string) Answer {
	id := ""
	answered_at := time.Now()
	return &answer{
		id,
		answered_by,
		content,
		answered_at,
	}
}

func (a *answer) SetID(id string) {
	a.id = id
}

func (a *answer) GetID() string {
	return a.id
}

func (a *answer) SetAnsweredBy(answered_by string) {
	a.answered_by = answered_by
}

func (a *answer) GetAnsweredBy() string {
	return a.answered_by
}

func (a *answer) SetContent(content string) {
	a.content = content
}

func (a *answer) GetContent() string {
	return a.content
}

func (a *answer) SetAnsweredAt(answered_at time.Time) {
	a.answered_at = answered_at
}

func (a *answer) GetAnsweredAt() time.Time {
	return a.answered_at
}
