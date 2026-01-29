package seed

import (
	"log"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

func ResultadoEvaluacionSeed(db *gorm.DB) {
	resultados := []models.ResultadoEvaluacion{
		// Resultados para la primera evaluación (ID: 1) - Algoritmos S11
		{
			Nota:          8.5,
			Devolucion:    "Excelente comprensión de los algoritmos de ordenamiento. Buena explicación de la complejidad temporal.",
			AlumnoID:      1, 
			EvaluacionID:  1,
		},
		{
			Nota:          7.0,
			Devolucion:    "Buen trabajo en general. Algunos errores en el análisis de complejidad espacial.",
			AlumnoID:      2, 
			EvaluacionID:  1,
		},
		{
			Nota:          9.0,
			Devolucion:    "Muy buen análisis y comprensión profunda de los temas. Ejemplos claros y bien explicados.",
			AlumnoID:      3, 
			EvaluacionID:  1,
		},

		// Resultados para la segunda evaluación (ID: 2) - Algoritmos S11
		{
			Nota:          8.0,
			Devolucion:    "Buena implementación de las estructuras de datos. Código limpio y bien documentado.",
			AlumnoID:      1, 
			EvaluacionID:  2,
		},
		{
			Nota:          6.5,
			Devolucion:    "Implementación correcta pero le falta optimización. Revisar casos edge.",
			AlumnoID:      2, 
			EvaluacionID:  2,
		},

		// Resultados para evaluación de Sistemas Operativos (ID: 4) - S21
		{
			Nota:          7.5,
			Devolucion:    "Buena comprensión de procesos y threads. Explicación clara de la sincronización.",
			AlumnoID:      4, 
			EvaluacionID:  4,
		},
		{
			Nota:          8.5,
			Devolucion:    "Excelente manejo de los conceptos de concurrencia. Muy buenos ejemplos prácticos.",
			AlumnoID:      5, 
			EvaluacionID:  4,
		},

		// Resultados para evaluación de Bases de Datos (ID: 6) - S31
		{
			Nota:          9.5,
			Devolucion:    "Excelente diseño de base de datos. Normalización perfecta y consultas muy eficientes.",
			AlumnoID:      6, 
			EvaluacionID:  6,
		},
		{
			Nota:          7.0,
			Devolucion:    "Buen diseño general pero con algunos errores en la normalización a 3FN.",
			AlumnoID:      7, 
			EvaluacionID:  6,
		},
	}

	for _, resultado := range resultados {
		if err := db.Create(&resultado).Error; err != nil {
			log.Printf("Error al crear resultado de evaluación: %v\n", err)
		}
	}

	log.Println("✅ Resultados de evaluación creados exitosamente")
}
