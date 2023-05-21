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
	tx := r.Sql.WithContext(ctx).Begin()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "mysql.Create Exception", nil)
		return nil, err
	}

	tx.Commit()
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

	return &result, nil
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

	return &result, nil

}

func (r *mysql) UpdateOne(ctx context.Context, bank entity.Bank, newData entity.Bank) (*entity.Bank, error) {
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

func (r *mysql) Delete(ctx context.Context, bank entity.Bank) error {
	tx := r.Sql.WithContext(ctx).Begin()

	stmt := tx.Model(&entity.Bank{}).
		Delete(bank)

	if stmt.Error != nil {
		tx.Rollback()
		r.Log.Error(stmt.Error, "mysql.Delete Exception", nil)
		return stmt.Error
	}

	tx.Commit()
	return nil
}
