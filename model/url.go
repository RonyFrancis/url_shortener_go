package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type Link struct {
	gorm.Model
	Url          string `json:"url"`
	ShortenedUrl string `json:"shortened_url"`
}

func CodeGenerator(size int, url string) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4",
		"5", "6", "7", "8", "9"}
	var shorten string

	// database connection
	// TODO: move db connection to separate function
	// and call it globally
	db, err := gorm.Open("sqlite3", "db/test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	var link Link
	db.Find(&link, "url= ?", url)

	// check record is present
	if link.ID == 0 {
		// save record if not present
		shorten = ""
		rand.Seed(time.Now().UnixNano())
		for i := 0; i <= size; i++ {
			shorten += letters[rand.Intn(len(letters))]
		}
		link = Link{Url: url, ShortenedUrl: shorten}
		db.NewRecord(link)
		db.Create(&link)
	} else {
		shorten = link.ShortenedUrl
	}
	return shorten
}
