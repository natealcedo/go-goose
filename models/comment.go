package models

import "time"

type Comment struct {
	ID        string    `gorm:"primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	PostId    string    `gorm:"column:post_id" json:"post_id"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Comment) TableName() string {
	return "comments"
}
