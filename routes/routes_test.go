package routes

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
)

func TestLoadRoute(t *testing.T) {
	// Create a new manager instance.
	mgr := mockdata.NewFakeManager(t, database.MySQL)

	// Create a new HTTP server.
	srv := mgr.Server.HttpServer

	// Start the server.
	go func() {
		LoadRoute(mgr)
		srv.ListenAndServe()
	}()

	// Create a new request.
	req, err := http.NewRequest("GET", "http://127.0.0.1:3030/ping", nil)
	assert.Nil(t, err)

	// Make the request.
	response, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	// Assert that the response is successful.
	assert.Equal(t, response.StatusCode, http.StatusOK)

	// Assert that the response body is "pong".
	ioRead, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(ioRead), "pong")

	// Shutdown the server gracefully.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	assert.NoError(t, srv.Shutdown(ctx))

}
