package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupTpRoutes(router *gin.Engine, service *services.TpService) {
	tpController := controllers.NewTpController(service)

	tps := router.Group("/tps")
	tps.Use(middleware.AuthMiddleware())
	{
		// Endpoint para alumnos: obtener sus propios TPs
		tps.GET("/mis-tps", middleware.RequireRole(models.RoleAlumno), tpController.GetMyTps)

		// Endpoints admin/profesor
		tps.GET("/", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), tpController.GetAllTps)
		tps.GET("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), tpController.GetTpByID)

		tps.POST("/", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), tpController.CreateTp)
		tps.PATCH("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), tpController.UpdateTp)

		tps.DELETE("/:id", middleware.RequireRole(models.RoleAdmin), tpController.DeleteTp)
	}
}
