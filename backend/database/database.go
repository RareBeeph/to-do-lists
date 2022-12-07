package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoEntry struct {
	gorm.Model
}

func OpenDatabase() *gorm.DB {
	r, _ := gorm.Open(sqlite.Open("test.db")) // unnecessary function as it stands
	r.AutoMigrate(&TodoEntry{})
	return r
}

func HandleCreate(entries ...TodoEntry) {
	db := OpenDatabase()
	for _, entry := range entries {
		db.Create(&entry) // unnecessary function as it stands
	}
}

func HandleQuery(id string) TodoEntry {
	db := OpenDatabase()
	var result TodoEntry
	db.Model(&TodoEntry{}).Where("ID = ?", id).First(&result) // especially untested
	fmt.Println(result)
	return result
}

func HandleDelete() {

}

func HandleUpdate() {

}
