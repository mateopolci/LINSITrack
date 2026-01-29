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

	// Rutas para administradores y profesores (gestión de evaluaciones)
	evaluaciones := router.Group("/evaluaciones")
	evaluaciones.Use(middleware.AuthMiddleware())
	evaluaciones.Use(middleware.RequireRole(models.RoleAdmin, models.RoleProfesor))
	{
		evaluaciones.GET("/", evaluacionController.GetAllEvaluaciones)
		evaluaciones.GET("/:id", evaluacionController.GetEvaluacionByID)
		evaluaciones.GET("/comision/:comisionId", evaluacionController.GetEvaluacionesByComisionID)
		evaluaciones.POST("/", evaluacionController.CreateEvaluacion)
		evaluaciones.PATCH("/:id", evaluacionController.UpdateEvaluacion)
	}

	// Solo admin puede eliminar evaluaciones
	adminOnlyEvaluaciones := router.Group("/evaluaciones")
	adminOnlyEvaluaciones.Use(middleware.AuthMiddleware())
	adminOnlyEvaluaciones.Use(middleware.RequireRole(models.RoleAdmin))
	{
		adminOnlyEvaluaciones.DELETE("/:id", evaluacionController.DeleteEvaluacion)
	}
}
