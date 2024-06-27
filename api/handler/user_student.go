package handler

import (
	"crm_api_gateway/genproto/student_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/student/getall [GET]
// @Summary Get all studentes
// @Description API for getting all studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllStudent(c *gin.Context) {
	student := &student_service.GetListStudentRequest{}

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

	student.Search = search
	student.Offset = int64(page)
	student.Limit = int64(limit)

	resp, err := h.grpcClient.StudentService().GetList(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating student")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/student/create [POST]
// @Summary Create student
// @Description API for creating studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		student body  student_service.CreateStudent true "student"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateStudent(c *gin.Context) {
	student := &student_service.CreateStudent{}
	if err := c.ShouldBindJSON(&student); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.StudentService().Create(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating student")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/student/update [PUT]
// @Summary Update student
// @Description API for Updating studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		student body  student_service.UpdateStudent true "student"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateStudent(c *gin.Context) {
	student := &student_service.UpdateStudent{}
	if err := c.ShouldBindJSON(&student); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.StudentService().Update(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating student")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/student/get/{id} [GET]
// @Summary Get student
// @Description API for getting student
// @Tags student
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetStudentById(c *gin.Context) {
	id := c.Param("id")
	student := &student_service.StudentPrimaryKey{Id: id}

	resp, err := h.grpcClient.StudentService().GetByID(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting student")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/student/delete/{id} [DELETE]
// @Summary Delete student
// @Description API for deleting student
// @Tags student
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	student := &student_service.StudentPrimaryKey{Id: id}

	resp, err := h.grpcClient.StudentService().Delete(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting student")
		return
	}
	c.JSON(http.StatusOK, resp)
}
