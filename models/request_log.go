package models

import "gorm.io/gorm"

type RequestLog struct {
	gorm.Model
	Email      string `json:"email"`
	Endpoint   string `json:"endpoint"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
