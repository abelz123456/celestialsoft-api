package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	mongo *mongo.Database
	log   log.Log
	repo  *repository
}

func (r *mongodb) GetRajaongkirProvince(ctx context.Context) (interface{}, error) {
	return r.repo.getRajaongkirProvince(ctx)
}

func (r *mongodb) GetRajaongkirCity(ctx context.Context, provinceID string) (interface{}, error) {
	return r.repo.getRajaongkirCity(ctx, provinceID)
}

func (r *mongodb) GetRajaongkirCost(ctx context.Context, deliveryData domain.CostInfoPayload) (map[string]interface{}, error) {
	return r.repo.getRajaongkirCost(ctx, deliveryData)
}

func (r *mongodb) Save(ctx context.Context, data entity.Rajaongkir) error {
	return nil
}

func (r *mongodb) GetOneByHashData(ctx context.Context, hash string) (*entity.Rajaongkir, error) {
	return nil, nil
}

func (r *mongodb) GetCollection(ctx context.Context) ([]entity.Rajaongkir, error) {
	return nil, nil
}
