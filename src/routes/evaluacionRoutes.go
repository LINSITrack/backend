package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupEvaluacionRoutes(router *gin.Engine, service *services.EvaluacionService) {
	evaluacionController := controllers.NewEvaluacionController(service)

	adminOnlyEvaluaciones := router.Group("/evaluaciones")
	adminOnlyEvaluaciones.Use(middleware.AuthMiddleware())
	adminOnlyEvaluaciones.Use(middleware.RequireRole(models.RoleAdmin))
	{
		adminOnlyEvaluaciones.GET("/", evaluacionController.GetAllEvaluaciones)
		adminOnlyEvaluaciones.GET("/:id", evaluacionController.GetEvaluacionByID)
		adminOnlyEvaluaciones.GET("/comision/:comisionId", evaluacionController.GetEvaluacionesByComisionID)
		adminOnlyEvaluaciones.POST("/", evaluacionController.CreateEvaluacion)
		adminOnlyEvaluaciones.PATCH("/:id", evaluacionController.UpdateEvaluacion)
		adminOnlyEvaluaciones.DELETE("/:id", evaluacionController.DeleteEvaluacion)
	}
}
