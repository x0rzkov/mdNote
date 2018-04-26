package model

type User struct {
	ID string `gorm:"PRIMARY_KEY;NOT NULL;UNIQUE;column:id" json:"id,omitempty"`
}
