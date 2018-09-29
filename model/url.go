package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

// Link Table
type Link struct {
	gorm.Model
	URL          string `json:"url"`
	ShortenedURL string `json:"shortened_url"`
}

// LinkURL fetch urls
type LinkURL struct {
	URL          string `json:"url"`
	ShortenedURL string `json:"shortened_url"`
}

// CodeGenerator generates 8 character code or
// retrieve code from database
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
		link = Link{URL: url, ShortenedURL: shorten}
		db.NewRecord(link)
		db.Create(&link)
	} else {
		shorten = link.ShortenedURL
	}
	return shorten
}

// Geturls fetch all url from database
func Geturls() []LinkURL {
	db, err := gorm.Open("sqlite3", "db/test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	var urls []LinkURL
	db.Raw("SELECT url,shortened_url FROM links").Scan(&urls)
	return urls
}
