package db

import (
	"gorm.io/gorm"
)

type MdContent struct {
	gorm.Model
	Name        string
	DateCreated string
	DateUpdated string
	Hidden      bool
	Content     uint
}
