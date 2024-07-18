package api

import (
	"novanxyz/api/controller"
	"novanxyz/utils"

	"net/http"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(taskController *controller.TaskController) *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20 // 10 MiB
	router.Use(nice.Recovery(utils.ErrorResponseRecovery))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "welcome home") })
	baseRouter := router.Group("/api")
	taskRouter := baseRouter.Group("/tasks")
	taskRouter.GET("", taskController.FindAll)
	taskRouter.GET("/:taskId", taskController.FindById)
	taskRouter.POST("", taskController.Create)
	taskRouter.PUT("/:taskId", taskController.Update)
	taskRouter.PATCH("/:taskId/:status", taskController.Mark)
	taskRouter.DELETE("/:taskId", taskController.Delete)

	taskRouter.POST("/:taskId/files", taskController.UploadTaskFile)
	taskRouter.GET("/:taskId/files", taskController.GetTaskFiles)
	taskRouter.GET("/:taskId/files/:fileId", taskController.DownloadTaskFile)
	taskRouter.DELETE("/:taskId/files/:fileId", taskController.DeleteTaskFile)

	return router
}
