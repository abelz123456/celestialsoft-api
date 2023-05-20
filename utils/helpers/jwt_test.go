package helpers

import (
	"testing"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	// Create a mock config
	mockConfig := config.Config{
		SecretKey:      "secret",
		JwtExpiredTime: 60, // 1 minute
	}

	// Create the JwtHelper
	jwtHelper := NewJwtHelper(mockConfig)

	// Generate a token
	oid := "user_id"
	token := jwtHelper.CreateToken(oid)

	// Parse the token to get the claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(mockConfig.SecretKey), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Extract the claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	// Check the claims
	assert.Equal(t, oid, claims["sub"])
	// exp := time.Now().Add(time.Minute * time.Duration(mockConfig.JwtExpiredTime)).Unix()
	// assert.True(t, claims.VerifyExpiresAt(exp, true))
}

func TestParseToken_ValidToken(t *testing.T) {
	// Create a mock config
	mockConfig := config.Config{
		SecretKey:      "secret",
		JwtExpiredTime: 60, // 1 minute
	}

	// Create the JwtHelper
	jwtHelper := NewJwtHelper(mockConfig)

	// Generate a token
	oid := "user_id"
	token := jwtHelper.CreateToken(oid)

	// Parse the token
	parsedOid := jwtHelper.ParseToken(token)

	// Assert that the parsed OID matches the original OID
	assert.Equal(t, oid, parsedOid)
}

func TestParseToken_InvalidToken(t *testing.T) {
	// Create a mock config
	mockConfig := config.Config{
		SecretKey:      "secret",
		JwtExpiredTime: 60, // 1 minute
	}

	// Create the JwtHelper
	jwtHelper := NewJwtHelper(mockConfig)

	// Parse the token
	parsedOid := jwtHelper.ParseToken("invalidJWTToken")

	// Assert that the parsed OID is empty (indicating an invalid token)
	assert.Empty(t, parsedOid)
}
