package Entities

import "time"

type ApiToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func NewApiToken() *ApiToken {
	return &ApiToken{}
}
