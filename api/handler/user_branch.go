package handler

import (
	"crm_api_gateway/genproto/branch_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/branch/getall [GET]
// @Summary Get all branches
// @Description API for getting all branches
// @Tags branch
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllBranch(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		branch := &branch_service.GetListBranchRequest{}

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

		branch.Search = search
		branch.Offset = int64(page)
		branch.Limit = int64(limit)

		resp, err := h.grpcClient.BranchService().GetList(c.Request.Context(), branch)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating branch")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change branch")
	}
}

// @Security ApiKeyAuth
// @Router /v1/branch/create [POST]
// @Summary Create branch
// @Description API for creating branches
// @Tags branch
// @Accept  json
// @Produce  json
// @Param		branch body  branch_service.CreateBranch true "branch"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateBranch(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		branch := &branch_service.CreateBranch{}
		if err := c.ShouldBindJSON(&branch); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.BranchService().Create(c.Request.Context(), branch)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating branch")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change branch")
	}
}

// @Security ApiKeyAuth
// @Router /v1/branch/update [PUT]
// @Summary Update branch
// @Description API for Updating branches
// @Tags branch
// @Accept  json
// @Produce  json
// @Param		branch body  branch_service.UpdateBranch true "branch"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateBranch(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		branch := &branch_service.UpdateBranch{}
		if err := c.ShouldBindJSON(&branch); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.BranchService().Update(c.Request.Context(), branch)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating branch")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change branch")
	}
}

// @Security ApiKeyAuth
// @Router /v1/branch/get/{id} [GET]
// @Summary Get branch
// @Description API for getting branch
// @Tags branch
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetBranchById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		id := c.Param("id")
		branch := &branch_service.BranchPrimaryKey{Id: id}

		resp, err := h.grpcClient.BranchService().GetByID(c.Request.Context(), branch)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting branch")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change branch")
	}
}

// @Security ApiKeyAuth
// @Router /v1/branch/delete/{id} [DELETE]
// @Summary Delete branch
// @Description API for deleting branch
// @Tags branch
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteBranch(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		id := c.Param("id")
		branch := &branch_service.BranchPrimaryKey{Id: id}

		resp, err := h.grpcClient.BranchService().Delete(c.Request.Context(), branch)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting branch")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change branch")
	}

}
