package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initPostgreSQLMock(t *testing.T) (sqlmock.Sqlmock, *postgresql, context.Context) {
	// Create a GORM DB instance with the mock DB
	mock, gormDB := mockdata.NewMySQLMock(t)

	// Create a new instance of the mysql repository with the GORM DB and a mock logger
	repo := &postgresql{
		Sql: gormDB,
		Log: log.NewLog(),
	}

	// Create a new context
	ctx := context.Background()

	return mock, repo, ctx
}

// postgresql.GetCollection
func Test_postgresql_GetCollection(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rows := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, data.BankName, data.UserInserted)
		mock.ExpectQuery("SELECT \\* FROM `bank`").
			WillReturnRows(rows)

		// Call the GetCollection method
		result, err := repo.GetCollection(ctx)

		// // Assert that there is no error returned
		assert.NoError(t, err)

		// // Assert that the returned result is not nil
		assert.NotNil(t, result)

		// // Assert that the returned result matches the expected result
		assert.GreaterOrEqual(t, len(result), 1)

		assert.Equal(t, result[0].BankName, data.BankName)

		// // Assert that the expected methods of the mock objects were called
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		mock.ExpectQuery("SELECT \\* FROM `bank`").
			WillReturnError(errors.New("test Error"))

		// Call the GetCollection method
		result, err := repo.GetCollection(ctx)

		assert.Error(t, err)

		assert.Contains(t, err.Error(), "test Error")

		assert.Nil(t, result)
	})
}

// postgresql.Create
func Test_postgresql_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `bank`").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Call the Create method
		result, err := repo.Create(ctx, data)

		// Assert that there is no error returned
		assert.NoError(t, err)

		// Assert that the returned result is not nil
		assert.NotNil(t, result)

		assert.Equal(t, result.BankName, data.BankName)

		// Assert that the expected methods of the mock objects were called
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `bank`").
			WillReturnError(errors.New("duplicate primary key error"))
		mock.ExpectRollback()

		// Call the Create method
		result, err := repo.Create(ctx, data)

		// Assert that the returned result is nil
		assert.Nil(t, result)

		// Assert that there is error returned
		assert.Error(t, err)

		assert.Equal(t, err.Error(), "duplicate primary key error")
	})

}

// postgresql.GetOneByCode
func Test_postgresql_GetOneByCode(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rows := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, data.BankName, data.UserInserted)
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `bankCode` = \\?").
			WillReturnRows(rows)

		// Call the GetOneByCode method
		result, err := repo.GetOneByCode(ctx, data.BankCode)

		assert.NoError(t, err)

		assert.NotNil(t, result)

		assert.Equal(t, result.BankName, data.BankName)
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rows := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, data.BankName, data.UserInserted)
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `bankCode` = \\?").
			WillReturnRows(rows).
			WillReturnError(errors.New("test error"))

		// Call the GetOneByCode method
		result, err := repo.GetOneByCode(ctx, data.BankCode)

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "test error")

		assert.Nil(t, result)

	})

	t.Run("Success but No Data", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `bankCode` = \\?").
			WillReturnError(gorm.ErrRecordNotFound)

		// Call the GetOneByCode method
		result, err := repo.GetOneByCode(ctx, data.BankCode)

		assert.NoError(t, err)

		assert.Nil(t, result)
	})
}

// postgresql.GetOneByOid
func Test_postgresql_GetOneByOid(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rows := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, data.BankName, data.UserInserted)
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `oid` = \\?").
			WillReturnRows(rows)

		// Call the GetOneByOid method
		result, err := repo.GetOneByOid(ctx, data.Oid)

		assert.NoError(t, err)

		assert.NotNil(t, result)

		assert.Equal(t, result.BankName, data.BankName)
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rows := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, data.BankName, data.UserInserted)
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `oid` = \\?").
			WillReturnRows(rows).
			WillReturnError(errors.New("test error"))

		// Call the GetOneByOid method
		result, err := repo.GetOneByOid(ctx, data.Oid)

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "test error")

		assert.Nil(t, result)
	})

	t.Run("Success but No Data", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `oid` = \\?").
			WillReturnError(gorm.ErrRecordNotFound)

		// Call the GetOneByOid method
		result, err := repo.GetOneByOid(ctx, data.Oid)

		assert.NoError(t, err)

		assert.Nil(t, result)
	})
}

// postgresql.UpdateOne
func Test_postgresql_UpdateOne(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		// Set the expected behavior for the mock DB
		rowsUpdated := sqlmock.NewRows([]string{"oid", "BankCode", "bankName", "userInserted"}).
			AddRow(data.Oid, data.BankCode, mockdata.BankMock[1].BankName, data.UserInserted)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `bank` SET `bankName`=\\? WHERE oid = \\?").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		mock.ExpectQuery("SELECT \\* FROM `bank` WHERE `oid` = \\?").
			WillReturnRows(rowsUpdated)

		mock.ExpectationsWereMet()

		// Call the UpdateOne method
		result, err := repo.UpdateOne(ctx, data, mockdata.BankMock[1])

		assert.NoError(t, err)

		assert.NotNil(t, result)

		assert.Equal(t, result.BankName, mockdata.BankMock[1].BankName)
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `bank` SET `bankName`=\\? WHERE oid = \\?").
			WillReturnError(errors.New("test error update"))
		mock.ExpectRollback()

		// Call the UpdateOne method
		result, err := repo.UpdateOne(ctx, data, mockdata.BankMock[1])

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "test error update")

		assert.Nil(t, result)
	})
}

// postgresql.Delete
func Test_postgresql_Delete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `bank` WHERE `bank`.`oid` = ?").
			WithArgs(data.Oid).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		// Call the Delete method
		err := repo.Delete(ctx, data)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mock, repo, ctx := initPostgreSQLMock(t)

		// Create a sample Bank object
		data := mockdata.BankMock[0]

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `bank` WHERE `bank`.`oid` = ?").
			WithArgs(data.Oid).
			WillReturnError(errors.New("test delete error"))
		mock.ExpectRollback()

		// Call the Delete method
		err := repo.Delete(ctx, data)

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "test delete error")
	})
}
