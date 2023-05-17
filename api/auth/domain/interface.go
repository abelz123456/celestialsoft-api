package domain

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type Service interface {
	Register(ctx context.Context, data PayloadRegister) (*entity.PermissionPolicyUser, error)
}

type Repository interface {
	Save(ctx context.Context, data entity.PermissionPolicyUser) (*entity.PermissionPolicyUser, error)
}
