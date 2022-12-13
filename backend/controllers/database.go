package controllers

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/models"
)

var db *gorm.DB

func init() {
	r, _ := gorm.Open(sqlite.Open("test.db"))
	db = r
	db.AutoMigrate(&models.TodoEntry{})
}

// TODO: in general, probably generalize what struct the actions are being performed on

func HandleCreate(entries ...models.TodoEntry) {
	// TODO: maybe generalize the creation process to allow for easier addition or removal of variables in the struct
	for _, entry := range entries {
		db.Create(&entry)
	}
}

func HandleQuery(id string) models.TodoEntry {
	var result models.TodoEntry

	db.Model(&models.TodoEntry{}).Where("ID = ?", id).First(&result)
	log.Println(result)
	return result
}

func HandleQueryAll() []models.TodoEntry {
	// TODO: it might be better to just make HandleQuery more versatile

	var test []models.TodoEntry
	db.Find(&test)
	return test
}

func HandleDelete(id string) string {
	entry := db.Model(&models.TodoEntry{}).Where("ID = ?", id).First(&models.TodoEntry{})

	if entry.Error == nil {
		// TODO: make this not fugly

		var entrySingle models.TodoEntry
		entry.First(&entrySingle)
		result := "Successfully deleted entry with ID " + (string)(entrySingle.ID)
		entry.Delete(&models.TodoEntry{})
		return result
	}
	return "a problem occurred" // TODO: specify; probably make this function return (string, error) or something
}

func HandleUpdate(id string, body string) (string, error) {
	var entry models.TodoEntry

	if db.Model(&models.TodoEntry{}).Where("ID = ?", id).First(&entry).Error == nil {
		// TODO: this is a cursed if condition

		entry.Body = body
		db.Save(&entry)
		log.Println(entry)
		return "Successfully modified entry with ID " + id, nil
	}
	return "", errors.New("a problem occurred") // TODO: specify
}
