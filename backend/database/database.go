package database

// TODO: pick a better name for the package

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoEntry struct {
	gorm.Model
	Body string
	// TODO: include actually useful information
}

var db *gorm.DB

func init() {
	// TODO: test the recent changes

	r, _ := gorm.Open(sqlite.Open("test.db"))
	db = r
	db.AutoMigrate(&TodoEntry{})
}

// TODO: in general, probably generalize what struct the actions are being performed on

func HandleCreate(entries ...TodoEntry) {
	// TODO: maybe generalize the creation process to allow for easier addition or removal of variables in the struct
	for _, entry := range entries {
		db.Create(&entry)
	}
}

func HandleQuery(id string) TodoEntry {
	var result TodoEntry

	db.Model(&TodoEntry{}).Where("ID = ?", id).First(&result)
	log.Println(result)
	return result
}

func HandleQueryAll() []TodoEntry {
	// TODO: it might be better to just make HandleQuery more versatile

	var test []TodoEntry
	db.Find(&test)
	return test
}

func HandleDelete(id string) string {
	entry := db.Model(&TodoEntry{}).Where("ID = ?", id).First(&TodoEntry{})

	if entry.Error == nil {
		// TODO: make this not fugly

		var entrySingle TodoEntry
		entry.First(&entrySingle)
		result := "Successfully deleted entry with ID " + (string)(entrySingle.ID)
		entry.Delete(&TodoEntry{})
		return result
	}
	return "a problem occurred" // TODO: specify; probably make this function return (string, error) or something
}

func HandleUpdate(id string, body string) (string, error) {
	var entry TodoEntry

	if db.Model(&TodoEntry{}).Where("ID = ?", id).First(&entry).Error == nil {
		// TODO: this is a cursed if condition

		entry.Body = body
		db.Save(&entry)
		log.Println(entry)
		return "Successfully modified entry with ID " + id, nil
	}
	return "", errors.New("a problem occurred") // TODO: specify
}
