package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/abelz123456/celestial-api/api/mail/domain"
	"github.com/abelz123456/celestial-api/api/mail/repositories"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/google/uuid"
)

type service struct {
	config     config.Config
	repository domain.Repository
}

func NewService(mgr manager.Manager) domain.Service {
	return &service{
		config:     mgr.Config,
		repository: repositories.NewRepository(mgr),
	}
}

func (s *service) SendEmail(ctx context.Context, data domain.SendEmailPayload) (*entity.EmailSent, error) {
	for i, recipient := range data.Recipient {
		recipient = strings.TrimSpace(recipient)
		if !helpers.IsValidEmail(recipient) {
			return nil, fmt.Errorf("'%s' is not a valid email", recipient)
		}

		data.Recipient[i] = recipient
	}

	mail := helpers.NewMailHelper(s.config)
	sendErr := mail.Send(data.Recipient, data.Subject, data.Body)

	var payload map[string]interface{}
	jsonData, _ := json.Marshal(data)
	json.Unmarshal(jsonData, &payload)

	var (
		now      time.Time
		strError = ""
	)

	if sendErr == nil {
		now = time.Now()
	} else {
		strError = sendErr.Error()
	}

	emailSent := entity.EmailSent{
		UID:              uuid.NewString(),
		SentBy:           data.SentBy,
		Payload:          payload,
		StringifyPayload: string(jsonData),
		SentAt:           &now,
		SentError:        strError,
	}

	if err := s.repository.Save(ctx, emailSent); err != nil {
		return nil, err
	}

	return &emailSent, nil
}

func (s *service) GetCollection(ctx context.Context) ([]entity.EmailSent, error) {
	results, err := s.repository.GetCollection(ctx)
	if err != nil {
		return results, err
	}

	for i, result := range results {
		var payload map[string]interface{}
		json.Unmarshal([]byte(result.StringifyPayload), &payload)
		results[i].Payload = payload
	}

	return results, nil
}

func (s *service) GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error) {
	result, err := s.repository.GetOneByUID(ctx, uid)
	if err != nil || result == nil {
		if result == nil {
			err = fmt.Errorf("no email data with uid '%s'", uid)
		}

		return nil, err
	}

	var payload map[string]interface{}
	json.Unmarshal([]byte(result.StringifyPayload), &payload)
	result.Payload = payload

	return result, nil
}
