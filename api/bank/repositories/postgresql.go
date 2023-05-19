package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type postgresql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.Bank, error) {
	return []entity.Bank{}, nil
}

func (r *postgresql) Create(ctx context.Context, data entity.Bank) (*entity.Bank, error) {
	return nil, nil
}

func (r *postgresql) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	return nil, nil
}

func (r *postgresql) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	return nil, nil
}

func (r *postgresql) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	return nil, nil
}

func (r *postgresql) Delete(ctx context.Context, bank entity.Bank) error {
	return nil
}
