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

func (r *postgresql) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	return nil, nil
}

func (r *postgresql) GetOneByEmail(ctx context.Context, email string) (*entity.PermissionPolicyUser, error) {

	return nil, nil
}
