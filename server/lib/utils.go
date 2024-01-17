package lib

import (
	"math/rand"
	"time"
)

func GenerateRandomString(lenght int, includeSpecial bool) string {

	var letterRunes []rune

	if includeSpecial {
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-.,;:_$!+*/")
	} else {
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}

	b := make([]rune, lenght)
	rand.NewSource(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CheckHashPassword(hash, password string) bool {
	return true
}

func HashPassword(password string) string {
	return password
}

func ValidatePhoneNumber(phone string) bool {
	return true
}
