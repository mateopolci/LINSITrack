package services

import (
	"fmt"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

type ResultadoEvaluacionService struct {
	db                  *gorm.DB
	notificacionService *NotificacionService
}

func NewResultadoEvaluacionService(db *gorm.DB) *ResultadoEvaluacionService {
	return &ResultadoEvaluacionService{
		db:                  db,
		notificacionService: NewNotificacionService(db),
	}
}

func (s *ResultadoEvaluacionService) GetAllResultados() ([]models.ResultadoEvaluacion, error) {
	var resultados []models.ResultadoEvaluacion
	result := s.db.Preload("Alumno").Preload("Evaluacion").Preload("Evaluacion.Comision").Find(&resultados)
	if result.Error != nil {
		return nil, result.Error
	}
	return resultados, nil
}

func (s *ResultadoEvaluacionService) GetResultadoByID(id int) (*models.ResultadoEvaluacion, error) {
	var resultado models.ResultadoEvaluacion
	result := s.db.Preload("Alumno").Preload("Evaluacion").Preload("Evaluacion.Comision").First(&resultado, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resultado, nil
}

func (s *ResultadoEvaluacionService) GetResultadosByAlumnoID(alumnoID int) ([]models.ResultadoEvaluacion, error) {
	var resultados []models.ResultadoEvaluacion
	result := s.db.Preload("Alumno").Preload("Evaluacion").Preload("Evaluacion.Comision").Where("alumno_id = ?", alumnoID).Find(&resultados)
	if result.Error != nil {
		return nil, result.Error
	}
	return resultados, nil
}

func (s *ResultadoEvaluacionService) GetResultadosByEvaluacionID(evaluacionID int) ([]models.ResultadoEvaluacion, error) {
	var resultados []models.ResultadoEvaluacion
	result := s.db.Preload("Alumno").Preload("Evaluacion").Preload("Evaluacion.Comision").Where("evaluacion_id = ?", evaluacionID).Find(&resultados)
	if result.Error != nil {
		return nil, result.Error
	}
	return resultados, nil
}

func (s *ResultadoEvaluacionService) CreateResultado(resultado *models.ResultadoEvaluacion) error {
	result := s.db.Create(resultado)
	if result.Error != nil {
		return result.Error
	}

	// Get evaluacion details for notification
	var evaluacion models.EvaluacionModel
	if err := s.db.Preload("Comision").Preload("Comision.Materia").First(&evaluacion, resultado.EvaluacionID).Error; err == nil {
		mensaje := fmt.Sprintf("Resultado de evaluación disponible para %s - Nota: %.2f",
			evaluacion.Comision.Materia.Nombre, resultado.Nota)
		s.notificacionService.NotifyAlumno(resultado.AlumnoID, mensaje)
	}

	return nil
}

func (s *ResultadoEvaluacionService) UpdateResultado(id int, updateRequest *models.ResultadoEvaluacionUpdateRequest) (*models.ResultadoEvaluacion, error) {
	var resultado models.ResultadoEvaluacion
	result := s.db.First(&resultado, id)
	if result.Error != nil {
		return nil, result.Error
	}

	notaChanged := false
	devolucionChanged := false

	if updateRequest.Nota != nil {
		if resultado.Nota != *updateRequest.Nota {
			notaChanged = true
		}
		resultado.Nota = *updateRequest.Nota
	}
	if updateRequest.Devolucion != nil {
		if resultado.Devolucion != *updateRequest.Devolucion {
			devolucionChanged = true
		}
		resultado.Devolucion = *updateRequest.Devolucion
	}
	if updateRequest.AlumnoID != nil {
		resultado.AlumnoID = *updateRequest.AlumnoID
	}
	if updateRequest.EvaluacionID != nil {
		resultado.EvaluacionID = *updateRequest.EvaluacionID
	}

	result = s.db.Save(&resultado)
	if result.Error != nil {
		return nil, result.Error
	}

	// Notify student if nota or devolucion changed
	if notaChanged || devolucionChanged {
		var evaluacion models.EvaluacionModel
		if err := s.db.Preload("Comision").Preload("Comision.Materia").First(&evaluacion, resultado.EvaluacionID).Error; err == nil {
			mensaje := fmt.Sprintf("Actualización de resultado de evaluación para %s - Nota: %.2f",
				evaluacion.Comision.Materia.Nombre, resultado.Nota)
			s.notificacionService.NotifyAlumno(resultado.AlumnoID, mensaje)
		}
	}

	return &resultado, nil
}

func (s *ResultadoEvaluacionService) DeleteResultado(id int) error {
	result := s.db.Delete(&models.ResultadoEvaluacion{}, id)
	return result.Error
}
