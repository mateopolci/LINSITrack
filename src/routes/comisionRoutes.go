package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupComisionRoutes(router *gin.Engine, service *services.ComisionService) {
	comisionController := controllers.NewComisionController(service)
	
	adminOnlyComisiones := router.Group("/comisiones")
	adminOnlyComisiones.Use(middleware.AuthMiddleware())
	adminOnlyComisiones.Use(middleware.RequireRole(models.RoleAdmin))
	{
		adminOnlyComisiones.GET("/", comisionController.GetAllComisiones)
		adminOnlyComisiones.GET("/:id", comisionController.GetComisionByID)
		adminOnlyComisiones.GET("/materia/:materiaId", comisionController.GetComisionesByMateriaID)
		adminOnlyComisiones.POST("/", comisionController.CreateComision)
		adminOnlyComisiones.PATCH("/:id", comisionController.UpdateComision)
		adminOnlyComisiones.DELETE("/:id", comisionController.DeleteComision)
	}
}