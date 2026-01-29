package controllers

import (
	"net/http"
	"strconv"

	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

type NotificacionController struct {
	notificacionService *services.NotificacionService
}

func NewNotificacionController(notificacionService *services.NotificacionService) *NotificacionController {
	return &NotificacionController{notificacionService: notificacionService}
}

func (c *NotificacionController) GetAllNotificaciones(ctx *gin.Context) {
	notificaciones, err := c.notificacionService.GetAllNotificaciones()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificaciones)
}

func (c *NotificacionController) GetNotificacionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	notificacion, err := c.notificacionService.GetNotificacionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificacion)
}

func (c *NotificacionController) GetNotificacionesByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}

	notificaciones, err := c.notificacionService.GetNotificacionesByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificaciones)
}

func (c *NotificacionController) GetUnreadNotificacionesByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}

	notificaciones, err := c.notificacionService.GetUnreadNotificacionesByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificaciones)
}

func (c *NotificacionController) GetReadNotificacionesByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}
	notificaciones, err := c.notificacionService.GetReadNotificacionesByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificaciones)
}

func (c *NotificacionController) MarkNotificacionAsRead(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.notificacionService.MarkNotificacionAsRead(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notificación marcada como leída"})
}

func (c *NotificacionController) MarkAllNotificacionAsReadByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}
	if err := c.notificacionService.MarkAllNotificacionAsReadByAlumnoID(alumnoID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Todas las notificaciones marcadas como leídas"})
}

func (c *NotificacionController) CreateNotificacion(ctx *gin.Context) {
	var notificacion models.Notificacion
	if err := ctx.ShouldBindJSON(&notificacion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}
	if notificacion.Mensaje == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El mensaje es obligatorio"})
		return
	}
	if notificacion.AlumnoID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID del alumno es obligatorio"})
		return
	}
	notificacion.Leida = false // Toda notificación nueva es no leída por defecto
	if notificacion.FechaHora.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "La fecha y hora son obligatorias"})
		return
	}
	if err := c.notificacionService.CreateNotificacion(&notificacion); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, notificacion)
}

func (c *NotificacionController) UpdateNotificacion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateRequest models.NotificacionUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	if updateRequest.Mensaje != nil && *updateRequest.Mensaje == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El mensaje no puede estar vacío"})
		return
	}
	notificacion, err := c.notificacionService.UpdateNotificacion(id, &updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificacion)
}

func (c *NotificacionController) DeleteNotificacion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.notificacionService.DeleteNotificacion(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notificación eliminada correctamente"})
}

func (c *NotificacionController) NotifyAlumnosByMateria(ctx *gin.Context) {
	var request models.NotifyByMateriaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	if request.Mensaje == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El mensaje es obligatorio"})
		return
	}
	if request.MateriaID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID de materia es obligatorio"})
		return
	}

	count, err := c.notificacionService.NotifyAlumnosByMateria(request.MateriaID, request.Mensaje)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":             "Notificaciones enviadas correctamente",
		"alumnos_notificados": count,
	})
}

func (c *NotificacionController) NotifyAlumnosByComision(ctx *gin.Context) {
	var request models.NotifyByComisionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	if request.Mensaje == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El mensaje es obligatorio"})
		return
	}
	if request.ComisionID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID de comisión es obligatorio"})
		return
	}

	count, err := c.notificacionService.NotifyAlumnosByComision(request.ComisionID, request.Mensaje)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":             "Notificaciones enviadas correctamente",
		"alumnos_notificados": count,
	})
}

// GetMyNotificaciones returns all notifications for the authenticated student
func (c *NotificacionController) GetMyNotificaciones(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// El JWT devuelve el ID como float64
	var userID int
	switch v := userIDInterface.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	notificaciones, err := c.notificacionService.GetNotificacionesByAlumnoID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notificaciones)
}

// GetMyUnreadNotificaciones returns unread notifications for the authenticated student
func (c *NotificacionController) GetMyUnreadNotificaciones(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var userID int
	switch v := userIDInterface.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	notificaciones, err := c.notificacionService.GetUnreadNotificacionesByAlumnoID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notificaciones)
}

// GetMyReadNotificaciones returns read notifications for the authenticated student
func (c *NotificacionController) GetMyReadNotificaciones(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var userID int
	switch v := userIDInterface.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	notificaciones, err := c.notificacionService.GetReadNotificacionesByAlumnoID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notificaciones)
}

// MarkMyNotificacionAsRead marks a notification as read for the authenticated student
func (c *NotificacionController) MarkMyNotificacionAsRead(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var userID int
	switch v := userIDInterface.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Verify that the notification belongs to the user
	notif, err := c.notificacionService.GetNotificacionByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Notificación no encontrada"})
		return
	}
	if notif.AlumnoID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para modificar esta notificación"})
		return
	}

	if err := c.notificacionService.MarkNotificacionAsRead(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notificación marcada como leída"})
}

// MarkAllMyNotificacionAsRead marks all notifications as read for the authenticated student
func (c *NotificacionController) MarkAllMyNotificacionAsRead(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var userID int
	switch v := userIDInterface.(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	if err := c.notificacionService.MarkAllNotificacionAsReadByAlumnoID(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Todas las notificaciones marcadas como leídas"})
}
