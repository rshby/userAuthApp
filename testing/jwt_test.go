package testing

import (
	"log"
	"testing"
	"userAuthApp/helper"
)

func TestJwtGenerateToken(t *testing.T) {
	id := 4

	token, err := helper.GenerateToken(id)
	if err != nil {
		log.Println(err)
	}

	log.Println(token)
}
