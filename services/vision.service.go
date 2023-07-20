package services

import "golang/battery-tracking/models"

type VisionService interface {
	CreateVision(*models.Vision) error
	GetVision(*string) (*models.Vision, error)
	GetAll() ([]*models.Vision, error)
	UpdateVision(*models.Vision) error
	DeleteVision(*string) error
}