package helpers

import (
	"time"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	CreateToken(oid string) string
	ParseToken(token string) string
}

type option struct {
	Config    config.Config
	secretKey string
	exp       int
	logger    log.Log
}

func NewJwtHelper(cfg config.Config) Jwt {
	return &option{
		Config:    cfg,
		secretKey: cfg.SecretKey,
		exp:       cfg.JwtExpiredTime,
		logger:    log.NewLog(),
	}
}

func (o *option) CreateToken(oid string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": oid,
			"exp": time.Now().Add(time.Minute * time.Duration(o.exp)).Unix(),
		})

	token, err := t.SignedString([]byte(o.secretKey))
	if err != nil {
		o.logger.Warning("Failed NewJwtHelper.CreateToken", err, nil)
		return ""
	}

	return token
}

func (o *option) ParseToken(token string) string {
	parsed, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(o.secretKey), nil
	})

	if err != nil || !parsed.Valid {
		return ""
	}

	sub, _ := parsed.Claims.GetSubject()
	return sub
}
