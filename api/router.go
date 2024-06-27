package api

import (
	_ "crm_api_gateway/api/docs" //for swagger
	"crm_api_gateway/api/handler"
	"crm_api_gateway/config"
	"crm_api_gateway/pkg/grpc_client"
	"crm_api_gateway/pkg/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// New ...
// @title           Swagger CRM system API
// @version         1.0
// @description     This is a CRM celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	// r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	///////////////////////// USER_service
	r.GET("/v1/branch/getall", handler.GetAllBranch)
	r.GET("/v1/branch/get/:id", handler.GetBranchById)
	r.POST("/v1/branch/create", handler.CreateBranch)
	r.PUT("/v1/branch/update", handler.UpdateBranch)
	r.DELETE("/v1/branch/delete/:id", handler.DeleteBranch)

	r.GET("/v1/administrator/getall", handler.GetAllAdministrator)
	r.GET("/v1/administrator/get/:id", handler.GetAdministratorById)
	r.POST("/v1/administrator/create", handler.CreateAdministrator)
	r.PUT("/v1/administrator/update", handler.UpdateAdministrator)
	r.DELETE("/v1/administrator/delete/:id", handler.DeleteAdministrator)

	r.GET("/v1/event_registrate/get/:id", handler.GetEventRegistrateById)
	r.POST("/v1/event_registrate/create", handler.CreateEventRegistrate)
	r.PUT("/v1/event_registrate/update", handler.UpdateEventRegistrate)
	r.DELETE("/v1/event_registrate/delete/:id", handler.DeleteEventRegistrate)

	r.GET("/v1/event/getall", handler.GetAllEvent)
	r.GET("/v1/event/get/:id", handler.GetEventById)
	r.POST("/v1/event/create", handler.CreateEvent)
	r.PUT("/v1/event/update", handler.UpdateEvent)
	r.DELETE("/v1/event/delete/:id", handler.DeleteEvent)

	r.GET("/v1/group/getall", handler.GetAllGroup)
	r.GET("/v1/group/get/:id", handler.GetGroupById)
	r.POST("/v1/group/create", handler.CreateGroup)
	r.PUT("/v1/group/update", handler.UpdateGroup)
	r.DELETE("/v1/group/delete/:id", handler.DeleteGroup)

	r.GET("/v1/manager/getall", handler.GetAllManager)
	r.GET("/v1/manager/get/:id", handler.GetManagerById)
	r.POST("/v1/manager/create", handler.CreateManager)
	r.PUT("/v1/manager/update", handler.UpdateManager)
	r.DELETE("/v1/manager/delete/:id", handler.DeleteManager)

	r.GET("/v1/student/getall", handler.GetAllStudent)
	r.GET("/v1/student/get/:id", handler.GetStudentById)
	r.POST("/v1/student/create", handler.CreateStudent)
	r.PUT("/v1/student/update", handler.UpdateStudent)
	r.DELETE("/v1/student/delete/:id", handler.DeleteStudent)

	r.GET("/v1/superadmin/get/:id", handler.GetSuperadminById)
	r.POST("/v1/superadmin/create", handler.CreateSuperadmin)
	r.PUT("/v1/superadmin/update", handler.UpdateSuperadmin)
	r.DELETE("/v1/superadmin/delete/:id", handler.DeleteSuperadmin)

	r.GET("/v1/support_teacher/getall", handler.GetAllSupportTeacher)
	r.GET("/v1/support_teacher/get/:id", handler.GetSupportTeacherById)
	r.POST("/v1/support_teacher/create", handler.CreateSupportTeacher)
	r.PUT("/v1/support_teacher/update", handler.UpdateSupportTeacher)
	r.DELETE("/v1/support_teacher/delete/:id", handler.DeleteSupportTeacher)

	r.GET("/v1/teacher/getall", handler.GetAllTeacher)
	r.GET("/v1/teacher/get/:id", handler.GetTeacherById)
	r.POST("/v1/teacher/create", handler.CreateTeacher)
	r.PUT("/v1/teacher/update", handler.UpdateTeacher)
	r.DELETE("/v1/teacher/delete/:id", handler.DeleteTeacher)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
// 	}
// 	c.Next()
// }
