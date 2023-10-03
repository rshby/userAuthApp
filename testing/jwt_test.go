package testing

import (
	"log"
	"testing"
	"userAuthApp/helper"
)

func TestJwtGenerateToken(t *testing.T) {
	email := "reoshby@gmailcom"

	token, err := helper.GenerateToken(email)
	if err != nil {
		log.Println(err)
	}

	log.Println(token)
}
