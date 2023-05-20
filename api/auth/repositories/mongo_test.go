package repositories

import (
	"context"
	"testing"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongoDB_Save(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		result, err := repo.Save(ctx, mockdata.PermissionPolicyUserMock[0])
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	mt.Run("error", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		result, err := repo.Save(ctx, mockdata.PermissionPolicyUserMock[0])
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestMongoDB_GetOneByEmail(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.PermissionPolicyUserMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.permissionPolicyUser", mtest.FirstBatch, bson.D{{"emailName", testData.EmailName}, {"oid", testData.Oid}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByEmail(ctx, testData.EmailName)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.EmailName, testData.EmailName)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.PermissionPolicyUserMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByEmail(ctx, testData.EmailName)

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
	})

	mt.Run("error Decode", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.PermissionPolicyUserMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.permissionPolicyUser", mtest.FirstBatch, bson.D{{"emailName", 123456}, {"oid", testData.Oid}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByEmail(ctx, testData.EmailName)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	mt.Run("ErrNoDocuments", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(0, "test.permissionPolicyUser", mtest.FirstBatch),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByEmail(ctx, mockdata.PermissionPolicyUserMock[1].EmailName)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	mt.Run("No Error but oid is blank", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.PermissionPolicyUserMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.permissionPolicyUser", mtest.FirstBatch, bson.D{{"emailName", testData.EmailName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByEmail(ctx, testData.EmailName)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}
