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

	// Student routes
	alumnoGroup := router.Group("/notificaciones/me")
	alumnoGroup.Use(middleware.AuthMiddleware())
	alumnoGroup.Use(middleware.RequireRole(models.RoleAlumno))
	{
		alumnoGroup.GET("", notificacionController.GetMyNotificaciones)
		alumnoGroup.GET("/unread", notificacionController.GetMyUnreadNotificaciones)
		alumnoGroup.GET("/read", notificacionController.GetMyReadNotificaciones)
		alumnoGroup.PATCH("/:id/mark-read", notificacionController.MarkMyNotificacionAsRead)
		alumnoGroup.PATCH("/mark-all-read", notificacionController.MarkAllMyNotificacionAsRead)
	}

	// Profesor & Admin bulk operations
	profesorAdminGroup := router.Group("/notificaciones")
	profesorAdminGroup.Use(middleware.AuthMiddleware())
	profesorAdminGroup.Use(middleware.RequireRole(models.RoleProfesor, models.RoleAdmin))
	{
		profesorAdminGroup.POST("/notify-materia", notificacionController.NotifyAlumnosByMateria)
		profesorAdminGroup.POST("/notify-comision", notificacionController.NotifyAlumnosByComision)
	}

	// Admin-only routes
	adminGroup := router.Group("/notificaciones")
	adminGroup.Use(middleware.AuthMiddleware())
	adminGroup.Use(middleware.RequireRole(models.RoleAdmin))
	{
		// General routes
		adminGroup.GET("", notificacionController.GetAllNotificaciones)
		adminGroup.POST("", notificacionController.CreateNotificacion)
		adminGroup.GET("/:id", notificacionController.GetNotificacionByID)
		adminGroup.PATCH("/:id", notificacionController.UpdateNotificacion)
		adminGroup.DELETE("/:id", notificacionController.DeleteNotificacion)
		adminGroup.PATCH("/:id/mark-read", notificacionController.MarkNotificacionAsRead)

		// Routes for specific students
		adminGroup.GET("/alumnos/:alumnoId", notificacionController.GetNotificacionesByAlumnoID)
		adminGroup.GET("/alumnos/:alumnoId/unread", notificacionController.GetUnreadNotificacionesByAlumnoID)
		adminGroup.GET("/alumnos/:alumnoId/read", notificacionController.GetReadNotificacionesByAlumnoID)
		adminGroup.PATCH("/alumnos/:alumnoId/mark-all-read", notificacionController.MarkAllNotificacionAsReadByAlumnoID)
	}
}
