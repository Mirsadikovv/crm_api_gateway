package handler

import (
	"crm_api_gateway/genproto/administrator_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"fmt"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change administrator")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		administrator := &administrator_service.CreateAdministrator{}
		if err := c.ShouldBindJSON(&administrator); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(administrator.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(administrator.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(administrator.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.AdministratorService().Create(c.Request.Context(), administrator)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating administrator")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change administrator")
	}
}

// @Security ApiKeyAuth
// @Router /v1/administrator/update [PUT]
// @Summary Update administrator
// @Description API for Updating administrators
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param		administrator body  administrator_service.UpdateAdministrator true "administrator"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateAdministrator(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		administrator := &administrator_service.UpdateAdministrator{}
		if err := c.ShouldBindJSON(&administrator); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(administrator.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(administrator.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.AdministratorService().Update(c.Request.Context(), administrator)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating administrator")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change administrator")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		administrator := &administrator_service.AdministratorPrimaryKey{Id: id}

		resp, err := h.grpcClient.AdministratorService().GetByID(c.Request.Context(), administrator)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting administrator")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrator and managers can change administrator")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		administrator := &administrator_service.AdministratorPrimaryKey{Id: id}

		resp, err := h.grpcClient.AdministratorService().Delete(c.Request.Context(), administrator)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting administrator")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrator and managers can change administrator")
	}
}

// AdministratorLogin godoc
// @Router       /v1/administrator/login [POST]
// @Summary      Administrator login
// @Description  Administrator login
// @Tags         administrator
// @Accept       json
// @Produce      json
// @Param        login body administrator_service.AdministratorLoginRequest true "login"
// @Success      201  {object}  administrator_service.AdministratorLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) AdministratorLogin(c *gin.Context) {
	loginReq := &administrator_service.AdministratorLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.AdministratorService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// AdministratorRegister godoc
// @Router       /v1/administrator/register [POST]
// @Summary      Administrator register
// @Description  Administrator register
// @Tags         administrator
// @Accept       json
// @Produce      json
// @Param        register body administrator_service.AdministratorRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) AdministratorRegister(c *gin.Context) {
	loginReq := &administrator_service.AdministratorRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.AdministratorService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr administrator")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// AdministratorRegister godoc
// @Router       /v1/administrator/register-confirm [POST]
// @Summary      Administrator register
// @Description  Administrator register
// @Tags         administrator
// @Accept       json
// @Produce      json
// @Param        register body administrator_service.AdministratorRegisterConfRequest true "register"
// @Success      201  {object}  administrator_service.AdministratorLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) AdministratorRegisterConfirm(c *gin.Context) {
	req := &administrator_service.AdministratorRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.AdministratorService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/administrator/change_password [PATCH]
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
func (h *handler) AdministratorChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		administrator := &administrator_service.AdministratorChangePassword{}
		if err := c.ShouldBindJSON(&administrator); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(administrator.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.AdministratorService().ChangePassword(c.Request.Context(), administrator)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing administrator's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, administrator and managers can change administrator")
	}
}
