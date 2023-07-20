package controller

import (
	"net/http"

	"github.com/Bright2704/battery-tracking/models"
	"github.com/Bright2704/battery-tracking/services"
	"github.com/gin-gonic/gin"
)

type VisionController struct {
	VisionService services.VisionService
}

func New(visionservice services.VisionService) VisionController {
	return VisionController{
		VisionService: visionservice,
	}
}

func (uc *VisionController) CreateVision(ctx *gin.Context) {
	var vision *models.Vision
	if err := ctx.ShouldBindJSON(&vision); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.VisionService.CreateVision(vision)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":  "success"})
}

func (uc *VisionController) GetVision(ctx *gin.Context) {
	var visionrelated string = ctx.Param("related")
	related, err := uc.VisionService.GetVision(&visionrelated)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, related)
}

func (uc *VisionController) GetAll(ctx *gin.Context) {
	visions, err := uc.VisionService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, visions)
}

func (uc *VisionController) UpdateVision(ctx *gin.Context) {
	var vision *models.Vision
	if err := ctx. ShouldBindJSON(&vision); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.VisionService.UpdateVision(vision)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *VisionController) DeleteVision(ctx *gin.Context) {
	var visionrelated string = ctx.Param("related")
	err := uc.VisionService.DeleteVision(&visionrelated)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *VisionController) RegisterVisionRoutes(rg *gin.RouterGroup) {
	visionroute := rg.Group("/vision")
	visionroute.POST("/create", uc.CreateVision)
	visionroute.GET("/get/:related", uc.GetVision)
	visionroute.GET("/getall", uc.GetAll)
	visionroute.PUT("/update", uc.UpdateVision)
	visionroute.DELETE("/delete/:related", uc.DeleteVision)
}