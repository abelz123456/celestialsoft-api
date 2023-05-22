package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type postgresql struct {
	sql  *gorm.DB
	log  log.Log
	repo *repository
}

func (r *postgresql) GetRajaongkirProvince(ctx context.Context) (interface{}, error) {
	return r.repo.getRajaongkirProvince(ctx)
}

func (r *postgresql) GetRajaongkirCity(ctx context.Context, provinceID string) (interface{}, error) {
	return r.repo.getRajaongkirCity(ctx, provinceID)
}

func (r *postgresql) GetRajaongkirCost(ctx context.Context, deliveryData domain.CostInfoPayload) (map[string]interface{}, error) {
	return r.repo.getRajaongkirCost(ctx, deliveryData)
}

func (r *postgresql) Save(ctx context.Context, data entity.Rajaongkir) error {
	return nil
}

func (r *postgresql) GetOneByHashData(ctx context.Context, hash string) (*entity.Rajaongkir, error) {
	return nil, nil
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.Rajaongkir, error) {
	return nil, nil
}
