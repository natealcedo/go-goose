package models

import "time"

type Post struct {
	ID        string    `gorm:"primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostId" json:"comments,omitempty"`
}
