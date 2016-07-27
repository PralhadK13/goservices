package model

import (
	"time"
)

type Message struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userid"`
	Message     string    `json:"message"`
	LastUpdated time.Time `json:"lastupdated"`
}
