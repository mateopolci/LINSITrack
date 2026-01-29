package services

import (
	"fmt"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

type TpService struct {
	db                  *gorm.DB
	notificacionService *NotificacionService
}

func NewTpService(db *gorm.DB) *TpService {
	return &TpService{
		db:                  db,
		notificacionService: NewNotificacionService(db),
	}
}

func (s *TpService) GetAllTps() ([]models.TpModel, error) {
	var tps []models.TpModel
	result := s.db.Preload("Comision").Find(&tps)
	if result.Error != nil {
		return nil, result.Error
	}
	return tps, nil
}

func (s *TpService) GetTpByID(id int) (*models.TpModel, error) {
	var tp models.TpModel
	result := s.db.Preload("Comision").First(&tp, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tp, nil
}

func (s *TpService) CreateTp(tp *models.TpModel) error {
	result := s.db.Create(tp)
	if result.Error != nil {
		return result.Error
	}

	// Get comision and materia for notification message
	var comision models.Comision
	if err := s.db.Preload("Materia").First(&comision, tp.ComisionId).Error; err == nil {
		mensaje := fmt.Sprintf("Nuevo trabajo práctico disponible para %s (Comisión: %s). Fecha de entrega: %s",
			comision.Materia.Nombre, comision.Nombre, tp.FechaHoraEntrega.Format("02/01/2006 15:04"))
		s.notificacionService.NotifyAlumnosByComision(tp.ComisionId, mensaje)
	}

	return nil
}

func (s *TpService) UpdateTp(id int, updateRequest *models.TpUpdateRequest) (*models.TpModel, error) {
	var tp models.TpModel
	result := s.db.First(&tp, id)
	if result.Error != nil {
		return nil, result.Error
	}

	oldComisionId := tp.ComisionId

	if updateRequest.Consigna != nil {
		tp.Consigna = *updateRequest.Consigna
	}
	if updateRequest.FechaHoraEntrega != nil {
		tp.FechaHoraEntrega = *updateRequest.FechaHoraEntrega
	}
	if updateRequest.Vigente != nil {
		tp.Vigente = *updateRequest.Vigente
	}
	if updateRequest.ComisionId != nil {
		tp.ComisionId = *updateRequest.ComisionId
	}

	saveResult := s.db.Save(&tp)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	// Notify students about the update
	var comision models.Comision
	if err := s.db.Preload("Materia").First(&comision, tp.ComisionId).Error; err == nil {
		mensaje := fmt.Sprintf("Actualización de trabajo práctico para %s (Comisión: %s). Fecha de entrega: %s",
			comision.Materia.Nombre, comision.Nombre, tp.FechaHoraEntrega.Format("02/01/2006 15:04"))

		// Notify old comision if it changed
		if updateRequest.ComisionId != nil && oldComisionId != tp.ComisionId {
			s.notificacionService.NotifyAlumnosByComision(oldComisionId, "Un trabajo práctico ha sido reasignado a otra comisión")
		}

		s.notificacionService.NotifyAlumnosByComision(tp.ComisionId, mensaje)
	}

	return &tp, nil
}

func (s *TpService) DeleteTp(id int) error {
	result := s.db.Delete(&models.TpModel{}, id)
	return result.Error
}

// GetTpsByAlumnoID obtiene todos los TPs de las comisiones a las que el alumno está inscrito
func (s *TpService) GetTpsByAlumnoID(alumnoID int) ([]models.TpModel, error) {
	var tps []models.TpModel

	// Obtener los TPs a través de las cursadas del alumno
	result := s.db.Table("tp_models").
		Select("DISTINCT tp_models.*").
		Joins("JOIN cursadas ON tp_models.comision_id = cursadas.comision_id").
		Where("cursadas.alumno_id = ? AND tp_models.vigente = ?", alumnoID, true).
		Preload("Comision").
		Preload("Comision.Materia").
		Order("tp_models.fecha_entrega DESC").
		Find(&tps)

	if result.Error != nil {
		return nil, result.Error
	}

	return tps, nil
}
