package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type mysql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *mysql) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	stmt := r.Sql.WithContext(ctx).
		Create(&data)

	if stmt.Error != nil {
		stmt.Rollback()
		r.Log.Error(stmt.Error, "mysql.Save Exception", nil)
		return nil, stmt.Error
	}

	return &data, nil
}

func (r *mysql) GetOneByEmail(ctx context.Context, email string) (*entity.PermissionPolicyUser, error) {
	var result entity.PermissionPolicyUser
	query := "select * from permissionPolicyUser where emailName = ?"
	stmt := r.Sql.WithContext(ctx).
		Raw(query, email).
		Scan(&result)

	if stmt.Error != nil {
		r.Log.Error(stmt.Error, "mysql.GetOneByEmail Exception", nil)
		return nil, stmt.Error
	}

	if stmt.RowsAffected > 0 {
		return &result, nil
	}

	return nil, nil
}
