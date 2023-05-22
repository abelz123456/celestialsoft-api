package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
)

func TestLoadRoute(t *testing.T) {
	// Create a new manager instance.
	mgr := mockdata.NewFakeManager(t, database.MySQL)
	mgr.Config.AppEnv = "testing"
	mgr.Config.StaticFilePath = "public"

	t.Run("Ping server", func(t *testing.T) {
		// Create an HTTP test server with the mock handler
		server := httptest.NewServer(mgr.Server.Engine)
		defer server.Close()

		LoadRoute(mgr)

		// Send a request to the mock server
		resp, err := http.Get(server.URL + "/ping")
		assert.NoError(t, err)
		defer resp.Body.Close()

		// Assert that the response is successful.
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		// // Assert that the response body is "pong".
		ioRead, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, string(ioRead), "pong")
		server.Close()
	})
}
