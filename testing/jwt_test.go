package testing

import (
	"github.com/stretchr/testify/assert"
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

	assert.Nil(t, err)
	assert.NotNil(t, token)
}
