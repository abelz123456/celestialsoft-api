package repositories

import (
	"context"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"gorm.io/gorm"
)

type postgresql struct {
	Sql *gorm.DB
	Log log.Log
}

func (r *postgresql) Save(ctx context.Context, data entity.EmailSent) error {
	return nil
}

func (r *postgresql) GetOneByUID(ctx context.Context, uid string) (*entity.EmailSent, error) {
	return nil, nil
}

func (r *postgresql) GetCollection(ctx context.Context) ([]entity.EmailSent, error) {
	return nil, nil
}
