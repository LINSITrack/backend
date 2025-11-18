package services

import (
	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

type TpService struct {
	db *gorm.DB
}

func NewTpService(db *gorm.DB) *TpService {
	return &TpService{db: db}
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
	return result.Error
}

func (s *TpService) UpdateTp(id int, updateRequest *models.TpUpdateRequest) (*models.TpModel, error) {
	var tp models.TpModel
	result := s.db.First(&tp, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if updateRequest.Consigna != nil {
		tp.Consigna = *updateRequest.Consigna
	}
	if updateRequest.FechaHoraEntrega != nil {
		tp.FechaHoraEntrega = *updateRequest.FechaHoraEntrega
	}
	if updateRequest.Vigente != nil {
		tp.Vigente = *updateRequest.Vigente
	}
	if updateRequest.Nota != nil {
		tp.Nota = *updateRequest.Nota
	}
	if updateRequest.Devolucion != nil {
		tp.Devolucion = *updateRequest.Devolucion
	}
	if updateRequest.ComisionId != nil {
		tp.ComisionId = *updateRequest.ComisionId
	}

	saveResult := s.db.Save(&tp)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	return &tp, nil
}

func (s *TpService) DeleteTp(id int) error {
	result := s.db.Delete(&models.TpModel{}, id)
	return result.Error
}