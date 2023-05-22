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

func (r *mysql) Save(ctx context.Context, data entity.EmailSent) error {
	tx := r.Sql.WithContext(ctx).Begin()

	if err := tx.Model(entity.EmailSent{}).Create(&data).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "mysql.Save Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *mysql) GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error) {
	var (
		result entity.EmailSent
		tx     = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.EmailSent{}).Where(map[string]interface{}{"uid": uid}).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(err, "mysql.GetOneByUID Exception", nil)
		return nil, err
	}

	return &result, nil
}

func (r *mysql) GetCollection(ctx context.Context) ([]entity.EmailSent, error) {
	var (
		results = make([]entity.EmailSent, 0)
		tx      = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.EmailSent{}).Find(&results).Error; err != nil {
		r.Log.Error(err, "mysql.GetCollection Exception", nil)
		return nil, err
	}

	return results, nil
}
