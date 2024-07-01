package handler

import (
	"crm_api_gateway/genproto/manager_service"
	"crm_api_gateway/pkg/validator"
	"errors"
	"fmt"
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

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
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins  can get managers")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" {

		manager := &manager_service.CreateManager{}
		if err := c.ShouldBindJSON(&manager); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(manager.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(manager.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(manager.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.ManagerService().Create(c.Request.Context(), manager)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating manager")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change manager")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		manager := &manager_service.UpdateManager{}
		if err := c.ShouldBindJSON(&manager); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(manager.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(manager.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.ManagerService().Update(c.Request.Context(), manager)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating manager")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change manager")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		manager := &manager_service.ManagerPrimaryKey{Id: id}

		resp, err := h.grpcClient.ManagerService().GetByID(c.Request.Context(), manager)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting manager")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change manager")
	}
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
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		id := c.Param("id")
		manager := &manager_service.ManagerPrimaryKey{Id: id}

		resp, err := h.grpcClient.ManagerService().Delete(c.Request.Context(), manager)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting manager")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change manager")
	}
}

// ManagerLogin godoc
// @Router       /v1/manager/login [POST]
// @Summary      Manager login
// @Description  Manager login
// @Tags         manager
// @Accept       json
// @Produce      json
// @Param        login body manager_service.ManagerLoginRequest true "login"
// @Success      201  {object}  manager_service.ManagerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) ManagerLogin(c *gin.Context) {
	loginReq := &manager_service.ManagerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.ManagerService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// ManagerRegister godoc
// @Router       /v1/manager/register [POST]
// @Summary      Manager register
// @Description  Manager register
// @Tags         manager
// @Accept       json
// @Produce      json
// @Param        register body manager_service.ManagerRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) ManagerRegister(c *gin.Context) {
	loginReq := &manager_service.ManagerRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.ManagerService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr manager")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// ManagerRegister godoc
// @Router       /v1/manager/register-confirm [POST]
// @Summary      Manager register
// @Description  Manager register
// @Tags         manager
// @Accept       json
// @Produce      json
// @Param        register body manager_service.ManagerRegisterConfRequest true "register"
// @Success      201  {object}  manager_service.ManagerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) ManagerRegisterConfirm(c *gin.Context) {
	req := &manager_service.ManagerRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.ManagerService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/manager/change_password [PATCH]
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
func (h *handler) ManagerChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "manager" {

		manager := &manager_service.ManagerChangePassword{}
		if err := c.ShouldBindJSON(&manager); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(manager.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.ManagerService().ChangePassword(c.Request.Context(), manager)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing manager's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and managers can change manager")
	}
}
