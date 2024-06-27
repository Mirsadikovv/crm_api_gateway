package handler

import (
	"crm_api_gateway/genproto/support_teacher_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/support_teacher/getall [GET]
// @Summary Get all support_teacheres
// @Description API for getting all support_teacheres
// @Tags support_teacher
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllSupportTeacher(c *gin.Context) {
	support_teacher := &support_teacher_service.GetListSupportTeacherRequest{}

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

	support_teacher.Search = search
	support_teacher.Offset = int64(page)
	support_teacher.Limit = int64(limit)

	resp, err := h.grpcClient.SupportTeacherService().GetList(c.Request.Context(), support_teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating support_teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/support_teacher/create [POST]
// @Summary Create support_teacher
// @Description API for creating support_teacheres
// @Tags support_teacher
// @Accept  json
// @Produce  json
// @Param		support_teacher body  support_teacher_service.CreateSupportTeacher true "support_teacher"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSupportTeacher(c *gin.Context) {
	support_teacher := &support_teacher_service.CreateSupportTeacher{}
	if err := c.ShouldBindJSON(&support_teacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Create(c.Request.Context(), support_teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating support_teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/support_teacher/update [PUT]
// @Summary Update support_teacher
// @Description API for Updating support_teacheres
// @Tags support_teacher
// @Accept  json
// @Produce  json
// @Param		support_teacher body  support_teacher_service.UpdateSupportTeacher true "support_teacher"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSupportTeacher(c *gin.Context) {
	support_teacher := &support_teacher_service.UpdateSupportTeacher{}
	if err := c.ShouldBindJSON(&support_teacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Update(c.Request.Context(), support_teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating support_teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/support_teacher/get/{id} [GET]
// @Summary Get support_teacher
// @Description API for getting support_teacher
// @Tags support_teacher
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetSupportTeacherById(c *gin.Context) {
	id := c.Param("id")
	support_teacher := &support_teacher_service.SupportTeacherPrimaryKey{Id: id}

	resp, err := h.grpcClient.SupportTeacherService().GetByID(c.Request.Context(), support_teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting support_teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/support_teacher/delete/{id} [DELETE]
// @Summary Delete support_teacher
// @Description API for deleting support_teacher
// @Tags support_teacher
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSupportTeacher(c *gin.Context) {
	id := c.Param("id")
	support_teacher := &support_teacher_service.SupportTeacherPrimaryKey{Id: id}

	resp, err := h.grpcClient.SupportTeacherService().Delete(c.Request.Context(), support_teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting support_teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}
