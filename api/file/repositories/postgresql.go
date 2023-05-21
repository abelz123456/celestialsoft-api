package repositories

import (
	"context"
	"errors"
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
		r.Log.Error(err, "postgresql.SaveLocalStorage Exception", nil)
		return err
	}
	defer file.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		r.Log.Error(err, "postgresql.SaveLocalStorage Exception", nil)
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		r.Log.Error(err, "postgresql.SaveLocalStorage Exception", nil)
		return err
	}

	return nil
}

func (r *postgresql) DeleteLocalStorage(ctx context.Context, filePath string) error {
	return os.Remove(filePath)
}

func (r *postgresql) Save(ctx context.Context, fileData entity.LocalFile) error {
	tx := r.Sql.WithContext(ctx).Begin()

	if err := tx.Create(&fileData).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "postgresql.Save Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.LocalFile, error) {
	var (
		results = make([]entity.LocalFile, 0)
		tx      = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.LocalFile{}).Find(&results).Error; err != nil {
		r.Log.Error(err, "postgresql.GetCollection Exception", nil)
		return nil, err
	}

	return results, nil
}

func (r *postgresql) GetOneByUID(ctx context.Context, uid string) (*entity.LocalFile, error) {
	var (
		result entity.LocalFile
		tx     = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.LocalFile{}).Where(map[string]interface{}{"uid": uid}).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(err, "postgresql.GetOneByUID Exception", nil)
		return nil, err
	}

	return &result, nil
}

func (r *postgresql) UpdateOne(ctx context.Context, data entity.LocalFile, newData entity.LocalFile) error {
	var tx = r.Sql.WithContext(ctx).Begin()

	stmt := tx.Model(&entity.LocalFile{}).
		Where(map[string]string{"uid": data.UID}).
		Updates(map[string]interface{}{"localPath": newData.LocalPath, "originalName": newData.OriginalName, "updatedAt": newData.UpdatedAt})

	if err := stmt.Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "postgresql.UpdateOne Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *postgresql) Delete(ctx context.Context, fileData entity.LocalFile) error {
	var tx = r.Sql.WithContext(ctx).Begin()

	if err := tx.Model(&entity.LocalFile{}).Delete(fileData).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "postgresql.Delete Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}
