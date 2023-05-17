package repositories

import (
	"context"
	"database/sql"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
)

type postgresql struct {
	Sql *sql.DB
	Log log.Log
}

func (r *postgresql) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	return nil, nil
}
