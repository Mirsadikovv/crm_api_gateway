package handler

import (
	"crm_api_gateway/genproto/superadmin_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"fmt"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		superadmin := &superadmin_service.CreateSuperadmin{}
		if err := c.ShouldBindJSON(&superadmin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(superadmin.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(superadmin.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(superadmin.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.SuperadminService().Create(c.Request.Context(), superadmin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating superadmin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can create superadmin")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		superadmin := &superadmin_service.UpdateSuperadmin{}
		if err := c.ShouldBindJSON(&superadmin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(superadmin.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(superadmin.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.SuperadminService().Update(c.Request.Context(), superadmin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating superadmin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can update superadmin")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		id := c.Param("id")
		superadmin := &superadmin_service.SuperadminPrimaryKey{Id: id}

		resp, err := h.grpcClient.SuperadminService().GetByID(c.Request.Context(), superadmin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting superadmin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can get superadmin")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		id := c.Param("id")
		superadmin := &superadmin_service.SuperadminPrimaryKey{Id: id}

		resp, err := h.grpcClient.SuperadminService().Delete(c.Request.Context(), superadmin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting superadmin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can delete superadmin")
	}
}

// SuperadminLogin godoc
// @Router       /v1/superadmin/login [POST]
// @Summary      Superadmin login
// @Description  Superadmin login
// @Tags         superadmin
// @Accept       json
// @Produce      json
// @Param        login body superadmin_service.SuperadminLoginRequest true "login"
// @Success      201  {object}  superadmin_service.SuperadminLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SuperadminLogin(c *gin.Context) {
	loginReq := &superadmin_service.SuperadminLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.SuperadminService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// SuperadminRegister godoc
// @Router       /v1/superadmin/register [POST]
// @Summary      Superadmin register
// @Description  Superadmin register
// @Tags         superadmin
// @Accept       json
// @Produce      json
// @Param        register body superadmin_service.SuperadminRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SuperadminRegister(c *gin.Context) {
	loginReq := &superadmin_service.SuperadminRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.SuperadminService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr superadmin")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// SuperadminRegister godoc
// @Router       /v1/superadmin/register-confirm [POST]
// @Summary      Superadmin register
// @Description  Superadmin register
// @Tags         superadmin
// @Accept       json
// @Produce      json
// @Param        register body superadmin_service.SuperadminRegisterConfRequest true "register"
// @Success      201  {object}  superadmin_service.SuperadminLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) SuperadminRegisterConfirm(c *gin.Context) {
	req := &superadmin_service.SuperadminRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.SuperadminService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/superadmin/change_password [PATCH]
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
func (h *handler) SuperadminChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		superadmin := &superadmin_service.SuperadminChangePassword{}
		if err := c.ShouldBindJSON(&superadmin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(superadmin.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.SuperadminService().ChangePassword(c.Request.Context(), superadmin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing superadmin's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can update superadmin's password")
	}
}
