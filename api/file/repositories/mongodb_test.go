package repositories

import (
	"context"
	"mime/multipart"
	"os"
	"testing"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_mongodb_SaveLocalStorage(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mtest := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Mongo: mtest.DB,
		Log:   log.NewLog(),
	}

	t.Run("Success", func(t *testing.T) {
		req := createRequestMultipartFiles(t)

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		assert.NoError(t, os.Mkdir(".test_public", 0755))

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(context.Background(), *fileHeader, dest)

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
		err = repo.SaveLocalStorage(context.Background(), *fileHeader, dest)

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

func Test_mongodb_DeleteLocalStorage(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mtest := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Mongo: mtest.DB,
		Log:   log.NewLog(),
	}

	t.Run("Success", func(t *testing.T) {
		req := createRequestMultipartFiles(t)

		assert.NotNil(t, req)

		_, fileHeader, err := req.FormFile("content")

		assert.NoError(t, err)

		assert.NoError(t, os.Mkdir(".test_public", 0755))

		// Call the SaveLocalStorage function
		dest := ".test_public/" + fileHeader.Filename
		err = repo.SaveLocalStorage(context.Background(), *fileHeader, dest)

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

func Test_mongodb_Save(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Log: log.NewLog(),
	}

	mt.Run("Success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := repo.Save(ctx, testData)
		assert.NoError(t, err)
	})

	mt.Run("error", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := repo.Save(ctx, testData)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func Test_mongodb_GetCollection(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Log: log.NewLog(),
	}

	mt.Run("Success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: testData.UID}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetCollection(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.GreaterOrEqual(t, len(result), 1)
		assert.Equal(t, result[0].UID, testData.UID)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetCollection(ctx)

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
		assert.Equal(t, len(result), 0)
	})

	mt.Run("error Decode", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: 123456}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetCollection(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding key uid")
		assert.Equal(t, len(result), 0)
		assert.Nil(t, result)
	})
}

func Test_mongodb_GetOneByUID(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Log: log.NewLog(),
	}

	mt.Run("Success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: testData.UID}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByUID(ctx, testData.UID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.LocalPath, testData.LocalPath)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByUID(ctx, testData.UID)

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
	})

	mt.Run("error Decode", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: 123456}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByUID(ctx, testData.UID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding key uid")
		assert.Nil(t, result)
	})

	mt.Run("ErrNoDocuments", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(0, "test.bank", mtest.FirstBatch),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByUID(ctx, testData.UID)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	mt.Run("No Error but uid is blank", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByUID(ctx, testData.UID)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func Test_mongodb_UpdateOne(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Log: log.NewLog(),
	}

	mt.Run("Success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: testData.UID}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: testData.UID}, {Key: "localPath", Value: mockdata.LocalFileMock[1].LocalPath}, {Key: "originalName", Value: mockdata.LocalFileMock[1].OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.UpdateOne(ctx, testData, mockdata.LocalFileMock[1])

		assert.NoError(t, err)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.UpdateOne(ctx, testData, mockdata.LocalFileMock[1])
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Test Error")
	})
}

func Test_mongodb_Delete(t *testing.T) {
	// Create a MongoDB instance with a test database connection
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)
	repo := &mongodb{
		Log: log.NewLog(),
	}

	mt.Run("Success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.localFile", mtest.FirstBatch, bson.D{{Key: "uid", Value: testData.UID}, {Key: "localPath", Value: testData.LocalPath}, {Key: "originalName", Value: testData.OriginalName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.Delete(ctx, testData)
		assert.NoError(t, err)

	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.LocalFileMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.Delete(ctx, testData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Test Error")
	})
}
