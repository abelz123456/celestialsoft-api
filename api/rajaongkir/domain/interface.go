package domain

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetHistories(ctx *gin.Context)
	GetProvince(ctx *gin.Context)
	GetCity(ctx *gin.Context)
	CostInfo(ctx *gin.Context)
}

type Service interface {
	GetProvince(ctx context.Context) (interface{}, error)
	GetCity(ctx context.Context, provinceID string) (interface{}, error)
	GetCostInfo(ctx context.Context, deliveryData CostInfoPayload) (*entity.Rajaongkir, error)
	GetCostHistories(ctx context.Context) ([]entity.Rajaongkir, error)
}

type Repository interface {
	GetRajaongkirProvince(ctx context.Context) (interface{}, error)
	GetRajaongkirCity(ctx context.Context, provinceID string) (interface{}, error)
	GetRajaongkirCost(ctx context.Context, deliveryData CostInfoPayload) (map[string]interface{}, error)

	Save(ctx context.Context, data entity.Rajaongkir) error
	GetOneByHashData(ctx context.Context, hash string) (*entity.Rajaongkir, error)
	GetCollection(ctx context.Context) ([]entity.Rajaongkir, error)
}
