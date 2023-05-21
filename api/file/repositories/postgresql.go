package repositories

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type postgresql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *postgresql) SaveLocalStorage(ctx context.Context, fileHeader multipart.FileHeader, destination string) error {
	file, err := fileHeader.Open()
	if err != nil {
		r.Log.Error(err, "mysql.SaveLocalStorage Exception", nil)
		return err
	}
	defer file.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		r.Log.Error(err, "mysql.SaveLocalStorage Exception", nil)
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		r.Log.Error(err, "mysql.SaveLocalStorage Exception", nil)
		return err
	}

	return nil
}

func (r *postgresql) DeleteLocalStorage(ctx context.Context, filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresql) Save(ctx context.Context, fileData entity.LocalFile) error {
	return nil
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.LocalFile, error) {
	return []entity.LocalFile{}, nil
}

func (r *postgresql) GetOneByUID(ctx context.Context, uid string) (*entity.LocalFile, error) {
	return nil, nil
}

func (r *postgresql) UpdateOne(ctx context.Context, data entity.LocalFile, newData entity.LocalFile) error {
	return nil
}

func (r *postgresql) Delete(ctx context.Context, fileData entity.LocalFile) error {
	return nil
}
