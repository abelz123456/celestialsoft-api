package services

import (
	"context"
	"fmt"
	"time"

	"github.com/abelz123456/celestial-api/api/auth/domain"
	"github.com/abelz123456/celestial-api/api/auth/repositories"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/google/uuid"
)

type service struct {
	manager    manager.Manager
	Repository domain.Repository
}

func NewService(mgr manager.Manager) domain.Service {
	return &service{
		manager:    mgr,
		Repository: repositories.NewRepository(mgr),
	}
}

func (s *service) Register(ctx context.Context, data domain.PayloadRegister) (*entity.PermissionPolicyUser, error) {
	user, err := s.Repository.GetOneByEmail(ctx, data.EmailName)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("'%s' already in use", data.EmailName)
	}

	passwordHash, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	result := entity.PermissionPolicyUser{
		Oid:                 uuid.New().String(),
		EmailName:           data.EmailName,
		Password:            passwordHash,
		Description:         "",
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        "",
		InsertedDate:        time.Now(),
		LastUserId:          "",
		LastUpdate:          time.Now(),
	}

	if _, err := s.Repository.Save(ctx, result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *service) Login(ctx context.Context, data domain.PayloadLogin) (*entity.PermissionPolicyUser, error) {
	user, err := s.Repository.GetOneByEmail(ctx, data.EmailName)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if helpers.CheckPasswordHash(data.Password, user.Password) {
		jwt := helpers.NewJwtHelper(s.manager.Config)
		user.AuthToken = jwt.CreateToken(user.Oid)
		return user, nil
	}

	return nil, nil
}
