package model

import "gorm.io/gorm"

type Mood struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   float64
	Color   uint32
	Entries []Entry
}

func AddMood(name string, value float64, color uint32) bool {
	var new Mood
	res := db.Where(Mood{Name: name}).Attrs(Mood{Value: value, Color: color}).FirstOrCreate(&new)
	return res.RowsAffected != 0
}

func SeedMoods() {
	var super, good, ok, bad, ugly Mood
	db.Where(Mood{Name: "Super"}).Attrs(Mood{Value: 5, Color: 0xFFBC42FF}).FirstOrCreate(&super)
	db.Where(Mood{Name: "Good"}).Attrs(Mood{Value: 4, Color: 0xD81159FF}).FirstOrCreate(&good)
	db.Where(Mood{Name: "Ok"}).Attrs(Mood{Value: 3, Color: 0x8F2D56FF}).FirstOrCreate(&ok)
	db.Where(Mood{Name: "Bad"}).Attrs(Mood{Value: 2, Color: 0x218380FF}).FirstOrCreate(&bad)
	db.Where(Mood{Name: "Ugly"}).Attrs(Mood{Value: 1, Color: 0x73D2DEFF}).FirstOrCreate(&ugly)
}
