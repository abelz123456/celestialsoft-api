package repositories

import (
	"context"
	"testing"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
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

	})
}
