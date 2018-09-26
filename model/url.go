package model

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type Link struct {
	gorm.Model
	Url          string `json:"url"`
	ShortenedUrl string `json:"shortened_url"`
}

func CodeGenerator(size int) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4",
		"5", "6", "7", "8", "9"}
	var shorten string
	shorten = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= size; i++ {
		shorten += letters[rand.Intn(len(letters))]
	}
	return shorten
}
