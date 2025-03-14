package models

import (
	"gorm.io/gorm"
)

type RateLimit struct {
	gorm.Model
	Requests   int `json:"requests"`
	TimeWindow int `json:"time_window"`
}
