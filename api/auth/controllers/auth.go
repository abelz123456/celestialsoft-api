package controllers

import (
	"errors"

	"github.com/abelz123456/celestial-api/api/auth/domain"
	"github.com/abelz123456/celestial-api/api/auth/mapping"
	"github.com/abelz123456/celestial-api/api/auth/services"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/api/response"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	Manager   manager.Manager
	Validator *validator.Validate
	Service   domain.Service
}

func NewController(mgr manager.Manager) domain.Controller {
	return &controller{
		Manager:   mgr,
		Validator: validator.New(),
		Service:   services.NewService(mgr),
	}
}

func (c *controller) Login(ctx *gin.Context) {
	c.Manager.Logger.Panic(errors.New("test fatal"), "", nil)
	response.SendJson(ctx, response.Ok, "", []map[string]interface{}{{"look": "at this.."}})
}

func (c *controller) Register(ctx *gin.Context) {
	var userDto domain.PayloadRegister

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", err)
		return
	}

	if err := c.Validator.Struct(userDto); err != nil {
		mapError := helpers.ValidationErrorToMap(err)
		response.SendJson(ctx, response.ErrForm1Forbidden, "", mapError)
		return
	}

	result, err := c.Service.Register(ctx.Request.Context(), userDto)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedRegister, "", err)
		return
	}

	response.SendJson(ctx, response.Created, "", mapping.ToPermissionPolicyUserResponse(*result))
}
