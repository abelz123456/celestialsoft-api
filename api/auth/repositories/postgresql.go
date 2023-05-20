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

func (r *postgresql) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	stmt := r.Sql.WithContext(ctx).
		Create(&data)

	if stmt.Error != nil {
		stmt.Rollback()
		r.Log.Error(stmt.Error, "postgresql.Save Exception", nil)
		return nil, stmt.Error
	}

	return &data, nil
}

func (r *postgresql) GetOneByEmail(ctx context.Context, email string) (*entity.PermissionPolicyUser, error) {
	var result entity.PermissionPolicyUser
	stmt := r.Sql.WithContext(ctx).
		Where(&entity.PermissionPolicyUser{EmailName: email}).
		First(&result)

	if stmt.Error != nil {
		if errors.Is(stmt.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(stmt.Error, "postgresql.GetOneByEmail Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}
