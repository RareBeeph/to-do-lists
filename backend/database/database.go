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

func OpenDatabase() *gorm.DB {
	// TODO: not a fan of having to open the database for each request handled. look into avoiding that

	r, _ := gorm.Open(sqlite.Open("test.db"))
	r.AutoMigrate(&TodoEntry{}) // TODO: figure out if this is a bad idea
	return r
}

// TODO: in general, probably generalize what struct the actions are being performed on

func HandleCreate(entries ...TodoEntry) {
	// TODO: maybe generalize the creation process to allow for easier addition or removal of variables in the struct

	db := OpenDatabase()
	for _, entry := range entries {
		db.Create(&entry)
	}
}

func HandleQuery(id string) TodoEntry {
	db := OpenDatabase()
	var result TodoEntry

	db.Model(&TodoEntry{}).Where("ID = ?", id).First(&result)
	log.Println(result)
	return result
}

func HandleQueryAll() []TodoEntry {
	// TODO: it might be better to just make HandleQuery more versatile

	db := OpenDatabase()
	var test []TodoEntry
	db.Find(&test)
	return test
}

func HandleDelete(id string) string {
	db := OpenDatabase()
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
	db := OpenDatabase()
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
