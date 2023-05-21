package repositories

import (
	"context"
	"mime/multipart"
	"os"
	"testing"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
)

func Test_postgresql_SaveLocalStorage(t *testing.T) {
	// Create a postgresql instance with a test database connection
	ctx := context.Background()
	_, db := mockdata.NewPosgreSQLMock(t)
	repo := &postgresql{
		Sql: db,
		Log: log.NewLog(),
	}

	t.Run("Success", func(t *testing.T) {
		req := createRequestMultipartFiles(t)

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		assert.NoError(t, os.Mkdir(".test_public", 0755))

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(ctx, *fileHeader, dest)

		assert.NoError(t, err)

		_, err = os.Stat(dest)
		if err != nil {
			assert.Equal(t, os.IsNotExist(err), false)
		}

		assert.NoError(t, os.RemoveAll(".test_public"))
	})

	t.Run("Error destination dir not exists", func(t *testing.T) {
		req := createRequestMultipartFiles(t)

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		assert.NoError(t, os.Mkdir(".test_public_", 0777))

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(ctx, *fileHeader, dest)

		assert.Error(t, err)

		assert.Contains(t, err.Error(), "no such file or directory")

		assert.NoError(t, os.RemoveAll(".test_public_"))
	})

	t.Run("Error to open multipart.FileHeader", func(t *testing.T) {
		req := createRequestMultipartFiles(t)
		ll := multipart.FileHeader{}

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(context.Background(), ll, dest)

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "open : no such file or directory")
	})
}

func Test_postgresql_DeleteLocalStorage(t *testing.T) {
	// Create a postgresql instance with a test database connection
	ctx := context.Background()
	_, db := mockdata.NewPosgreSQLMock(t)
	repo := &postgresql{
		Sql: db,
		Log: log.NewLog(),
	}

	t.Run("Success", func(t *testing.T) {
		req := createRequestMultipartFiles(t)

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		assert.NoError(t, os.Mkdir(".test_public", 0755))

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(ctx, *fileHeader, dest)

		assert.NoError(t, err)

		_, err = os.Stat(dest)
		if err != nil {
			assert.Equal(t, os.IsNotExist(err), false)
		}

		err = repo.DeleteLocalStorage(ctx, dest)

		assert.NoError(t, err)

		assert.NoError(t, os.RemoveAll(".test_public"))
	})
}
