package router

import (
	"filesProcessor/action/fileProcess"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.New()
	api := r.Group("api/files-processor-api")

	api.GET(
		"/health", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK, gin.H{
					"status": http.StatusOK,
				},
			)
		},
	)
	SetFilesProcessor(api)
	return r
}

func SetFilesProcessor(api *gin.RouterGroup) {
	repository := fileProcess.NewRepository("jdcv0116@gmail.com", "A$i7kr.4m4NsiB9", "smtp.gmail.com", "587")
	services := fileProcess.NewServer(repository)
	fileProcessHandler := fileProcess.ProcessHandler(services)
	api.POST("/action/file/process", fileProcessHandler)
}
