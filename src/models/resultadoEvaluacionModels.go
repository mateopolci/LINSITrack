package models

type ResultadoEvaluacion struct {
	ID            int             `json:"id" gorm:"primaryKey;autoIncrement"`
	Nota          float64         `json:"nota" gorm:"column:nota;type:float;not null"`
	Devolucion    string          `json:"devolucion" gorm:"column:devolucion;type:text;null"`
	AlumnoID      int             `json:"alumno_id" gorm:"column:alumno_id;type:int;not null"`
	Alumno        Alumno          `json:"alumno" gorm:"foreignKey:AlumnoID;references:ID"`
	EvaluacionID  int             `json:"evaluacion_id" gorm:"column:evaluacion_id;type:int;not null"`
	Evaluacion    EvaluacionModel `json:"evaluacion" gorm:"foreignKey:EvaluacionID;references:ID"`
}

type ResultadoEvaluacionUpdateRequest struct {
	Nota          *float64 `json:"nota,omitempty"`
	Devolucion    *string  `json:"devolucion,omitempty"`
	Observaciones *string  `json:"observaciones,omitempty"`
	AlumnoID      *int     `json:"alumno_id,omitempty"`
	EvaluacionID  *int     `json:"evaluacion_id,omitempty"`
}
