package services

import "golang/battery-tracking/models"

type VisionService interface {
	CreateVision(*models.Vision) error
	GetVision(Serial_number *string) (*models.Vision, error)
	GetAll() ([]*models.Vision, error)
	UpdateVision(vision *models.Vision,visionSerialNumber *string) error
	UpdateStageProcessFlags(vision *models.Vision,visionSerialNumber *string) error
	DeleteVision(*string) error
	GetVisionFromFile(filePath string) (*models.Vision, error)
}