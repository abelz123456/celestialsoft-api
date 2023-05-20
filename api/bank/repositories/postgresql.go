package repositories

import (
	"context"
	"errors"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type postgresql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.Bank, error) {
	var results = make([]entity.Bank, 0)
	stmt := r.Sql.WithContext(ctx).
		Find(&results)

	if stmt.Error != nil {
		r.Log.Error(stmt.Error, "postgresql.GetCollection Exception", nil)
		return nil, stmt.Error
	}

	return results, nil
}

func (r *postgresql) Create(ctx context.Context, data entity.Bank) (*entity.Bank, error) {
	stmt := r.Sql.WithContext(ctx).
		Create(&data)

	if stmt.Error != nil {
		stmt.Rollback()
		r.Log.Error(stmt.Error, "postgresql.Create Exception", nil)
		return nil, stmt.Error
	}

	return &data, nil
}

func (r *postgresql) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where(&entity.Bank{BankCode: code}).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "postgresql.GetOneByCode Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}
func (r *postgresql) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", oid).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "postgresql.GetOneByOid Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}

func (r *postgresql) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", bank.Oid).
		Updates(&newData)

	if stmt.Error != nil {
		stmt.Rollback()
		return nil, stmt.Error
	}

	return r.GetOneByOid(ctx, bank.Oid)
}

func (r *postgresql) Delete(ctx context.Context, bank entity.Bank) error {
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", bank.Oid).
		Delete(&entity.Bank{})

	if stmt.Error != nil {
		r.Log.Error(stmt.Error, "postgresql.Delete Exception", nil)
		return stmt.Error
	}

	return nil
}
