package repositories

import (
	"context"
	"database/sql"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
)

type mysql struct {
	Sql *sql.DB
	Log log.Log
}

func (r *mysql) Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error) {
	query := "insert into permissionPolicyUser(description, optimisticLockField, gCRecord, deleted, userInserted, insertedDate,oid,emailName,password) values (?,?,?,?,?,?,?,?,?)"
	stmt, err := r.Sql.Prepare(query)
	if err != nil {
		r.Log.Error(err, "mysql.Save Exception", nil)
		return nil, err
	}
	defer stmt.Close()

	args := []interface{}{
		data.Description,
		data.OptimisticLockField,
		data.GCRecord,
		data.Deleted,
		data.UserInserted,
		data.InsertedDate,
		data.Oid,
		data.EmailName,
		data.Password,
	}

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.Log.Error(err, "mysql.Save Exception", nil)
		return nil, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		r.Log.Error(err, "mysql.Save Exception", nil)
		return nil, err
	}

	return &data, nil
}
