package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupResultadoEvaluacionRoutes(router *gin.Engine, service *services.ResultadoEvaluacionService) {
	resultadoController := controllers.NewResultadoEvaluacionController(service)

	resultadoEvaluacion := router.Group("/resultado-evaluacion")
	resultadoEvaluacion.Use(middleware.AuthMiddleware())
	{
		resultadoEvaluacion.GET("", resultadoController.GetAllResultados)
		resultadoEvaluacion.GET("/:id", resultadoController.GetResultadoByID)
		resultadoEvaluacion.GET("/alumno/:alumnoId", resultadoController.GetResultadosByAlumnoID)
		resultadoEvaluacion.GET("/evaluacion/:evaluacionId", resultadoController.GetResultadosByEvaluacionID)
		resultadoEvaluacion.POST("", resultadoController.CreateResultado)
		resultadoEvaluacion.PUT("/:id", resultadoController.UpdateResultado)
		resultadoEvaluacion.DELETE("/:id", resultadoController.DeleteResultado)
	}
}
