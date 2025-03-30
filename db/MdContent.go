package db

import (
	"gorm.io/gorm"
)

// Consider https://github.com/swaggest/rest/
// for open api spec
type MdContent struct {
	gorm.Model
	Name            string
	CreatedByUserId string
	Hidden          bool
	Content         string
}

type MdContentCreated struct {
	CreatedByUserId string    `json:"createdByUserId"`
	CreatedEntity   MdContent `json:"createdEntity"`
}
