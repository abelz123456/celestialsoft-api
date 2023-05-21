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
	tx := r.Sql.WithContext(ctx).Begin()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "postgresql.Create Exception", nil)
		return nil, err
	}

	tx.Commit()
	return &data, nil
}

func (r *postgresql) GetOneByCode(ctx context.Context, code string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where(map[string]string{"bankCode": code}).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "postgresql.GetOneByCode Exception", nil)
		return nil, stmt.Error
	}

	return &result, nil
}

func (r *postgresql) GetOneByOid(ctx context.Context, oid string) (*entity.Bank, error) {
	var result entity.Bank
	stmt := r.Sql.WithContext(ctx).
		Where(map[string]string{"oid": oid}).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "postgresql.GetOneByOid Exception", nil)
		return nil, stmt.Error
	}

	return &result, nil

}

func (r *postgresql) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
	tx := r.Sql.WithContext(ctx).Begin()

	stmt := tx.Model(&entity.Bank{}).
		Where("oid = ?", bank.Oid).
		Updates(map[string]interface{}{"bankName": newData.BankName})

	if stmt.Error != nil {
		tx.Rollback()
		return nil, stmt.Error
	}

	tx.Commit()
	return r.GetOneByOid(ctx, bank.Oid)
}

func (r *postgresql) Delete(ctx context.Context, bank entity.Bank) error {
	tx := r.Sql.WithContext(ctx).Begin()

	stmt := tx.Model(&entity.Bank{}).
		Delete(bank)

	if stmt.Error != nil {
		tx.Rollback()
		r.Log.Error(stmt.Error, "postgresql.Delete Exception", nil)
		return stmt.Error
	}

	tx.Commit()
	return nil
}
