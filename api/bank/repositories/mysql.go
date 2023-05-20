package repositories

import (
	"context"
	"errors"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type mysql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *mysql) GetCollection(ctx context.Context) ([]entity.Bank, error) {
	var results = make([]entity.Bank, 0)
	stmt := r.Sql.WithContext(ctx).
		Find(&results)

	if stmt.Error != nil {
		r.Log.Error(stmt.Error, "mysql.GetCollection Exception", nil)
		return nil, stmt.Error
	}

	return results, nil
}

func (r *mysql) Create(ctx context.Context, data entity.Bank) (*entity.Bank, error) {
	stmt := r.Sql.WithContext(ctx).
		Create(&data)

	if stmt.Error != nil {
		stmt.Rollback()
		r.Log.Error(stmt.Error, "mysql.Create Exception", nil)
		return nil, stmt.Error
	}

	return &data, nil
}

func (r *mysql) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where("bankCode = ?", code).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "mysql.GetOneByCode Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}

func (r *mysql) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", oid).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "mysql.GetOneByOid Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}

func (r *mysql) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", bank.Oid).
		Updates(&newData)

	if stmt.Error != nil {
		stmt.Rollback()
		return nil, stmt.Error
	}

	return r.GetOneByOid(ctx, bank.Oid)
}

func (r *mysql) Delete(ctx context.Context, bank entity.Bank) error {
	stmt := r.Sql.WithContext(ctx).
		Where("oid = ?", bank.Oid).
		Delete(&entity.Bank{})

	if stmt.Error != nil {
		r.Log.Error(stmt.Error, "mysql.Delete Exception", nil)
		return stmt.Error
	}

	return nil
}
