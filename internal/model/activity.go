package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Icon    string
	Entries []*Entry `gorm:"many2many:entry_activities;"`
}
