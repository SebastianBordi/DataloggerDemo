package model

import "time"

type BasicResponse struct {
	StatusCode int
	Message    string
	Timestamp  time.Time
}
