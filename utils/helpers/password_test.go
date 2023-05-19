package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	// Create a password.
	password := "password"

	// Hash the password.
	hash, err := HashPassword(password)
	assert.Nil(t, err)

	// Assert that the hash is not empty.
	assert.NotEmpty(t, hash)

	// Assert that the hash is a valid bcrypt hash.
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.Nil(t, err)
}

func TestCheckPasswordHash(t *testing.T) {
	// Create a password and hash.
	password := "password"
	hash, err := HashPassword(password)
	assert.Nil(t, err)

	// Assert that the password matches the hash.
	assert.True(t, CheckPasswordHash(password, hash))

	// Assert that an incorrect password does not match the hash.
	assert.False(t, CheckPasswordHash("incorrect password", hash))
}
