package models

type EvaluacionModel struct {
	ID              int           `json:"id" gorm:"primaryKey;autoIncrement"`
	FechaEvaluacion string        `json:"fecha_evaluacion" gorm:"column:fecha_evaluacion;type:date;not null"`
	Temas           string        `json:"temas" gorm:"column:temas;type:text;not null"`
	Observaciones   string        `json:"observaciones" gorm:"column:observaciones;type:text;null"`
	ComisionId      int           `json:"comision_id" gorm:"column:comision_id;type:int;not null"`
	Comision        Comision `json:"comision" gorm:"foreignKey:ComisionId;references:ID"`
}

type EvaluacionUpdateRequest struct {
	FechaEvaluacion *string  `json:"fecha_evaluacion,omitempty"`
	Temas           *string  `json:"temas,omitempty"`
	Observaciones   *string  `json:"observaciones,omitempty"`
	ComisionId      *int     `json:"comision_id,omitempty"`
}