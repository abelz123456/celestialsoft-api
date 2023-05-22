package services

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/api/rajaongkir/repositories"
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

func (s *service) GetProvince(ctx context.Context) (interface{}, error) {
	return s.repository.GetRajaongkirProvince(ctx)
}

func (s *service) GetCity(ctx context.Context, provinceID string) (interface{}, error) {
	return s.repository.GetRajaongkirCity(ctx, provinceID)
}

func (s *service) GetCostInfo(ctx context.Context, deliveryData domain.CostInfoPayload) (*entity.Rajaongkir, error) {
	now := time.Now()

	input, _ := json.Marshal(deliveryData)
	hash := sha1.New()
	io.WriteString(hash, string(input))
	hashBytes := hash.Sum(nil)
	hashData := fmt.Sprintf("%x", hashBytes)

	prevData, err := s.repository.GetOneByHashData(ctx, hashData)
	if err != nil || prevData != nil {
		if prevData != nil {
			var apiResult map[string]interface{}
			json.Unmarshal([]byte(prevData.APIResponse), &apiResult)
			prevData.Response = apiResult
			return prevData, nil
		}

		return nil, err
	}

	apiResult, err := s.repository.GetRajaongkirCost(ctx, deliveryData)
	if err != nil {
		return nil, err
	}

	apiStatus := 0
	if status, ok := apiResult["status"]; ok {
		if mapStatus, ok := status.(map[string]interface{}); ok {
			if code, ok := mapStatus["code"]; ok {
				if val, ok := code.(float64); ok {
					apiStatus = int(val)
				}
			}
		}
	}

	mapResult, _ := json.Marshal(apiResult)
	rajaongkir := entity.Rajaongkir{
		UID:         uuid.NewString(),
		HashData:    hashData,
		Origin:      deliveryData.Origin,
		Destination: deliveryData.Destination,
		Wight:       deliveryData.Weight,
		Courier:     deliveryData.Courier,
		APIResponse: string(mapResult),
		CreatedAt:   &now,
		CreatedBy:   deliveryData.CreatedBy,
		ApiStatus:   apiStatus,
	}

	if err := s.repository.Save(ctx, rajaongkir); err != nil {
		return nil, err
	}

	rajaongkir.Response = apiResult
	return &rajaongkir, nil
}

func (s *service) GetCostHistories(ctx context.Context) ([]entity.Rajaongkir, error) {
	results, err := s.repository.GetCollection(ctx)
	if err != nil {
		return nil, err
	}

	for i, result := range results {
		var apiResult map[string]interface{}
		json.Unmarshal([]byte(result.APIResponse), &apiResult)
		results[i].Response = apiResult
	}

	return results, nil
}
