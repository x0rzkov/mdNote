package model

import (
	"time"
)

type Note struct {
	ID   string `gorm:"PRIMARY_KEY;NOT NULL;UNIQUE;column:id",json:"id,omitempty"`
	UserID   string `gorm:"FOREIGNKEY;NOT NULL;column:user_id",json:"user_id,omitempty"`
	Category string `gorm:"default:'default';NOT NULL;column:category",json:"category,omitempty"`
	Title string `gorm:"NOT NULL;column:title",json:"title,omitempty"`
	Content string `gorm:"NOT NULL;column:content",json:"content,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at",json:"created_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at",json:"deleted_at,omitempty"`
}