package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Simulate the path to the config file
	path := "."

	// Create a temporary environment variable for testing
	os.Setenv("APPL_ENV", "testing")

	// Initialize the configuration
	cfg := Init(path)

	// Assert the expected values from the configuration
	assert.Equal(t, path, cfg.BasePath)
	assert.Equal(t, "development", cfg.AppEnv)
	assert.Equal(t, ":3000", cfg.DevelopmentPort)

	// Clean up the temporary environment variable
	os.Unsetenv("APP_ENV")
}
