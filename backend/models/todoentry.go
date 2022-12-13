package models

import "gorm.io/gorm"

type TodoEntry struct {
	gorm.Model
	Body string
	// TODO: include actually useful information
}
