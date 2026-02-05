package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupAnexoRoutes(router *gin.Engine, service *services.AnexoService) {
	anexoController := controllers.NewAnexoController(service)

	anexos := router.Group("/anexos")
	anexos.Use(middleware.AuthMiddleware())
	{
		// Endpoints de lectura - Todos los roles (Admin, Profesor, Alumno)
		anexos.GET("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor, models.RoleAlumno), anexoController.GetAnexoByID)
		anexos.GET("/tp/:tp_id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor, models.RoleAlumno), anexoController.GetAnexosByTpID)
		anexos.GET("/:id/archivos", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor, models.RoleAlumno), anexoController.GetAnexoArchivosByAnexoID)
		anexos.GET("/:id/archivo/download", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor, models.RoleAlumno), anexoController.DownloadAnexoArchivo)

		// Endpoints de escritura - Solo Admin y Profesor
		anexos.GET("/", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.GetAllAnexos)
		anexos.POST("/", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.CreateAnexo)
		anexos.PATCH("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.UpdateAnexo)
		anexos.DELETE("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.DeleteAnexo)
		anexos.POST("/:id/upload", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.UploadAnexoArchivo)
		anexos.DELETE("/:id/archivos", middleware.RequireRole(models.RoleAdmin, models.RoleProfesor), anexoController.DeleteAnexoArchivo)
	}
}
