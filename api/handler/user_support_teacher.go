package handler

import (
	"crm_api_gateway/genproto/support_teacher_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"fmt"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can update support_teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		support_teacher := &support_teacher_service.CreateSupportTeacher{}
		if err := c.ShouldBindJSON(&support_teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(support_teacher.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(support_teacher.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(support_teacher.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.SupportTeacherService().Create(c.Request.Context(), support_teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating support_teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and teacher can crate support_teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "support_teacher" {
		support_teacher := &support_teacher_service.UpdateSupportTeacher{}
		if err := c.ShouldBindJSON(&support_teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(support_teacher.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(support_teacher.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.SupportTeacherService().Update(c.Request.Context(), support_teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating support_teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and support_teacher can update support_teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "support_teacher" {
		id := c.Param("id")
		support_teacher := &support_teacher_service.SupportTeacherPrimaryKey{Id: id}

		resp, err := h.grpcClient.SupportTeacherService().GetByID(c.Request.Context(), support_teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting support_teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and support_teacher can update support_teacher")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "support_teacher" {

		id := c.Param("id")
		support_teacher := &support_teacher_service.SupportTeacherPrimaryKey{Id: id}

		resp, err := h.grpcClient.SupportTeacherService().Delete(c.Request.Context(), support_teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting support_teacher")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and support_teacher can delete support_teacher")
	}
}

// SupportTeacherLogin godoc
// @Router       /v1/support_teacher/login [POST]
// @Summary      SupportTeacher login
// @Description  SupportTeacher login
// @Tags         support_teacher
// @Accept       json
// @Produce      json
// @Param        login body support_teacher_service.SupportTeacherLoginRequest true "login"
// @Success      201  {object}  support_teacher_service.SupportTeacherLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SupportTeacherLogin(c *gin.Context) {
	loginReq := &support_teacher_service.SupportTeacherLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.SupportTeacherService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// SupportTeacherRegister godoc
// @Router       /v1/support_teacher/register [POST]
// @Summary      SupportTeacher register
// @Description  SupportTeacher register
// @Tags         support_teacher
// @Accept       json
// @Produce      json
// @Param        register body support_teacher_service.SupportTeacherRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SupportTeacherRegister(c *gin.Context) {
	loginReq := &support_teacher_service.SupportTeacherRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.SupportTeacherService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr support_teacher")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// SupportTeacherRegister godoc
// @Router       /v1/support_teacher/register-confirm [POST]
// @Summary      SupportTeacher register
// @Description  SupportTeacher register
// @Tags         support_teacher
// @Accept       json
// @Produce      json
// @Param        register body support_teacher_service.SupportTeacherRegisterConfRequest true "register"
// @Success      201  {object}  support_teacher_service.SupportTeacherLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SupportTeacherRegisterConfirm(c *gin.Context) {
	req := &support_teacher_service.SupportTeacherRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.SupportTeacherService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/support_teacher/change_password [PATCH]
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
func (h *handler) SupportTeacherChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" || authInfo.UserRole == "support_teacher" {
		support_teacher := &support_teacher_service.SupportTeacherChangePassword{}
		if err := c.ShouldBindJSON(&support_teacher); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		err := validator.ValidatePassword(support_teacher.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.SupportTeacherService().ChangePassword(c.Request.Context(), support_teacher)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing support_teacher's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and support_teacher can update password support_teacher")
	}
}
