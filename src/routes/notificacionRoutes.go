package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupNotificacionRoutes(router *gin.Engine, service *services.NotificacionService) {
	notificacionController := controllers.NewNotificacionController(service)

	adminOnlyNotificaciones := router.Group("/notificaciones")
	adminOnlyNotificaciones.Use(middleware.AuthMiddleware())
	adminOnlyNotificaciones.Use(middleware.RequireRole(models.RoleAdmin))
	{
		// Rutas generales
		adminOnlyNotificaciones.GET("/", notificacionController.GetAllNotificaciones)
		adminOnlyNotificaciones.POST("/", notificacionController.CreateNotificacion)
		adminOnlyNotificaciones.GET("/:id", notificacionController.GetNotificacionByID)
		adminOnlyNotificaciones.PATCH("/:id", notificacionController.UpdateNotificacion)
		adminOnlyNotificaciones.DELETE("/:id", notificacionController.DeleteNotificacion)

		// Rutas específicas de alumno
		adminOnlyNotificaciones.GET("/alumnos/:alumnoId", notificacionController.GetNotificacionesByAlumnoID)
		adminOnlyNotificaciones.GET("/alumnos/:alumnoId/unread", notificacionController.GetUnreadNotificacionesByAlumnoID)
		adminOnlyNotificaciones.GET("/alumnos/:alumnoId/read", notificacionController.GetReadNotificacionesByAlumnoID)

		// Acciones
		adminOnlyNotificaciones.PATCH("/:id/mark-read", notificacionController.MarkNotificacionAsRead)
		adminOnlyNotificaciones.PATCH("/alumnos/:alumnoId/mark-all-read", notificacionController.MarkAllNotificacionAsReadByAlumnoID)
	}
}
