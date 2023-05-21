package controllers

import (
	"errors"

	"github.com/abelz123456/celestial-api/api/file/domain"
	"github.com/abelz123456/celestial-api/api/file/services"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/utils/api/response"
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

// @Tags		File
// @Router		/file [post]
// @Summary		Upload new File
// @Description	Upload new File
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param 		content formData file true "File to upload"
// @Produce 	application/json
func (c *controller) Upload(ctx *gin.Context) {
	userOid := ctx.GetString("oid")
	file, err := ctx.FormFile("content")
	if err != nil || file == nil {
		if file == nil {
			err = errors.New("form 'content' is required")
		}

		response.SendJson(ctx, response.ErrFailedUploadFile, "", map[string]string{"error": err.Error()})
		return
	}

	data := domain.UploadFileData{UserOid: userOid, File: *file}
	result, err := c.service.UploadFile(ctx.Request.Context(), data)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedUploadFile, "", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Created, "New file uploaded", result)
}

// @Tags		File
// @Router		/file [get]
// @Summary		Get File Collection
// @Description	Get File Collection
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Produce 	application/json
func (c *controller) GetCollection(ctx *gin.Context) {
	results, err := c.service.GetCollection(ctx.Request.Context())
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Error handle get file collection", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", results)
}

// @Tags		File
// @Router		/file/{uid} [get]
// @Summary		Get File Info
// @Description	Get File Info
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param		uid path string true "uid of File"
// @Produce 	application/json
func (c *controller) GetInfo(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	result, err := c.service.GetInfo(ctx.Request.Context(), uid)
	if err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "", map[string]string{"error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "", result)
}

// @Tags		File
// @Router		/file/{uid} [put]
// @Summary		Change File
// @Description	Change File
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param		uid path string true "uid of File"
// @Param 		content formData file true "File to upload"
// @Produce 	application/json
func (c *controller) Replace(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	file, err := ctx.FormFile("content")
	if err != nil || file == nil {
		if file == nil {
			err = errors.New("form 'content' is required")
		}

		response.SendJson(ctx, response.ErrFailedUploadFile, "", map[string]string{"error": err.Error()})
		return
	}

	fileData := domain.ReplaceFileData{FileUID: uid, File: *file}
	result, err := c.service.ReplaceFile(ctx.Request.Context(), fileData)
	if err != nil {
		response.SendJson(ctx, response.ErrFailedUploadFile, "", map[string]string{"uid": uid, "error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "File updated", result)
}

// @Tags		File
// @Router		/file/{uid} [delete]
// @Summary		Delete File
// @Description	Delete File
// @Param 		Authorization header string true "With value 'Bearer {authToken}'"
// @Param		uid path string true "uid of File"
// @Produce 	application/json
func (c *controller) Unlink(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	if err := c.service.UnlinkFile(ctx.Request.Context(), uid); err != nil {
		response.SendJson(ctx, response.ErrBadRequest, "Error handle delete file", map[string]string{"uid": uid, "error": err.Error()})
		return
	}

	response.SendJson(ctx, response.Ok, "File deleted", map[string]string{"uid": uid})
}
