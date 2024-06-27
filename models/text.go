package models

import "gorm.io/gorm"

type Text struct {
	gorm.Model
	Content string
}

type History struct {
	gorm.Model
	Content string
	Action  string
}
