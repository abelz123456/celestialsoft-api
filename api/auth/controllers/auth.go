package controllers

import (
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

// @Tags		Authentication
// @Router		/auth/login [post]
// @Summary		User Authentication
// @Produce 	application/json
// @Param 		data body domain.PayloadLogin true "Login info"
// @Failure 	401 {object} response.ResponseProperties
func (c *controller) Login(ctx *gin.Context) {
	var loginInfo domain.PayloadLogin

	if err := ctx.ShouldBindJSON(&loginInfo); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", err)
		return
	}

	if err := c.Validator.Struct(loginInfo); err != nil {
		mapError := helpers.ValidationErrorToMap(err)
		response.SendJson(ctx, response.ErrForm1Forbidden, "", mapError)
		return
	}

	result, err := c.Service.Login(ctx.Request.Context(), loginInfo)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedLogin, "", err)
		return
	}

	if result == nil {
		response.SendJson(ctx, response.ErrFailedLogin, "Incorrect email or password", err)
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		Authentication
// @Router		/auth/register [post]
// @Summary		User Registration
// @Produce 	application/json
// @Param 		data body domain.PayloadRegister true "Login info"
// @Failure 	400 {object} response.ResponseProperties
// @Failure 	403 {object} response.ResponseProperties
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
