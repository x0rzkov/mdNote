package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func OpenAndCreate(dialect, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		return db, fmt.Errorf("gorm.Open: %v", err)
	}

	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			return nil, fmt.Errorf("create user table: %v", err)
		}
	}

	if !db.HasTable(&Note{}) {
		if err := db.CreateTable(&Note{}).Error; err != nil {
			return nil, fmt.Errorf("create note table: %v", err)
		}
		if err := db.Model(&Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
			return nil, fmt.Errorf("add user_id foreign key in course table: %v", err)
		}
	}

	return db, nil
}
