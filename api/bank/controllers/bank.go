package controllers

import (
	"github.com/abelz123456/celestial-api/api/bank/domain"
	"github.com/abelz123456/celestial-api/api/bank/services"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/api/response"
	"github.com/abelz123456/celestial-api/utils/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service   domain.Service
	Validator *validator.Validate
}

func NewController(mgr manager.Manager) domain.Controller {
	return &controller{
		service:   services.NewService(mgr),
		Validator: validator.New(),
	}
}

func (c *controller) GetList(ctx *gin.Context) {
	result, err := c.service.GetList(ctx.Request.Context())
	if err != nil {
		response.SendJson(ctx, response.ErrFailedGetBankCollection, "", nil)
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

func (c *controller) CreateNew(ctx *gin.Context) {
	var data domain.CreateBankDto

	if err := ctx.ShouldBindJSON(&data); err != nil {
		response.SendJson(ctx, response.ErrFailedCreateBank, err.Error(), nil)
		return
	}

	if err := c.Validator.Struct(data); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", helpers.ValidationErrorToMap(err))
		return
	}

	data.UserInserted = ctx.GetString("oid")
	result, err := c.service.CreateNew(ctx, data)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedCreateBank, err.Error(), nil)
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

func (c *controller) GetOne(ctx *gin.Context) {
	oid, _ := ctx.Params.Get("oid")

	result, err := c.service.GetOne(ctx, oid)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedGetBank, "", map[string]string{"oid": oid, "error": err.Error()})
		return
	}

	if result == nil {
		response.SendJson(ctx, response.ErrFailedGetBank, "Bank data not found", map[string]string{"oid": oid, "error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

func (c *controller) UpdateOne(ctx *gin.Context) {
	var data domain.UpdateBankDto
	oid, _ := ctx.Params.Get("oid")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		response.SendJson(ctx, response.ErrFailedUpdateBank, err.Error(), nil)
		return
	}

	if err := c.Validator.Struct(data); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", helpers.ValidationErrorToMap(err))
		return
	}

	result, err := c.service.UpdateOne(ctx, oid, data)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedUpdateBank, "", map[string]string{"oid": oid, "error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", map[string]interface{}{"newData": result})
}

func (c *controller) Delete(ctx *gin.Context) {
	oid, _ := ctx.Params.Get("oid")

	if err := c.service.Delete(ctx.Request.Context(), oid); err != nil {
		response.SendJson(ctx, response.ErrFailedRemoveBank, "", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "Bank data deleted", map[string]string{"oid": oid})
}
