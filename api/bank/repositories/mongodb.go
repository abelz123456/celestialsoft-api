package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	Mongo *mongo.Database
	Log   log.Log
}

func (r *mongodb) GetCollection(ctx context.Context) ([]entity.Bank, error) {
	return []entity.Bank{}, nil
}

func (r *mongodb) Create(ctx context.Context, data entity.Bank) (*entity.Bank, error) {
	return nil, nil
}

func (r *mongodb) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	return nil, nil
}

func (r *mongodb) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	return nil, nil
}

func (r *mongodb) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	return nil, nil
}

func (r *mongodb) Delete(ctx context.Context, bank entity.Bank) error {
	return nil
}
