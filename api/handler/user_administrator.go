package handler

import (
	"crm_api_gateway/genproto/administrator_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/administrator/getall [GET]
// @Summary Get all administratores
// @Description API for getting all administratores
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllAdministrator(c *gin.Context) {
	administrator := &administrator_service.GetListAdministratorRequest{}

	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
		return
	}

	administrator.Search = search
	administrator.Offset = int64(page)
	administrator.Limit = int64(limit)

	resp, err := h.grpcClient.AdministratorService().GetList(c.Request.Context(), administrator)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/administrator/create [POST]
// @Summary Create administrator
// @Description API for creating administratores
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param		administrator body  administrator_service.CreateAdministrator true "administrator"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateAdministrator(c *gin.Context) {
	administrator := &administrator_service.CreateAdministrator{}
	if err := c.ShouldBindJSON(&administrator); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.AdministratorService().Create(c.Request.Context(), administrator)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/administrator/update [PUT]
// @Summary Update administrator
// @Description API for Updating administratores
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param		administrator body  administrator_service.UpdateAdministrator true "administrator"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateAdministrator(c *gin.Context) {
	administrator := &administrator_service.UpdateAdministrator{}
	if err := c.ShouldBindJSON(&administrator); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.AdministratorService().Update(c.Request.Context(), administrator)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/administrator/get/{id} [GET]
// @Summary Get administrator
// @Description API for getting administrator
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAdministratorById(c *gin.Context) {
	id := c.Param("id")
	administrator := &administrator_service.AdministratorPrimaryKey{Id: id}

	resp, err := h.grpcClient.AdministratorService().GetByID(c.Request.Context(), administrator)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/administrator/delete/{id} [DELETE]
// @Summary Delete administrator
// @Description API for deleting administrator
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteAdministrator(c *gin.Context) {
	id := c.Param("id")
	administrator := &administrator_service.AdministratorPrimaryKey{Id: id}

	resp, err := h.grpcClient.AdministratorService().Delete(c.Request.Context(), administrator)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}
