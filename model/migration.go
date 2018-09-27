package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitalMigration() {
	fmt.Println("Migration started")

	db, err := gorm.Open("sqlite3", "db/test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Link{})
	fmt.Println("Migration ended")
}
