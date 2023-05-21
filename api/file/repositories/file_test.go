package repositories

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/test/mockdata"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

type testFile struct {
	Fieldname string
	Filename  string
	Content   []byte
}

func createRequestMultipartFiles(t *testing.T) *http.Request {
	var (
		body  bytes.Buffer
		files = []testFile{
			{"content", "test.jpg", []byte("hello")},
		}
	)

	mw := multipart.NewWriter(&body)
	for _, file := range files {
		fw, err := mw.CreateFormFile(file.Fieldname, file.Filename)
		assert.NoError(t, err)

		n, err := fw.Write(file.Content)
		assert.NoError(t, err)
		assert.Equal(t, len(file.Content), n)
	}
	err := mw.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/", &body)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", binding.MIMEMultipartPOSTForm+"; boundary="+mw.Boundary())
	return req
}

func TestNewMySQLRepository(t *testing.T) {
	mysqlMgr := mockdata.NewFakeManager(t, database.MySQL)
	NewRepository(mysqlMgr)
}

func TestNewPostgreSQLRepository(t *testing.T) {
	postgresqlMgr := mockdata.NewFakeManager(t, database.PostgreSQL)
	NewRepository(postgresqlMgr)
}

func TestNewMongoDBRepository(t *testing.T) {
	mongodbMgr := mockdata.NewFakeManager(t, database.Mongo)
	NewRepository(mongodbMgr)
}
