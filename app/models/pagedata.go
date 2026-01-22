package models

import (
	"time"

	"gorm.io/gorm"
)

type PageLinks struct {
	Internal int `json:"internal"`
	External int `json:"external"`
}

type PageData struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	URL         string         `json:"url" gorm:"size:1000;not null;uniqueIndex"`
	Title       string         `json:"title" gorm:"size:255"`
	Description string         `json:"description" gorm:"size:5000"`
	WordCount   int            `json:"word_count" gorm:"default:0"`
	Links       PageLinks      `json:"links" gorm:"embedded"`
	StatusCode  int            `json:"status_code" gorm:"index:idx_status_code"`
	LoadTime    int64          `json:"load_time" gorm:"type:double precision;default:0"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index:idx_created_at,sort:desc,autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index:idx_updated_at,sort:desc,autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
