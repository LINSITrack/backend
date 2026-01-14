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
			Nota:            8.5,
			Devolucion:      "Excelente comprensión de los algoritmos básicos. Mejorar análisis de complejidad espacial.",
			Observaciones:   "Llevar notebook",
			ComisionId:      1,
		},
		{
			FechaEvaluacion: "2024-04-10",
			Temas:           "Estructuras de datos: listas, pilas, colas. Implementación y casos de uso.",
			Nota:            0, // Sin nota aún
			Devolucion:      "",
			Observaciones:   "Llevar notebook",
			ComisionId:      1,
		},

		// Algoritmos y estructuras de datos - S12 (Comision ID: 2)
		{
			FechaEvaluacion: "2024-03-18",
			Temas:           "Árboles binarios: recorridos, inserción, eliminación. Árboles de búsqueda binaria.",
			Nota:            7.0,
			Devolucion:      "Buen manejo de la teoría, necesita practicar más la implementación.",
			Observaciones:   "",
			ComisionId:      2,
		},

		// Sistemas operativos - S21 (Comision ID: 3)
		{
			FechaEvaluacion: "2024-03-22",
			Temas:           "Procesos y threads: creación, sincronización, comunicación entre procesos.",
			Nota:            9.2,
			Devolucion:      "Excelente dominio conceptual y práctico. Muy buena resolución de ejercicios.",
			Observaciones:   "Llevar notebook",
			ComisionId:      3,
		},
		{
			FechaEvaluacion: "2024-05-10",
			Temas:           "Gestión de memoria: paginación, segmentación, memoria virtual.",
			Nota:            0, // Sin calificar
			Devolucion:      "",
			Observaciones:   "Llevar notebook",
			ComisionId:      3,
		},

		// Bases de datos - S31 (Comision ID: 4)
		{
			FechaEvaluacion: "2024-04-05",
			Temas:           "Modelo relacional: normalización, dependencias funcionales, formas normales.",
			Nota:            6.8,
			Devolucion:      "Comprende los conceptos básicos pero necesita reforzar normalización de 3FN.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      4,
		},
		{
			FechaEvaluacion: "2024-05-20",
			Temas:           "SQL avanzado: subconsultas, joins, funciones agregadas, vistas.",
			Nota:            8.8,
			Devolucion:      "Muy buen manejo de SQL. Consultas complejas resueltas correctamente.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      4,
		},

		// Bases de datos - S32 (Comision ID: 5)
		{
			FechaEvaluacion: "2024-04-08",
			Temas:           "Diseño de bases de datos: diagramas ER, cardinalidades, restricciones.",
			Nota:            7.5,
			Devolucion:      "Buen diseño general, mejorar especificación de restricciones de integridad.",
			Observaciones:   "",
			ComisionId:      5,
		},

		// Desarrollo de software - S31 (Comision ID: 6)
		{
			FechaEvaluacion: "2024-03-25",
			Temas:           "Paradigmas de programación: POO, encapsulamiento, herencia, polimorfismo.",
			Nota:            0, // Pendiente
			Devolucion:      "",
			Observaciones:   "Llevar notebook",
			ComisionId:      6,
		},
		{
			FechaEvaluacion: "2024-05-15",
			Temas:           "Patrones de diseño: Singleton, Factory, Observer, MVC.",
			Nota:            9.0,
			Devolucion:      "Excelente aplicación de patrones. Código limpio y bien documentado.",
			Observaciones:   "Llevar calculadora",
			ComisionId:      6,
		},

		// Ingeniería y calidad de software - S41 (Comision ID: 7)
		{
			FechaEvaluacion: "2024-04-12",
			Temas:           "Metodologías ágiles: Scrum, Kanban, historias de usuario, planning poker.",
			Nota:            8.3,
			Devolucion:      "Buena comprensión de metodologías ágiles. Mejorar estimación de tareas.",
			Observaciones:   "",
			ComisionId:      7,
		},

		// Administración SI - S41 (Comision ID: 8)
		{
			FechaEvaluacion: "2024-04-18",
			Temas:           "ITIL: gestión de servicios, incidentes, problemas, cambios.",
			Nota:            7.8,
			Devolucion:      "Conceptos claros de ITIL. Practicar más casos de estudio reales.",
			Observaciones:   "",
			ComisionId:      8,
		},

		// Ciencia de datos - S51 (Comision ID: 9)
		{
			FechaEvaluacion: "2024-03-30",
			Temas:           "Análisis exploratorio de datos: estadística descriptiva, visualizaciones, outliers.",
			Nota:            0, // Sin nota
			Devolucion:      "",
			Observaciones:   "Llevar notebook",
			ComisionId:      9,
		},
		{
			FechaEvaluacion: "2024-05-25",
			Temas:           "Machine Learning: regresión lineal, validación cruzada, métricas de evaluación.",
			Nota:            8.7,
			Devolucion:      "Excelente implementación de modelos. Interpretación correcta de métricas.",
			Observaciones:   "Llevar notebook",
			ComisionId:      9,
		},

		// Inteligencia artificial - S51 (Comision ID: 10)
		{
			FechaEvaluacion: "2024-04-22",
			Temas:           "Búsqueda en espacios de estados: BFS, DFS, A*, heurísticas.",
			Nota:            7.2,
			Devolucion:      "Algoritmos implementados correctamente. Mejorar diseño de heurísticas.",
			Observaciones:   "",
			ComisionId:      10,
		},

		// Proyecto final - S51 (Comision ID: 11)
		{
			FechaEvaluacion: "2024-05-30",
			Temas:           "Presentación de propuesta de proyecto: objetivos, alcance, tecnologías.",
			Nota:            0, // Evaluación en curso
			Devolucion:      "",
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
