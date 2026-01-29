package controllers

import (
	"net/http"
	"strconv"

	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

type ResultadoEvaluacionController struct {
	resultadoService *services.ResultadoEvaluacionService
}

func NewResultadoEvaluacionController(resultadoService *services.ResultadoEvaluacionService) *ResultadoEvaluacionController {
	return &ResultadoEvaluacionController{resultadoService: resultadoService}
}

func (c *ResultadoEvaluacionController) GetAllResultados(ctx *gin.Context) {
	resultados, err := c.resultadoService.GetAllResultados()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) GetResultadoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	resultado, err := c.resultadoService.GetResultadoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultado)
}

func (c *ResultadoEvaluacionController) GetResultadosByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}

	resultados, err := c.resultadoService.GetResultadosByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) GetResultadosByEvaluacionID(ctx *gin.Context) {
	evaluacionID, err := strconv.Atoi(ctx.Param("evaluacionId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid evaluacion ID"})
		return
	}

	resultados, err := c.resultadoService.GetResultadosByEvaluacionID(evaluacionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) CreateResultado(ctx *gin.Context) {
	var resultado models.ResultadoEvaluacion
	if err := ctx.ShouldBindJSON(&resultado); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Validaciones requeridas
	if resultado.AlumnoID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "el ID del alumno es obligatorio"})
		return
	}
	if resultado.EvaluacionID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "el ID de la evaluación es obligatorio"})
		return
	}

	if err := c.resultadoService.CreateResultado(&resultado); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resultado)
}

func (c *ResultadoEvaluacionController) UpdateResultado(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateRequest models.ResultadoEvaluacionUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Validaciones para campos que no pueden estar vacíos si se envían
	if updateRequest.Nota != nil && (*updateRequest.Nota < 0 || *updateRequest.Nota > 10) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "la nota debe estar entre 0 y 10"})
		return
	}
	if updateRequest.Devolucion != nil && *updateRequest.Devolucion == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "la devolución no puede estar vacía"})
		return
	}
	if updateRequest.Observaciones != nil && *updateRequest.Observaciones == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "las observaciones no pueden estar vacías"})
		return
	}

	resultado, err := c.resultadoService.UpdateResultado(id, &updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultado)
}

func (c *ResultadoEvaluacionController) DeleteResultado(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.resultadoService.DeleteResultado(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// GetMisResultados obtiene los resultados de evaluación del alumno autenticado
func (c *ResultadoEvaluacionController) GetMisResultados(ctx *gin.Context) {
	// Obtener el ID del alumno desde el contexto (establecido por el middleware de autenticación)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	alumnoID := int(userID.(float64))

	resultados, err := c.resultadoService.GetResultadosByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultados)
}

// GetMiResultadoByID obtiene un resultado específico del alumno autenticado
func (c *ResultadoEvaluacionController) GetMiResultadoByID(ctx *gin.Context) {
	// Obtener el ID del resultado desde los parámetros
	resultadoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Obtener el ID del alumno desde el contexto
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	alumnoID := int(userID.(float64))

	// Obtener el resultado
	resultado, err := c.resultadoService.GetResultadoByID(resultadoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verificar que el resultado pertenece al alumno autenticado
	if resultado.AlumnoID != alumnoID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para ver este resultado"})
		return
	}

	ctx.JSON(http.StatusOK, resultado)
}
