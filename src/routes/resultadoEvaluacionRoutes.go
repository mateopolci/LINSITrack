package routes

import (
	"github.com/LINSITrack/backend/src/controllers"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

func SetupResultadoEvaluacionRoutes(router *gin.Engine, service *services.ResultadoEvaluacionService) {
	resultadoController := controllers.NewResultadoEvaluacionController(service)

	// Rutas para administradores y profesores (gestión completa)
	adminProfesorResultados := router.Group("/resultado-evaluacion")
	adminProfesorResultados.Use(middleware.AuthMiddleware())
	adminProfesorResultados.Use(middleware.RequireRole(models.RoleAdmin, models.RoleProfesor))
	{
		adminProfesorResultados.GET("", resultadoController.GetAllResultados)
		adminProfesorResultados.GET("/:id", resultadoController.GetResultadoByID)
		adminProfesorResultados.GET("/alumno/:alumnoId", resultadoController.GetResultadosByAlumnoID)
		adminProfesorResultados.GET("/evaluacion/:evaluacionId", resultadoController.GetResultadosByEvaluacionID)
		adminProfesorResultados.POST("", resultadoController.CreateResultado)
		adminProfesorResultados.PUT("/:id", resultadoController.UpdateResultado)
	}

	// Solo admin puede eliminar resultados
	adminOnlyResultados := router.Group("/resultado-evaluacion")
	adminOnlyResultados.Use(middleware.AuthMiddleware())
	adminOnlyResultados.Use(middleware.RequireRole(models.RoleAdmin))
	{
		adminOnlyResultados.DELETE("/:id", resultadoController.DeleteResultado)
	}

	// Rutas para alumnos (solo ver sus propios resultados)
	alumnoResultados := router.Group("/mis-resultados")
	alumnoResultados.Use(middleware.AuthMiddleware())
	alumnoResultados.Use(middleware.RequireRole(models.RoleAlumno))
	{
		alumnoResultados.GET("", resultadoController.GetMisResultados)
		alumnoResultados.GET("/:id", resultadoController.GetMiResultadoByID)
	}
}
