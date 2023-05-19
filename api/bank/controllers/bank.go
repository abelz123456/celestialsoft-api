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

// @Tags		Bank
// @Router		/bank [get]
// @Summary		Get Bank Collection
// @Description	Get Bank Collection
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Success		200 {object} response.ResponseProperties "Ok"
// @Failure 	400 {object} response.ResponseProperties "General Error"
// @Failure 	401 {object} response.ResponseProperties "Authentication Error"
func (c *controller) GetList(ctx *gin.Context) {
	result, err := c.service.GetList(ctx.Request.Context())
	if err != nil {
		response.SendJson(ctx, response.ErrFailedGetBankCollection, "", nil)
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		Bank
// @Router		/bank [post]
// @Summary		Create New Bank
// @Description	Create New Bank
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
// @Param 		data body domain.CreateBankDto true "Bank Information"
// @Success		201 {object} response.ResponseProperties "Created"
// @Failure 	400 {object} response.ResponseProperties "General Error"
// @Failure 	401 {object} response.ResponseProperties "Authentication Error"
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

// @Tags		Bank
// @Router		/bank/{oid} [get]
// @Summary		Detail Bank
// @Description	Detail Bank
// @Param		oid path string true "oid of Bank"
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Success		200 {object} response.ResponseProperties "OK"
// @Failure 	400 {object} response.ResponseProperties "General Error"
// @Failure 	401 {object} response.ResponseProperties "Authentication Error"
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

// @Tags		Bank
// @Router		/bank/{oid} [patch]
// @Summary		Update Bank
// @Description	Update Bank
// @Param		oid path string true "oid of Bank"
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
// @Param 		data body domain.UpdateBankDto true "New Bank Information"
// @Success		200 {object} response.ResponseProperties "OK"
// @Failure 	400 {object} response.ResponseProperties "General Error"
// @Failure 	401 {object} response.ResponseProperties "Authentication Error"
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

// @Tags		Bank
// @Router		/bank/{oid} [delete]
// @Summary		Delete Bank
// @Description	Detail Bank
// @Param		oid path string true "oid of Bank"
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Success		200 {object} response.ResponseProperties "OK"
// @Failure 	400 {object} response.ResponseProperties "General Error"
// @Failure 	401 {object} response.ResponseProperties "Authentication Error"
func (c *controller) Delete(ctx *gin.Context) {
	oid, _ := ctx.Params.Get("oid")

	if err := c.service.Delete(ctx.Request.Context(), oid); err != nil {
		response.SendJson(ctx, response.ErrFailedRemoveBank, "", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "Bank data deleted", map[string]string{"oid": oid})
}
