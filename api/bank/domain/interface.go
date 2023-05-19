package domain

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetList(ctx *gin.Context)
	CreateNew(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateOne(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Service interface {
	GetList(ctx context.Context) ([]entity.Bank, error)
	CreateNew(ctx context.Context, data CreateBankDto) (*entity.Bank, error)
	GetOne(ctx context.Context, oid string) (*entity.Bank, error)
	UpdateOne(ctx context.Context, oid string, data UpdateBankDto) (*entity.Bank, error)
	Delete(ctx context.Context, oid string) error
}

type Repository interface {
	GetCollection(ctx context.Context) ([]entity.Bank, error)
	Create(ctx context.Context, data entity.Bank) (*entity.Bank, error)
	GetOneByCode(ctx context.Context, code string) (*entity.Bank, error)
	GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error)
	UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error)
	Delete(ctx context.Context, bank entity.Bank) error
}
