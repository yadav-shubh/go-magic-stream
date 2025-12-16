package models

import "time"

type ContactMessage struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Query     string    `json:"query" bson:"query"`
	IPAddress string    `json:"-" bson:"ip_address"`
	UserAgent string    `json:"-" bson:"user_agent"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
}

type ContactMessageDTO struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Query string `json:"query" validate:"required,min=2"`
}
