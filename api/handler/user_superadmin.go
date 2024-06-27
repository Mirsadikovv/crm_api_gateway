package handler

import (
	"crm_api_gateway/genproto/superadmin_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/superadmin/create [POST]
// @Summary Create superadmin
// @Description API for creating superadmines
// @Tags superadmin
// @Accept  json
// @Produce  json
// @Param		superadmin body  superadmin_service.CreateSuperadmin true "superadmin"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSuperadmin(c *gin.Context) {
	superadmin := &superadmin_service.CreateSuperadmin{}
	if err := c.ShouldBindJSON(&superadmin); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().Create(c.Request.Context(), superadmin)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating superadmin")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/superadmin/update [PUT]
// @Summary Update superadmin
// @Description API for Updating superadmines
// @Tags superadmin
// @Accept  json
// @Produce  json
// @Param		superadmin body  superadmin_service.UpdateSuperadmin true "superadmin"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSuperadmin(c *gin.Context) {
	superadmin := &superadmin_service.UpdateSuperadmin{}
	if err := c.ShouldBindJSON(&superadmin); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().Update(c.Request.Context(), superadmin)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating superadmin")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/superadmin/get/{id} [GET]
// @Summary Get superadmin
// @Description API for getting superadmin
// @Tags superadmin
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetSuperadminById(c *gin.Context) {
	id := c.Param("id")
	superadmin := &superadmin_service.SuperadminPrimaryKey{Id: id}

	resp, err := h.grpcClient.SuperadminService().GetByID(c.Request.Context(), superadmin)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting superadmin")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/superadmin/delete/{id} [DELETE]
// @Summary Delete superadmin
// @Description API for deleting superadmin
// @Tags superadmin
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSuperadmin(c *gin.Context) {
	id := c.Param("id")
	superadmin := &superadmin_service.SuperadminPrimaryKey{Id: id}

	resp, err := h.grpcClient.SuperadminService().Delete(c.Request.Context(), superadmin)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting superadmin")
		return
	}
	c.JSON(http.StatusOK, resp)
}
