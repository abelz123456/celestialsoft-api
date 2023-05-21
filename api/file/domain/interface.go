package domain

import (
	"context"
	"mime/multipart"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Upload(ctx *gin.Context)
	GetCollection(ctx *gin.Context)
	GetInfo(ctx *gin.Context)
	Replace(ctx *gin.Context)
	Unlink(ctx *gin.Context)
}

type Service interface {
	UploadFile(ctx context.Context, fileData UploadFileData) (*entity.LocalFile, error)
	GetCollection(ctx context.Context) ([]entity.LocalFile, error)
	GetInfo(ctx context.Context, uid string) (*entity.LocalFile, error)
	ReplaceFile(ctx context.Context, fileData ReplaceFileData) (*entity.LocalFile, error)
	UnlinkFile(ctx context.Context, uid string) error
}

type Repository interface {
	SaveLocalStorage(ctx context.Context, fileHeader multipart.FileHeader, destination string) error
	DeleteLocalStorage(ctx context.Context, filePath string) error

	Save(ctx context.Context, fileData entity.LocalFile) error
	GetCollection(ctx context.Context) ([]entity.LocalFile, error)
	GetOneByUID(ctx context.Context, uid string) (*entity.LocalFile, error)
	UpdateOne(ctx context.Context, data entity.LocalFile, newData entity.LocalFile) error
	Delete(ctx context.Context, fileData entity.LocalFile) error
}
