package handler

import (
	"crm_api_gateway/genproto/teacher_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/teacher/getall [GET]
// @Summary Get all teacheres
// @Description API for getting all teacheres
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllTeacher(c *gin.Context) {
	teacher := &teacher_service.GetListTeacherRequest{}

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

	teacher.Search = search
	teacher.Offset = int64(page)
	teacher.Limit = int64(limit)

	resp, err := h.grpcClient.TeacherService().GetList(c.Request.Context(), teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/teacher/create [POST]
// @Summary Create teacher
// @Description API for creating teacheres
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param		teacher body  teacher_service.CreateTeacher true "teacher"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateTeacher(c *gin.Context) {
	teacher := &teacher_service.CreateTeacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.TeacherService().Create(c.Request.Context(), teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/teacher/update [PUT]
// @Summary Update teacher
// @Description API for Updating teacheres
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param		teacher body  teacher_service.UpdateTeacher true "teacher"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateTeacher(c *gin.Context) {
	teacher := &teacher_service.UpdateTeacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.TeacherService().Update(c.Request.Context(), teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/teacher/get/{id} [GET]
// @Summary Get teacher
// @Description API for getting teacher
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetTeacherById(c *gin.Context) {
	id := c.Param("id")
	teacher := &teacher_service.TeacherPrimaryKey{Id: id}

	resp, err := h.grpcClient.TeacherService().GetByID(c.Request.Context(), teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/teacher/delete/{id} [DELETE]
// @Summary Delete teacher
// @Description API for deleting teacher
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	teacher := &teacher_service.TeacherPrimaryKey{Id: id}

	resp, err := h.grpcClient.TeacherService().Delete(c.Request.Context(), teacher)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}
