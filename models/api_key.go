package models

import "gorm.io/gorm"

type APIKey struct {
	gorm.Model
	UserId uint `json:"user_id" gorm:"index"`
	ApiKey string `json:"api_key" gorm:"uniqueIndex"`
	BackendUrl string `json:"backend_url"`
	Description string `json:"description"`
}


