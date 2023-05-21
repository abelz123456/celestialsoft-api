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

type mysql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *mysql) SaveLocalStorage(ctx context.Context, fileHeader multipart.FileHeader, destination string) error {
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

func (r *mysql) DeleteLocalStorage(ctx context.Context, filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysql) Save(ctx context.Context, fileData entity.LocalFile) error {
	tx := r.Sql.WithContext(ctx).Begin()

	if err := tx.Create(&fileData).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "mysql.Save Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *mysql) GetCollection(ctx context.Context) ([]entity.LocalFile, error) {
	var (
		results = make([]entity.LocalFile, 0)
		tx      = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.LocalFile{}).Find(&results).Error; err != nil {
		r.Log.Error(err, "mysql.GetCollection Exception", nil)
		return nil, err
	}

	return results, nil
}

func (r *mysql) GetOneByUID(ctx context.Context, uid string) (*entity.LocalFile, error) {
	var (
		result entity.LocalFile
		tx     = r.Sql.WithContext(ctx).Begin()
	)

	if err := tx.Model(&entity.LocalFile{}).Where(map[string]interface{}{"uid": uid}).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Log.Error(err, "mysql.GetOneByUID Exception", nil)
		return nil, err
	}

	return &result, nil
}

func (r *mysql) UpdateOne(ctx context.Context, data entity.LocalFile, newData entity.LocalFile) error {
	var tx = r.Sql.WithContext(ctx).Begin()

	stmt := tx.Model(&entity.LocalFile{}).
		Where(map[string]string{"uid": data.UID}).
		Updates(map[string]interface{}{"localPath": newData.LocalPath, "originalName": newData.OriginalName, "updatedAt": newData.UpdatedAt})

	if err := stmt.Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "mysql.UpdateOne Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r *mysql) Delete(ctx context.Context, fileData entity.LocalFile) error {
	var tx = r.Sql.WithContext(ctx).Begin()

	if err := tx.Model(&entity.LocalFile{}).Delete(fileData).Error; err != nil {
		tx.Rollback()
		r.Log.Error(err, "mysql.Delete Exception", nil)
		return err
	}

	tx.Commit()
	return nil
}
