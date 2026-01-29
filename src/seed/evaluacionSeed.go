package seed

import (
	"log"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

func EvaluacionSeed(db *gorm.DB) {
	evaluaciones := []models.EvaluacionModel{
		// Algoritmos y estructuras de datos - S11 (Comision ID: 1)
		{
			FechaEvaluacion: "2024-03-15",
			Temas:           "Algoritmos de ordenamiento: burbuja, inserción, selección. Complejidad temporal y espacial.",
			Observaciones:   "Llevar notebook",
			ComisionId:      1,
		},
		{
			FechaEvaluacion: "2024-04-10",
			Temas:           "Estructuras de datos: listas, pilas, colas. Implementación y casos de uso.",
			Observaciones:   "Llevar notebook",
			ComisionId:      1,
		},

		// Algoritmos y estructuras de datos - S12 (Comision ID: 2)
		{
			FechaEvaluacion: "2024-03-18",
			Temas:           "Árboles binarios: recorridos, inserción, eliminación. Árboles de búsqueda binaria.",
			Observaciones:   "",
			ComisionId:      2,
		},

		// Sistemas operativos - S21 (Comision ID: 3)
		{
			FechaEvaluacion: "2024-03-22",
			Temas:           "Procesos y threads: creación, sincronización, comunicación entre procesos.",
			Observaciones:   "Llevar notebook",
			ComisionId:      3,
		},
		{
			FechaEvaluacion: "2024-05-10",
			Temas:           "Gestión de memoria: paginación, segmentación, memoria virtual.",
			Observaciones:   "Llevar notebook",
			ComisionId:      3,
		},

		// Bases de datos - S31 (Comision ID: 4)
		{
			FechaEvaluacion: "2024-04-05",
			Temas:           "Modelo relacional: normalización, dependencias funcionales, formas normales.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      4,
		},
		{
			FechaEvaluacion: "2024-05-20",
			Temas:           "SQL avanzado: subconsultas, joins, funciones agregadas, vistas.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      4,
		},

		// Bases de datos - S32 (Comision ID: 5)
		{
			FechaEvaluacion: "2024-04-08",
			Temas:           "Diseño de bases de datos: diagramas ER, cardinalidades, restricciones.",
			Observaciones:   "",
			ComisionId:      5,
		},

		// Desarrollo de software - S31 (Comision ID: 6)
		{
			FechaEvaluacion: "2024-03-25",
			Temas:           "Paradigmas de programación: POO, encapsulamiento, herencia, polimorfismo.",
			Observaciones:   "Llevar notebook",
			ComisionId:      6,
		},
		{
			FechaEvaluacion: "2024-05-15",
			Temas:           "Patrones de diseño: Singleton, Factory, Observer, MVC.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      6,
		},

		// Ingeniería y calidad de software - S41 (Comision ID: 7)
		{
			FechaEvaluacion: "2024-04-12",
			Temas:           "Metodologías ágiles: Scrum, Kanban, historias de usuario, planning poker.",
			Observaciones:   "",
			ComisionId:      7,
		},

		// Administración SI - S41 (Comision ID: 8)
		{
			FechaEvaluacion: "2024-04-18",
			Temas:           "ITIL: gestión de servicios, incidentes, problemas, cambios.",
			Observaciones:   "",
			ComisionId:      8,
		},

		// Ciencia de datos - S51 (Comision ID: 9)
		{
			FechaEvaluacion: "2024-03-30",
			Temas:           "Análisis exploratorio de datos: estadística descriptiva, visualizaciones, outliers.",
			Observaciones:   "Llevar notebook",
			ComisionId:      9,
		},
		{
			FechaEvaluacion: "2024-05-25",
			Temas:           "Machine Learning: regresión lineal, validación cruzada, métricas de evaluación.",
			Observaciones:   "Llevar notebook",
			ComisionId:      9,
		},

		// Inteligencia artificial - S51 (Comision ID: 10)
		{
			FechaEvaluacion: "2024-04-22",
			Temas:           "Búsqueda en espacios de estados: BFS, DFS, A*, heurísticas.",
			Observaciones:   "",
			ComisionId:      10,
		},

		// Proyecto final - S51 (Comision ID: 11)
		{
			FechaEvaluacion: "2024-05-30",
			Temas:           "Presentación de propuesta de proyecto: objetivos, alcance, tecnologías.",
			Observaciones:   "",
			ComisionId:      11,
		},
	}

	for _, evaluacion := range evaluaciones {
		var existingEvaluacion models.EvaluacionModel
		result := db.Where("fecha_evaluacion = ? AND comision_id = ? AND temas = ?",
			evaluacion.FechaEvaluacion, evaluacion.ComisionId, evaluacion.Temas).First(&existingEvaluacion)

		if result.Error == nil {
			log.Printf("Evaluación for comision %d on %s already exists", evaluacion.ComisionId, evaluacion.FechaEvaluacion)
		} else {
			if err := db.Create(&evaluacion).Error; err != nil {
				log.Printf("Failed to create evaluación for comision %d: %v", evaluacion.ComisionId, err)
			} else {
				log.Printf("Evaluación created successfully for comision %d on %s", evaluacion.ComisionId, evaluacion.FechaEvaluacion)
			}
		}
	}

	// Log
	log.Println("Evaluacion seed completed successfully")
}
