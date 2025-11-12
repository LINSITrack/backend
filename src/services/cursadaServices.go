package services

import (
	"errors"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

type CursadaService struct {
	db *gorm.DB
}

func NewCursadaService(db *gorm.DB) *CursadaService {
	return &CursadaService{db: db}
}

func (s *CursadaService) GetAllCursadas() ([]models.CursadaResponse, error) {
	var cursadas []models.Cursada
	result := s.db.Preload("Alumno").Preload("Comision").Preload("Comision.Materia").Find(&cursadas)
	if result.Error != nil {
		return nil, result.Error
	}

	responses := make([]models.CursadaResponse, 0, len(cursadas))
	for _, cursada := range cursadas {
		response := models.CursadaResponse{
			ID:             cursada.ID,
			AnoLectivo:     cursada.AnoLectivo,
			NotaFinal:      cursada.NotaFinal,
			NotaConceptual: cursada.NotaConceptual,
			Feedback:       cursada.Feedback,
			AlumnoID:       cursada.AlumnoID,
			Alumno: models.AlumnoResponse{
				ID:       cursada.Alumno.ID,
				Nombre:   cursada.Alumno.Nombre,
				Apellido: cursada.Alumno.Apellido,
				Legajo:   cursada.Alumno.Legajo,
				Email:    cursada.Alumno.Email,
			},
			ComisionID: cursada.ComisionID,
			Comision:   cursada.Comision,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *CursadaService) GetCursadaByID(id int) (*models.CursadaResponse, error) {
	var cursada models.Cursada
	result := s.db.Preload("Alumno").Preload("Comision").Preload("Comision.Materia").First(&cursada, id)
	if result.Error != nil {
		return nil, result.Error
	}

	response := &models.CursadaResponse{
		ID:             cursada.ID,
		AnoLectivo:     cursada.AnoLectivo,
		NotaFinal:      cursada.NotaFinal,
		NotaConceptual: cursada.NotaConceptual,
		Feedback:       cursada.Feedback,
		AlumnoID:       cursada.AlumnoID,
		Alumno: models.AlumnoResponse{
			ID:       cursada.Alumno.ID,
			Nombre:   cursada.Alumno.Nombre,
			Apellido: cursada.Alumno.Apellido,
			Legajo:   cursada.Alumno.Legajo,
			Email:    cursada.Alumno.Email,
		},
		ComisionID: cursada.ComisionID,
		Comision:   cursada.Comision,
	}

	return response, nil
}

func (s *CursadaService) GetCursadaByAlumnoID(alumnoID int) ([]models.CursadaResponse, error) {
	var cursadas []models.Cursada
	result := s.db.Preload("Alumno").Preload("Comision").Preload("Comision.Materia").Where("alumno_id = ?", alumnoID).Find(&cursadas)
	if result.Error != nil {
		return nil, result.Error
	}

	responses := make([]models.CursadaResponse, 0, len(cursadas))
	for _, cursada := range cursadas {
		response := models.CursadaResponse{
			ID:             cursada.ID,
			AnoLectivo:     cursada.AnoLectivo,
			NotaFinal:      cursada.NotaFinal,
			NotaConceptual: cursada.NotaConceptual,
			Feedback:       cursada.Feedback,
			AlumnoID:       cursada.AlumnoID,
			Alumno: models.AlumnoResponse{
				ID:       cursada.Alumno.ID,
				Nombre:   cursada.Alumno.Nombre,
				Apellido: cursada.Alumno.Apellido,
				Legajo:   cursada.Alumno.Legajo,
				Email:    cursada.Alumno.Email,
			},
			ComisionID: cursada.ComisionID,
			Comision:   cursada.Comision,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *CursadaService) CreateCursada(cursada *models.Cursada) error {
	var alumno models.Alumno
	if err := s.db.First(&alumno, cursada.AlumnoID).Error; err != nil {
		return errors.New("el alumno especificado no existe")
	}

	var comision models.Comision
	if err := s.db.First(&comision, cursada.ComisionID).Error; err != nil {
		return errors.New("la comisión especificada no existe")
	}

	return s.db.Create(cursada).Error
}

func (s *CursadaService) UpdateCursada(id int, updateRequest *models.CursadaUpdateRequest) (*models.Cursada, error) {
	var cursada models.Cursada
	result := s.db.First(&cursada, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if updateRequest.AnoLectivo != nil {
		cursada.AnoLectivo = *updateRequest.AnoLectivo
	}
	if updateRequest.NotaFinal != nil {
		cursada.NotaFinal = *updateRequest.NotaFinal
	}
	if updateRequest.NotaConceptual != nil {
		cursada.NotaConceptual = *updateRequest.NotaConceptual
	}
	if updateRequest.Feedback != nil {
		cursada.Feedback = *updateRequest.Feedback
	}
	if updateRequest.AlumnoID != nil {
		cursada.AlumnoID = *updateRequest.AlumnoID
	}
	if updateRequest.ComisionID != nil {
		cursada.ComisionID = *updateRequest.ComisionID
	}

	result = s.db.Save(&cursada)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cursada, nil
}

func (s *CursadaService) DeleteCursada(id int) error {
	result := s.db.Delete(&models.Cursada{}, id)
	return result.Error
}
