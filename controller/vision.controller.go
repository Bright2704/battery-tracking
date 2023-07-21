package controller

import (
	"net/http"

	"golang/battery-tracking/models"
	"golang/battery-tracking/services"
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
	var vision models.Vision
	if err := ctx.ShouldBindJSON(&vision); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.VisionService.CreateVision(&vision)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":  "success"})
}

func (uc *VisionController) GetVision(ctx *gin.Context) {
	var visionserial_number string = ctx.Param("serial_number")
	related, err := uc.VisionService.GetVision(&visionserial_number)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, related)
}

func (uc *VisionController) GetVisionFromFile(ctx *gin.Context) {
	// ใช้ ctx.DefaultQuery เพื่อรับค่าของ serial_number จาก URL
	// ถ้าไม่มีค่าให้ใช้ค่าเริ่มต้นเป็นว่างๆ
	visionSerialNumber := ctx.DefaultQuery("serial_number", "")

	// ตรวจสอบว่า visionSerialNumber มีค่าหรือไม่
	if visionSerialNumber == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Serial number is required"})
		return
	}

	related, err := uc.VisionService.GetVision(&visionSerialNumber)
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
	var vision models.Vision
	if err := ctx. ShouldBindJSON(&vision); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.VisionService.UpdateVision(&vision)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *VisionController) DeleteVision(ctx *gin.Context) {
	var visionserial_number string = ctx.Param("serial_number")
	err := uc.VisionService.DeleteVision(&visionserial_number)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *VisionController) RegisterVisionRoutes(rg *gin.RouterGroup) {
	visionroute := rg.Group("/vision")
	visionroute.POST("/create", uc.CreateVision)
	visionroute.GET("/get/:serial_number", uc.GetVision)
	visionroute.GET("/getfromfile", uc.GetVisionFromFile)
	visionroute.GET("/getall", uc.GetAll)
	visionroute.PUT("/update", uc.UpdateVision)
	visionroute.DELETE("/delete/:serial_number", uc.DeleteVision)
}