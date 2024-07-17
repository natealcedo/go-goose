package models

import (
	"encoding/json"
	"time"
)

type Comment struct {
	ID            string    `gorm:"primaryKey;default:gen_random_uuid()" json:"id"`
	PostId        string    `gorm:"column:post_id" json:"-"`
	Content       string    `gorm:"column:content" json:"content,omitempty"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	IncludePostID bool      `gorm:"-" json:"-"` // Temporary field to control JSON serialization
}

func (c Comment) MarshalJSON() ([]byte, error) {
	type Alias Comment
	if c.IncludePostID {
		return json.Marshal(&struct {
			PostId string `json:"post_id,omitempty"`
			*Alias
		}{
			PostId: c.PostId,
			Alias:  (*Alias)(&c),
		})
	}
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&c),
	})
}
