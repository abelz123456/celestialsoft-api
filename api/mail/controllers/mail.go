package controllers

import (
	"github.com/abelz123456/celestial-api/api/mail/domain"
	"github.com/abelz123456/celestial-api/api/mail/services"
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

// @Tags		Email
// @Router		/mail [post]
// @Summary		Send Email API
// @Description	Send Email API
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param 		data body domain.SendEmailPayload true "Send Email Data"
// @Produce 	application/json
func (c *controller) Send(ctx *gin.Context) {
	var data domain.SendEmailPayload

	if err := ctx.ShouldBindJSON(&data); err != nil {
		response.SendJson(ctx, response.ErrFailedCreateBank, err.Error(), nil)
		return
	}

	if err := c.Validator.Struct(data); err != nil {
		response.SendJson(ctx, response.ErrForm1Forbidden, "", helpers.ValidationErrorToMap(err))
		return
	}

	data.SentBy = ctx.GetString("oid")
	result, err := c.service.SendEmail(ctx, data)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedSendEmail, "", map[string]interface{}{"payload": data, "error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "Email sent", result)
}

// @Tags		Email
// @Router		/mail/{uid} [get]
// @Summary		Get Info Email History
// @Description	Get Info Email History
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param		uid path string true "String of Email History UID"
// @Produce 	application/json
func (c *controller) InfoByUID(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	result, err := c.service.GetOneByUID(ctx.Request.Context(), uid)
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Failed to handle get email info", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		Email
// @Router		/mail [get]
// @Summary		Get Email histories
// @Description	Get Email histories
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
func (c *controller) EmailSentCollection(ctx *gin.Context) {
	results, err := c.service.GetCollection(ctx)
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Failed to handle get email sent collection", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", results)
}
