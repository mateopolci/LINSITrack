package services

import (
	"fmt"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

type EvaluacionService struct {
	db                  *gorm.DB
	notificacionService *NotificacionService
}

func NewEvaluacionService(db *gorm.DB) *EvaluacionService {
	return &EvaluacionService{
		db:                  db,
		notificacionService: NewNotificacionService(db),
	}
}

func (s *EvaluacionService) GetAllEvaluaciones() ([]models.EvaluacionModel, error) {
	var evaluaciones []models.EvaluacionModel
	result := s.db.Preload("Comision").Preload("Comision.Materia").Find(&evaluaciones)
	if result.Error != nil {
		return nil, result.Error
	}
	return evaluaciones, nil
}

func (s *EvaluacionService) GetEvaluacionByID(id int) (*models.EvaluacionModel, error) {
	var evaluacion models.EvaluacionModel
	result := s.db.Preload("Comision").Preload("Comision.Materia").First(&evaluacion, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &evaluacion, nil
}

func (s *EvaluacionService) GetEvaluacionesByComisionID(comisionID int) ([]models.EvaluacionModel, error) {
	var evaluaciones []models.EvaluacionModel
	result := s.db.Preload("Comision").Preload("Comision.Materia").Where("comision_id = ?", comisionID).Find(&evaluaciones)
	if result.Error != nil {
		return nil, result.Error
	}
	return evaluaciones, nil
}

func (s *EvaluacionService) CreateEvaluacion(evaluacion *models.EvaluacionModel) error {
	result := s.db.Create(evaluacion)
	if result.Error != nil {
		return result.Error
	}

	// Get comision name for notification message
	var comision models.Comision
	if err := s.db.Preload("Materia").First(&comision, evaluacion.ComisionId).Error; err == nil {
		mensaje := fmt.Sprintf("Nueva evaluación programada para la materia %s (Comisión: %s) el %s. Temas: %s",
			comision.Materia.Nombre, comision.Nombre, evaluacion.FechaEvaluacion, evaluacion.Temas)
		s.notificacionService.NotifyAlumnosByComision(evaluacion.ComisionId, mensaje)
	}

	return nil
}

func (s *EvaluacionService) UpdateEvaluacion(id int, updateRequest *models.EvaluacionUpdateRequest) (*models.EvaluacionModel, error) {
	var evaluacion models.EvaluacionModel
	result := s.db.First(&evaluacion, id)
	if result.Error != nil {
		return nil, result.Error
	}

	oldComisionId := evaluacion.ComisionId

	if updateRequest.FechaEvaluacion != nil {
		evaluacion.FechaEvaluacion = *updateRequest.FechaEvaluacion
	}
	if updateRequest.Temas != nil {
		evaluacion.Temas = *updateRequest.Temas
	}
	if updateRequest.Observaciones != nil {
		evaluacion.Observaciones = *updateRequest.Observaciones
	}
	if updateRequest.ComisionId != nil {
		evaluacion.ComisionId = *updateRequest.ComisionId
	}

	result = s.db.Save(&evaluacion)
	if result.Error != nil {
		return nil, result.Error
	}

	// Notify students about the update
	var comision models.Comision
	if err := s.db.Preload("Materia").First(&comision, evaluacion.ComisionId).Error; err == nil {
		mensaje := fmt.Sprintf("Actualización de evaluación para %s (Comisión: %s) - Fecha: %s. Temas: %s",
			comision.Materia.Nombre, comision.Nombre, evaluacion.FechaEvaluacion, evaluacion.Temas)

		// Notify old comision if it changed
		if updateRequest.ComisionId != nil && oldComisionId != evaluacion.ComisionId {
			s.notificacionService.NotifyAlumnosByComision(oldComisionId, "Una evaluación ha sido reasignada a otra comisión")
		}

		s.notificacionService.NotifyAlumnosByComision(evaluacion.ComisionId, mensaje)
	}

	return &evaluacion, nil
}

func (s *EvaluacionService) DeleteEvaluacion(id int) error {
	result := s.db.Delete(&models.EvaluacionModel{}, id)
	return result.Error
}
