package utils

import (
	"fmt"
	"math/rand"
)

func ShortLinksCreator() (link string) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, 8)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Print(link)
	return
}
