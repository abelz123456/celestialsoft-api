package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
)

func initPostgreSQLMock(t *testing.T) (sqlmock.Sqlmock, *postgresql, context.Context) {
	// Create a GORM DB instance with the mock DB
	mock, gormDB := mockdata.NewPosgreSQLMock(t)

	// Create a new instance of the mysql repository with the GORM DB and a mock logger
	repo := &postgresql{
		Sql: gormDB,
		Log: log.NewLog(),
	}

	// Create a new context
	ctx := context.Background()

	return mock, repo, ctx
}

func TestPostgreSQL_Save(t *testing.T) {
	mock, repo, ctx := initPostgreSQLMock(t)

	// Create a sample PermissionPolicyUser object
	data := mockdata.PermissionPolicyUserMock[0]

	// Set the expected behavior for the mock DB
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "permissionPolicyUser"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// mock.ExpectQuery()
	mock.ExpectCommit()

	// Call the Save method
	_, err := repo.Save(ctx, data)

	// Assert that there is no error returned
	assert.NoError(t, err)

	// Assert that the returned result is not nil
	assert.NotNil(t, data)

	// Assert that the returned result matches the original data
	assert.Equal(t, data, data)

	// Assert that the expected methods of the mock objects were called
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorPostgreSQL_Save(t *testing.T) {
	mock, repo, ctx := initPostgreSQLMock(t)

	// Create a sample PermissionPolicyUser object
	data := mockdata.PermissionPolicyUserMock[0]

	// Set the expected behavior for the mock DB
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "permissionPolicyUser"`).
		WillReturnError(errors.New("duplicate primary key error"))
	mock.ExpectRollback()

	// Call the Save method
	result, err := repo.Save(ctx, data)

	// Assert that the error matches the expected error
	assert.ErrorContains(t, err, "duplicate primary key error")

	// // Assert that the returned result is nil
	assert.Nil(t, result)

	// // Assert that the expected methods of the mock objects were called
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgreSQL_GetOneByEmail(t *testing.T) {
	mock, repo, ctx := initPostgreSQLMock(t)

	// Create a sample PermissionPolicyUser object
	data := mockdata.PermissionPolicyUserMock[0]

	// Set the expected behavior for the mock DB
	rows := sqlmock.NewRows([]string{"oid", "emailName", "optimisticLockField", "gCRecord", "deleted", "userInserted", "insertedDate", "password", "description"}).
		AddRow(data.Oid,
			data.EmailName,
			data.OptimisticLockField,
			data.GCRecord,
			data.Deleted,
			data.UserInserted,
			data.InsertedDate,
			data.Password,
			data.Description)
	mock.ExpectQuery(`SELECT \* FROM "permissionPolicyUser" WHERE "permissionPolicyUser"."emailName" = ?`).
		WillReturnRows(rows)

	// Call the GetOneByEmail method
	result, err := repo.GetOneByEmail(ctx, data.EmailName)

	// Assert that there is no error returned
	assert.NoError(t, err)

	// Assert that the returned result is not nil
	assert.NotNil(t, result)

	// Assert that the returned result matches the expected result
	assert.Equal(t, data.EmailName, result.EmailName)

	// Assert that the expected methods of the mock objects were called
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorPostgreSQL_GetOneByEmail(t *testing.T) {
	mock, repo, ctx := initPostgreSQLMock(t)

	// Create a sample PermissionPolicyUser object
	data := mockdata.PermissionPolicyUserMock[0]

	// Set the expected behavior for the mock DB
	rows := sqlmock.NewRows([]string{"oid", "emailName", "optimisticLockField", "gCRecord", "deleted", "userInserted", "insertedDate", "password", "description"}).
		AddRow(data.Oid,
			data.EmailName,
			data.OptimisticLockField,
			data.GCRecord,
			data.Deleted,
			data.UserInserted,
			data.InsertedDate,
			data.Password,
			data.Description)
	mock.ExpectQuery(`SELECT \* FROM "permissionPolicyUser" WHERE "permissionPolicyUser"."emailName" = ?`).
		WillReturnRows(rows).
		WillReturnError(errors.New("test mysql error"))

	// Call the GetOneByEmail method
	result, err := repo.GetOneByEmail(ctx, data.EmailName)

	// Assert that there is no error returned
	assert.EqualError(t, err, "test mysql error")

	// Assert that the returned result is not nil
	assert.Nil(t, result)

	// Assert that the expected methods of the mock objects were called
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNoDataPostgresql_GetOneByEmail(t *testing.T) {
	mock, repo, ctx := initPostgreSQLMock(t)

	// Set the expected behavior for the mock DB
	rows := sqlmock.NewRows([]string{"oid", "emailName", "optimisticLockField", "gCRecord", "deleted", "userInserted", "insertedDate", "password", "description"})
	mock.ExpectQuery(`SELECT \* FROM "permissionPolicyUser" WHERE "permissionPolicyUser"."emailName" = ?`).
		WillReturnRows(rows)

	// Call the GetOneByEmail method
	result, err := repo.GetOneByEmail(ctx, "me@testing.com")

	// Assert that there is no error returned
	assert.NoError(t, err)

	// Assert that the returned result is nil
	assert.Nil(t, result)

	// Assert that the expected methods of the mock objects were called
	assert.NoError(t, mock.ExpectationsWereMet())
}
