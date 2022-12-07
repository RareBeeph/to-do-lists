package main

import (
	"gorm.io/gorm"

	"backend/server"
)

type tester struct {
	gorm.Model
	contents []byte
}

func main() {
	/*
		db, err := gorm.Open(sqlite.Open("test.db"))
		if err != nil {
			log.Fatalln(err.Error())
		}
		//bytesRead, _ := io.ReadAll(os.Stdin)

		db.AutoMigrate(&tester{})

		//db.Create(&tester{contents: bytesRead})
		result := tester{}
		db.Model(&tester{}).First(&result)
		fmt.Println(result.ID)
	*/

	server.CreateServer()

}
