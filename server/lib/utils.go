package lib

import (
	"bytes"
	crypto "crypto/rand"
	"encoding/json"
	"io"
	rand "math/rand"
	"net/http"
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

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(crypto.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func SendSMS(phone, token string) {
	type PostBody struct {
		MobileNumber string `json:"mobileNumber"`
		Message      string `json:"message"`
	}

	pb := &PostBody{MobileNumber: phone, Message: "Please use the code " + token + " to authenticate"}
	jsonStr, err := json.Marshal(pb)
	posturl := "https://m183.gibz-informatik.ch/api/sms/message"

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Api-Key", "NwAzADkAOQAxADUANgA5ADMAMAA1ADIAMgA0ADAANQA0ADQA")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
