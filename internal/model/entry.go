package model

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Notes      string
	MoodId     uint
	Mood       Mood
	Activities []*Activity `gorm:"many2many:entry_activities;"`
}
