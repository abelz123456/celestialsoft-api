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

func TestMongoDB_GetCollection(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetCollection(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.GreaterOrEqual(t, len(result), 1)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		// testData := mockdata.BankMock[0]

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
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: 123456}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetCollection(ctx)
		assert.Error(t, err)
		assert.Equal(t, len(result), 0)
		assert.Nil(t, result)
	})
}

func TestMongoDB_Create(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		result, err := repo.Create(ctx, mockdata.BankMock[0])
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

		result, err := repo.Create(ctx, mockdata.BankMock[0])
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestMongoDB_GetOneByCode(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByCode(ctx, testData.BankCode)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.BankName, testData.BankName)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByCode(ctx, testData.BankCode)

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
	})

	mt.Run("error Decode", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: 123456}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByCode(ctx, testData.BankCode)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	mt.Run("ErrNoDocuments", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(0, "test.bank", mtest.FirstBatch),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByCode(ctx, mockdata.BankMock[1].BankCode)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	mt.Run("No Error but oid is blank", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByCode(ctx, testData.BankCode)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestMongoDB_GetOneByOid(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByOid(ctx, testData.Oid)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.BankName, testData.BankName)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByOid(ctx, testData.Oid)

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
	})

	mt.Run("error Decode", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: 123456}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByOid(ctx, testData.Oid)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	mt.Run("ErrNoDocuments", func(mt *mtest.T) {
		repo.Mongo = mt.DB

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(0, "test.bank", mtest.FirstBatch),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.GetOneByOid(ctx, mockdata.BankMock[1].Oid)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestMongoDB_UpdateOne(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: mockdata.BankMock[1].BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.UpdateOne(ctx, testData, mockdata.BankMock[1])

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.BankName, mockdata.BankMock[1].BankName)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		result, err := repo.UpdateOne(ctx, testData, mockdata.BankMock[1])

		assert.Contains(t, err.Error(), "Test Error")
		assert.Nil(t, result)
	})
}

func TestMongoDB_Delete(t *testing.T) {
	ctx := context.Background()
	mt := mockdata.NewMongoDBMock(t, ctx)

	repo := &mongodb{
		Mongo: mt.DB,
		Log:   log.NewLog(),
	}

	mt.Run("success", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCursorResponse(1, "test.bank", mtest.FirstBatch, bson.D{{Key: "oid", Value: testData.Oid}, {Key: "bankCode", Value: testData.BankCode}, {Key: "bankName", Value: testData.BankName}}),
			mtest.CreateSuccessResponse(),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.Delete(ctx, testData)
		assert.NoError(t, err)
	})

	mt.Run("error mongoResult", func(mt *mtest.T) {
		repo.Mongo = mt.DB
		testData := mockdata.BankMock[0]

		mockResponses := []primitive.D{
			mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 1100, Message: "Test Error"}),
		}
		mt.AddMockResponses(mockResponses...)

		err := repo.Delete(ctx, testData)

		assert.Contains(t, err.Error(), "Test Error")
	})
}
