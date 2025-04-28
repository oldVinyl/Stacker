package server

import "time"

type Stack struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreationTime time.Time `json:"creationTime"`
	UpdateTime   time.Time `json:"updateTime"`
}
