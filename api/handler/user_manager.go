package handler

import (
	"crm_api_gateway/genproto/manager_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/manager/getall [GET]
// @Summary Get all manageres
// @Description API for getting all manageres
// @Tags manager
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllManager(c *gin.Context) {
	manager := &manager_service.GetListManagerRequest{}

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

	manager.Search = search
	manager.Offset = int64(page)
	manager.Limit = int64(limit)

	resp, err := h.grpcClient.ManagerService().GetList(c.Request.Context(), manager)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/manager/create [POST]
// @Summary Create manager
// @Description API for creating manageres
// @Tags manager
// @Accept  json
// @Produce  json
// @Param		manager body  manager_service.CreateManager true "manager"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateManager(c *gin.Context) {
	manager := &manager_service.CreateManager{}
	if err := c.ShouldBindJSON(&manager); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ManagerService().Create(c.Request.Context(), manager)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/manager/update [PUT]
// @Summary Update manager
// @Description API for Updating manageres
// @Tags manager
// @Accept  json
// @Produce  json
// @Param		manager body  manager_service.UpdateManager true "manager"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateManager(c *gin.Context) {
	manager := &manager_service.UpdateManager{}
	if err := c.ShouldBindJSON(&manager); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ManagerService().Update(c.Request.Context(), manager)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/manager/get/{id} [GET]
// @Summary Get manager
// @Description API for getting manager
// @Tags manager
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetManagerById(c *gin.Context) {
	id := c.Param("id")
	manager := &manager_service.ManagerPrimaryKey{Id: id}

	resp, err := h.grpcClient.ManagerService().GetByID(c.Request.Context(), manager)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/manager/delete/{id} [DELETE]
// @Summary Delete manager
// @Description API for deleting manager
// @Tags manager
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteManager(c *gin.Context) {
	id := c.Param("id")
	manager := &manager_service.ManagerPrimaryKey{Id: id}

	resp, err := h.grpcClient.ManagerService().Delete(c.Request.Context(), manager)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}
