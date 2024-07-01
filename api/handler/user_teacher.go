package handler

import (
	"crm_api_gateway/genproto/teacher_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"fmt"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers  can get teachers")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		teacher := &teacher_service.CreateTeacher{}
		if err := c.ShouldBindJSON(&teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(teacher.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(teacher.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(teacher.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.TeacherService().Create(c.Request.Context(), teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers can create teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "manager" {

		teacher := &teacher_service.UpdateTeacher{}
		if err := c.ShouldBindJSON(&teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(teacher.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(teacher.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.TeacherService().Update(c.Request.Context(), teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can update teacher")
	}

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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		teacher := &teacher_service.TeacherPrimaryKey{Id: id}

		resp, err := h.grpcClient.TeacherService().GetByID(c.Request.Context(), teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can get teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		teacher := &teacher_service.TeacherPrimaryKey{Id: id}

		resp, err := h.grpcClient.TeacherService().Delete(c.Request.Context(), teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can update teacher")
	}

}

// TeacherLogin godoc
// @Router       /v1/teacher/login [POST]
// @Summary      Teacher login
// @Description  Teacher login
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        login body teacher_service.TeacherLoginRequest true "login"
// @Success      201  {object}  teacher_service.TeacherLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) TeacherLogin(c *gin.Context) {
	loginReq := &teacher_service.TeacherLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.TeacherService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// TeacherRegister godoc
// @Router       /v1/teacher/register [POST]
// @Summary      Teacher register
// @Description  Teacher register
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        register body teacher_service.TeacherRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) TeacherRegister(c *gin.Context) {
	loginReq := &teacher_service.TeacherRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.TeacherService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr teacher")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// TeacherRegister godoc
// @Router       /v1/teacher/register-confirm [POST]
// @Summary      Teacher register
// @Description  Teacher register
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        register body teacher_service.TeacherRegisterConfRequest true "register"
// @Success      201  {object}  teacher_service.TeacherLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) TeacherRegisterConfirm(c *gin.Context) {
	req := &teacher_service.TeacherRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.TeacherService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/teacher/change_password [PATCH]
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
func (h *handler) TeacherChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "teacher" || authInfo.UserRole == "manager" {

		teacher := &teacher_service.TeacherChangePassword{}
		if err := c.ShouldBindJSON(&teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		err := validator.ValidatePassword(teacher.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.TeacherService().ChangePassword(c.Request.Context(), teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing teacher's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can update password")
	}
}
