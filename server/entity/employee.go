package entity

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
)

type employee struct {
	registry      string
	name          string
	email         string
	sector        string
	unit          string
	administrator bool
	password      string
	feedbacks     []Feedback
}

type Employee interface {
	SetRegistry(string)
	GetRegistry() string
	SetName(string)
	GetName() string
	SetEmail(string)
	GetEmail() string
	SetSector(string)
	GetSector() string
	SetUnit(string)
	GetUnit() string
	SetAdministrator(bool)
	GetAdministrator() bool
	SetPassword(string)
	GetPassword() string
	SetFeedbacks(feedbacks []Feedback)
	GetFeedbacks() []Feedback
	Encrypt()
	Authenticate() (string, *rest.RestErr)
}

func NewEmployee(name, email, sector, unit, password string, administrator bool) Employee {
	registry := ""
	feedbacks := []Feedback{}

	e := employee{
		registry,
		name,
		email,
		sector,
		unit,
		administrator,
		password,
		feedbacks,
	}

	e.Encrypt()
	return &e
}

func (e *employee) SetRegistry(registry string) {
	e.registry = registry
}

func (e *employee) GetRegistry() string {
	return e.registry
}

func (e *employee) SetName(name string) {
	e.name = name
}

func (e *employee) GetName() string {
	return e.name
}

func (e *employee) SetEmail(email string) {
	e.email = email
}

func (e *employee) GetEmail() string {
	return e.email
}

func (e *employee) SetSector(sector string) {
	e.sector = sector
}

func (e *employee) GetSector() string {
	return e.sector
}

func (e *employee) SetUnit(unit string) {
	e.unit = unit
}

func (e *employee) GetUnit() string {
	return e.unit
}

func (e *employee) SetAdministrator(administrator bool) {
	e.administrator = administrator
}

func (e *employee) GetAdministrator() bool {
	return e.administrator
}

func (e *employee) SetPassword(password string) {
	e.password = password
	e.Encrypt()
}

func (e *employee) GetPassword() string {
	return e.password
}

func (e *employee) SetFeedbacks(feedbacks []Feedback) {
	e.feedbacks = feedbacks
}

func (e *employee) GetFeedbacks() []Feedback {
	return e.feedbacks
}

func (e *employee) Encrypt() {
	h := md5.New()
	defer h.Reset()

	h.Write([]byte(e.password))
	e.password = hex.EncodeToString(h.Sum(nil))
}

func (e *employee) Authenticate() (string, *rest.RestErr) {
	jwt_secret := config.GetConfig().JWT

	claims := jwt.MapClaims{
		"registry":      e.registry,
		"email":         e.email,
		"administrator": e.administrator,
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token_str, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", rest.NewInternalServerErr("error trying to generate jwt token")
	}

	return token_str, nil
}
