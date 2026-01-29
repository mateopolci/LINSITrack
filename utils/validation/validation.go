package validation

import (
	"errors"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

// ValidateEmailUniqueness verifica que el email no exista en ninguna de las tres tablas
func ValidateEmailUniqueness(db *gorm.DB, email string, excludeTable string, excludeID string) error {
	// Verificar en Admin
	if excludeTable != "admin" {
		var admin models.Admin
		err := db.Where("email = ?", email).First(&admin).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else if excludeID != "" {
		// Si estamos actualizando un admin, excluir el registro actual
		var admin models.Admin
		err := db.Where("email = ? AND id <> ?", email, excludeID).First(&admin).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Verificar en Profesor
	if excludeTable != "profesor" {
		var profesor models.Profesor
		err := db.Where("email = ?", email).First(&profesor).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else if excludeID != "" {
		var profesor models.Profesor
		err := db.Where("email = ? AND id <> ?", email, excludeID).First(&profesor).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Verificar en Alumno
	if excludeTable != "alumno" {
		var alumno models.Alumno
		err := db.Where("email = ?", email).First(&alumno).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else if excludeID != "" {
		var alumno models.Alumno
		err := db.Where("email = ? AND id <> ?", email, excludeID).First(&alumno).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese email")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}

// ValidateLegajoUniqueness verifica que el legajo no exista en profesores o alumnos
func ValidateLegajoUniqueness(db *gorm.DB, legajo string, excludeTable string, excludeID string) error {
	// Verificar en Profesor
	if excludeTable != "profesor" {
		var profesor models.Profesor
		err := db.Where("legajo = ?", legajo).First(&profesor).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese legajo")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else if excludeID != "" {
		var profesor models.Profesor
		err := db.Where("legajo = ? AND id <> ?", legajo, excludeID).First(&profesor).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese legajo")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Verificar en Alumno
	if excludeTable != "alumno" {
		var alumno models.Alumno
		err := db.Where("legajo = ?", legajo).First(&alumno).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese legajo")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else if excludeID != "" {
		var alumno models.Alumno
		err := db.Where("legajo = ? AND id <> ?", legajo, excludeID).First(&alumno).Error
		if err == nil {
			return errors.New("ya existe una cuenta con ese legajo")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}
