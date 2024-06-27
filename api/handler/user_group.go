package handler

import (
	"crm_api_gateway/genproto/group_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/group/getall [GET]
// @Summary Get all groupes
// @Description API for getting all groupes
// @Tags group
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllGroup(c *gin.Context) {
	group := &group_service.GetListGroupRequest{}

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

	group.Search = search
	group.Offset = int64(page)
	group.Limit = int64(limit)

	resp, err := h.grpcClient.GroupService().GetList(c.Request.Context(), group)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating group")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/group/create [POST]
// @Summary Create group
// @Description API for creating groupes
// @Tags group
// @Accept  json
// @Produce  json
// @Param		group body  group_service.CreateGroup true "group"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateGroup(c *gin.Context) {
	group := &group_service.CreateGroup{}
	if err := c.ShouldBindJSON(&group); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.GroupService().Create(c.Request.Context(), group)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating group")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/group/update [PUT]
// @Summary Update group
// @Description API for Updating groupes
// @Tags group
// @Accept  json
// @Produce  json
// @Param		group body  group_service.UpdateGroup true "group"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateGroup(c *gin.Context) {
	group := &group_service.UpdateGroup{}
	if err := c.ShouldBindJSON(&group); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.GroupService().Update(c.Request.Context(), group)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating group")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/group/get/{id} [GET]
// @Summary Get group
// @Description API for getting group
// @Tags group
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetGroupById(c *gin.Context) {
	id := c.Param("id")
	group := &group_service.GroupPrimaryKey{Id: id}

	resp, err := h.grpcClient.GroupService().GetByID(c.Request.Context(), group)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting group")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/group/delete/{id} [DELETE]
// @Summary Delete group
// @Description API for deleting group
// @Tags group
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	group := &group_service.GroupPrimaryKey{Id: id}

	resp, err := h.grpcClient.GroupService().Delete(c.Request.Context(), group)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting group")
		return
	}
	c.JSON(http.StatusOK, resp)
}
