package services

import "github.com/Bright2704/battery-tracking/models"

type VisionService interface {
	CreateVision(vision *models.Vision) error
	GetVision(*string) (*models.Vision, error)
	GetAll() ([]*models.Vision, error)
	UpdateVision (*models.Vision) error
	DeleteVision (*string) error
}