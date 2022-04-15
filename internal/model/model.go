package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Title    string
	SubTitle string
	Text     string
}
