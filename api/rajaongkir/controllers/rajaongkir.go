package controllers

import (
	"strings"

	"github.com/abelz123456/celestial-api/api/rajaongkir/domain"
	"github.com/abelz123456/celestial-api/api/rajaongkir/services"
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

// @Tags		Rajaongkir
// @Router		/rajaongkir [get]
// @Summary		Get Rajaongkir Cost histories
// @Description	Get Rajaongkir Cost histories
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
func (c *controller) GetHistories(ctx *gin.Context) {
	result, err := c.service.GetCostHistories(ctx.Request.Context())
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Error handle get histories", map[string]interface{}{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		Rajaongkir
// @Router		/rajaongkir/province [get]
// @Summary		Get Rajaongkir Province data
// @Description	Get Rajaongkir Province data
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
func (c *controller) GetProvince(ctx *gin.Context) {
	result, err := c.service.GetProvince(ctx.Request.Context())
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Error handle get province", map[string]interface{}{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		Rajaongkir
// @Router		/rajaongkir/province/{id}/city [get]
// @Summary		Get Rajaongkir City data by Province ID
// @Description	Get Rajaongkir City data by Province ID
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param		id path string true "number of Province ID"
// @Produce 	application/json
func (c *controller) GetCity(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	results, err := c.service.GetCity(ctx.Request.Context(), id)
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Error handle get city", map[string]interface{}{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", results)
}

// @Tags		Rajaongkir
// @Router		/rajaongkir/cost [post]
// @Summary		Get delivery cost with Rajaongkir
// @Description	Get delivery cost with Rajaongkir
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param 		data body domain.CostInfoPayload true "Rajaongkir cost data"
// @Produce 	application/json
func (c *controller) CostInfo(ctx *gin.Context) {
	var (
		deliveryData domain.CostInfoPayload
		userOid      = ctx.GetString("oid")
	)

	if err := ctx.ShouldBindJSON(&deliveryData); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", err)
		return
	}

	if err := c.Validator.Struct(deliveryData); err != nil {
		mapError := helpers.ValidationErrorToMap(err)
		response.SendJson(ctx, response.ErrForm1Forbidden, "", mapError)
		return
	}

	deliveryData.CreatedBy = userOid
	deliveryData.Courier = strings.ToLower(deliveryData.Courier)
	result, err := c.service.GetCostInfo(ctx.Request.Context(), deliveryData)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedGetCost, "", err)
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}
