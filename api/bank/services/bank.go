package services

import (
	"context"
	"fmt"
	"time"

	"github.com/abelz123456/celestial-api/api/bank/domain"
	"github.com/abelz123456/celestial-api/api/bank/repositories"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/google/uuid"
)

type service struct {
	manager    manager.Manager
	repository domain.Repository
}

func NewService(mgr manager.Manager) domain.Service {
	return &service{
		manager:    mgr,
		repository: repositories.NewRepository(mgr),
	}
}

func (s *service) GetList(ctx context.Context) ([]entity.Bank, error) {
	return s.repository.GetCollection(ctx)
}

func (s *service) CreateNew(ctx context.Context, data domain.CreateBankDto) (*entity.Bank, error) {
	bank, err := s.repository.GetOneByCode(ctx, data.BankCode)
	if err != nil {
		return nil, err
	}

	if bank != nil {
		return nil, fmt.Errorf("code '%s' is registered", bank.BankCode)
	}

	bank = &entity.Bank{
		Oid:          uuid.NewString(),
		BankCode:     data.BankCode,
		BankName:     data.BankName,
		UserInserted: data.UserInserted,
		InsertedDate: time.Now(),
		LastUserId:   "",
		LastUpdate:   time.Now(),
	}

	result, err := s.repository.Create(ctx, *bank)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetOne(ctx context.Context, oid string) (*entity.Bank, error) {
	bank, err := s.repository.GetOneByOid(ctx, oid)
	if err != nil {
		return nil, err
	}

	if bank != nil {
		return bank, nil
	}

	return nil, nil
}

func (s *service) UpdateOne(ctx context.Context, oid string, data domain.UpdateBankDto) (*entity.Bank, error) {
	bank, err := s.repository.GetOneByOid(ctx, oid)
	if err != nil || bank == nil {
		if bank == nil {
			return nil, fmt.Errorf("no data bank found with oid '%s'", oid)
		}
		return nil, err
	}

	newData := entity.Bank{
		BankName:   data.BankName,
		LastUpdate: time.Now(),
	}

	return s.repository.UpdateOne(ctx, *bank, newData)
}

func (s *service) Delete(ctx context.Context, oid string) error {
	bank, err := s.repository.GetOneByOid(ctx, oid)
	if err != nil {
		return err
	}

	if bank == nil {
		return fmt.Errorf("no data bank found with oid '%s'", oid)
	}

	return s.repository.Delete(ctx, *bank)
}
