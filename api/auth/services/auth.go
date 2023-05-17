package services

import (
	"context"
	"time"

	"github.com/abelz123456/celestial-api/api/auth/domain"
	"github.com/abelz123456/celestial-api/api/auth/repositories"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/google/uuid"
)

type service struct {
	Repository domain.Repository
}

func NewService(mgr manager.Manager) domain.Service {
	return &service{
		Repository: repositories.NewRepository(mgr),
	}
}

func (s *service) Register(ctx context.Context, data domain.PayloadRegister) (*entity.PermissionPolicyUser, error) {
	passwordHash, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	result := entity.PermissionPolicyUser{
		Oid:                 uuid.New().String(),
		EmailName:           data.EmailName,
		Password:            passwordHash,
		Description:         new(string),
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        new(string),
		InsertedDate:        time.Now(),
		LastUserId:          new(string),
		LastUpdate:          time.Now(),
	}

	if _, err := s.Repository.Save(ctx, result); err != nil {
		return nil, err
	}

	return &result, nil
}
