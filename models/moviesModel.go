package models

import (
	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	Title    string
	Year     int
	Genre    string
	Director string
}
