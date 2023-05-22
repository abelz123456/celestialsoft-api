package domain

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Send(ctx *gin.Context)
	InfoByUID(ctx *gin.Context)
	EmailSentCollection(ctx *gin.Context)
}

type Service interface {
	SendEmail(ctx context.Context, data SendEmailPayload) (*entity.EmailSent, error)
	GetCollection(ctx context.Context) ([]entity.EmailSent, error)
	GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error)
}

type Repository interface {
	Save(ctx context.Context, data entity.EmailSent) error
	GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error)
	GetCollection(ctx context.Context) ([]entity.EmailSent, error)
}
