package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoEntry struct {
	gorm.Model
}

func OpenDatabase() *gorm.DB {
	r, _ := gorm.Open(sqlite.Open("test.db")) // unnecessary function as it stands
	return r
}

func HandleCreate(entries ...TodoEntry) {
	db := OpenDatabase()
	for _, entry := range entries {
		db.Create(&entry) // unnecessary function as it stands
	}
}

func HandleQuery(id string) {
	db := OpenDatabase()
	db.Model(&TodoEntry{}).Where("ID = ?", id) // especially untested
}

func HandleDelete() {

}

func HandleUpdate() {

}
