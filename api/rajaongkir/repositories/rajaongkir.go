package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/manager"
)

type repository struct {
	config     config.Config
	httpClient *http.Client
	log        log.Log
}

func NewRepository(mgr manager.Manager) domain.Repository {
	repo := &repository{
		config:     mgr.Config,
		httpClient: http.DefaultClient,
		log:        mgr.Logger,
	}

	switch mgr.Config.DBUsed {
	case database.MySQL.String():
		return &mysql{
			sql:  mgr.Database.Sql,
			log:  mgr.Logger,
			repo: repo,
		}
	case database.PostgreSQL.String():
		return &postgresql{
			sql:  mgr.Database.Sql,
			log:  mgr.Logger,
			repo: repo,
		}
	default:
		return &mongodb{
			mongo: mgr.Database.Mongo,
			log:   mgr.Logger,
			repo:  repo,
		}
	}
}

func (r *repository) getRajaongkirProvince(ctx context.Context) (interface{}, error) {
	var (
		apiResponse struct {
			Rajaongkir map[string]interface{} `json:"rajaongkir"`
		}
	)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/province", r.config.RajaongkirApiUrl), nil)
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirProvince Exception", nil)
		return nil, err
	}

	req.Header.Set("Key", r.config.RajaongkirApiKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirProvince Exception", nil)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		r.log.Error(err, "repository.getRajaongkirProvince Exception", nil)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		r.log.Warning("repository.getRajaongkirProvince http client warning", fmt.Errorf("%d http status code", resp.StatusCode), map[string]interface{}{
			"method":   req.Method,
			"url":      req.URL,
			"headers":  req.Header,
			"response": apiResponse.Rajaongkir,
		})
	}

	if _, exists := apiResponse.Rajaongkir["results"]; exists {
		return apiResponse.Rajaongkir["results"], nil
	}

	return apiResponse, nil
}

func (r *repository) getRajaongkirCity(ctx context.Context, provinceID string) (interface{}, error) {
	var (
		apiResponse struct {
			Rajaongkir map[string]interface{} `json:"rajaongkir"`
		}
	)

	u, err := url.Parse(fmt.Sprintf("%s/city", r.config.RajaongkirApiUrl))
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirCity Exception", nil)
		return nil, err
	}

	params := url.Values{}
	params.Set("province", provinceID)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirCity Exception", nil)
		return nil, err
	}

	req.Header.Set("Key", r.config.RajaongkirApiKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirCity Exception", nil)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		r.log.Error(err, "repository.getRajaongkirCity Exception", nil)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		r.log.Warning("repository.getRajaongkirCity http client warning", fmt.Errorf("%d http status code", resp.StatusCode), map[string]interface{}{
			"method":   req.Method,
			"url":      req.URL,
			"headers":  req.Header,
			"response": apiResponse.Rajaongkir,
		})
	}

	if _, exists := apiResponse.Rajaongkir["results"]; exists {
		return apiResponse.Rajaongkir["results"], nil
	}

	return apiResponse.Rajaongkir, nil
}

func (r *repository) getRajaongkirCost(ctx context.Context, deliveryData domain.CostInfoPayload) (map[string]interface{}, error) {
	var (
		apiResponse struct {
			Rajaongkir map[string]interface{} `json:"rajaongkir"`
		}

		data, _ = json.Marshal(&deliveryData)
	)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cost", r.config.RajaongkirApiUrl), bytes.NewBuffer(data))
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirCost Exception", nil)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Key", r.config.RajaongkirApiKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log.Error(err, "repository.getRajaongkirCost Exception", nil)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		r.log.Error(err, "repository.getRajaongkirCost Exception", nil)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		r.log.Warning("repository.getRajaongkirCost http client warning", fmt.Errorf("%d http status code", resp.StatusCode), map[string]interface{}{
			"method":   req.Method,
			"url":      req.URL,
			"headers":  req.Header,
			"body":     string(data),
			"response": apiResponse.Rajaongkir,
		})
	}

	return apiResponse.Rajaongkir, nil
}
