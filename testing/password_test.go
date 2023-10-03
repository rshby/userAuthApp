package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"userAuthApp/helper"
)

func TestCheckPassword(t *testing.T) {
	hashedPassword := "$2a$10$aItOOiMWWNbgNlu0gHP/0OYhAomlXbi/5X.yv8khqcCGb6HaIIRtC"
	password := "123456789012"

	result, err := helper.CheckPasword(hashedPassword, password)

	assert.Nil(t, err)
	assert.True(t, result)
}
