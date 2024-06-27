package models

import "gorm.io/gorm"

type Text struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
}

type History struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	Action  string `gorm:"type:enum('add', 'delete');not null"`
}
