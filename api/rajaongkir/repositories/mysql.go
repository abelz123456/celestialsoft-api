package repositories

import (
	"context"
	"errors"
	"net/http"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type mysql struct {
	sql  *gorm.DB
	log  log.Log
	repo *repository
}

func (r *mysql) GetRajaongkirProvince(ctx context.Context) (interface{}, error) {
	return r.repo.getRajaongkirProvince(ctx)
}

func (r *mysql) GetRajaongkirCity(ctx context.Context, provinceID string) (interface{}, error) {
	return r.repo.getRajaongkirCity(ctx, provinceID)
}

func (r *mysql) GetRajaongkirCost(ctx context.Context, deliveryData domain.CostInfoPayload) (map[string]interface{}, error) {
	return r.repo.getRajaongkirCost(ctx, deliveryData)
}

func (r *mysql) Save(ctx context.Context, data entity.Rajaongkir) error {
	tx := r.sql.WithContext(ctx).Begin()

	if err := tx.Model(entity.Rajaongkir{}).Create(&data).Error; err != nil {
		tx.Rollback()
		r.log.Error(err, "mysql.Save Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *mysql) GetOneByHashData(ctx context.Context, hash string) (*entity.Rajaongkir, error) {
	var (
		result entity.Rajaongkir
		tx     = r.sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.Rajaongkir{}).Where(map[string]interface{}{"hashData": hash}).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.log.Error(err, "mysql.GetOneByHashData Exception", nil)
		return nil, err
	}

	return &result, nil
}

func (r *mysql) GetCollection(ctx context.Context) ([]entity.Rajaongkir, error) {
	var (
		results = make([]entity.Rajaongkir, 0)
		tx      = r.sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.Rajaongkir{}).Where(map[string]interface{}{"apiStatus": http.StatusOK}).Find(&results).Error; err != nil {
		r.log.Error(err, "mysql.GetCollection Exception", nil)
		return nil, err
	}

	return results, nil
}
