package url

import (
	"math/rand"
	"time"
)

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
