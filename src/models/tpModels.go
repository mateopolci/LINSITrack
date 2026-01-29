package models

import "time"

type TpModel struct {
	ID               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Consigna         string    `json:"consigna" gorm:"column:consigna;type:text;not null"`
	FechaHoraEntrega time.Time `json:"fecha_entrega" gorm:"column:fecha_entrega;type:date;not null"`
	Vigente          bool      `json:"vigente" gorm:"column:vigente;type:boolean;not null;default:true"`
	ComisionId       int       `json:"comision_id" gorm:"column:comision_id;type:int;not null"`
	Comision         Comision  `json:"comision" gorm:"foreignKey:ComisionId;references:ID"`
}

type TpUpdateRequest struct {
	Consigna         *string    `json:"consigna,omitempty"`
	FechaHoraEntrega *time.Time `json:"fecha_entrega,omitempty"`
	Vigente          *bool      `json:"vigente,omitempty"`
	ComisionId       *int       `json:"comision_id,omitempty"`
}