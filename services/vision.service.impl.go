package services

import (
	"context"
	"encoding/json"
	"errors"
	

	
	"io/ioutil"
	"os"

	"golang/battery-tracking/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
)

type VisionServiceImpl struct {
	visioncollection *mongo.Collection
	ctx              context.Context
}

func NewVisionService(visioncollection *mongo.Collection, ctx context.Context) VisionService {
	return &VisionServiceImpl{
		visioncollection: visioncollection,
		ctx:              ctx,
	}
}


func (u *VisionServiceImpl) CreateVision(vision *models.Vision) error {
	_, err := u.visioncollection.InsertOne(u.ctx, vision)
	return err
}

func (u *VisionServiceImpl) GetVision(serial_number *string) (*models.Vision, error) {
	var vision *models.Vision
	query := bson.D{bson.E{Key: "serial_number", Value: serial_number}}
	err := u.visioncollection.FindOne(u.ctx, query).Decode(&vision)
	return vision, err
}

func (u *VisionServiceImpl) GetVisionFromFile(filePath string) (*models.Vision, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a Vision struct
	var vision *models.Vision
	err = json.Unmarshal(fileData, &vision)
	if err != nil {
		return nil, err
	}

	return vision, nil
}

func (u *VisionServiceImpl) GetAll() ([]*models.Vision, error) {
	var visions []*models.Vision
	cursor, err := u.visioncollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var vision *models.Vision
		err := cursor.Decode(&vision)
		if err != nil {
			return nil, err
		}
		visions = append(visions, vision)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(visions) == 0 {
		return nil, errors.New("No visions found")
	}
	return visions, nil
}

func (u *VisionServiceImpl) UpdateVision(vision *models.Vision, visionSerialNumber *string) error {

	filter := bson.D{primitive.E{Key: "serial_number", Value: vision.Serial_number}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "serial_number", Value: vision.Serial_number},
		primitive.E{Key: "stage_1", Value: vision.Stage_1},
		primitive.E{Key: "stage_2", Value: vision.Stage_2},
		primitive.E{Key: "stage_3", Value: vision.Stage_3},
		primitive.E{Key: "stage_4", Value: vision.Stage_4},
		primitive.E{Key: "stage_5", Value: vision.Stage_5}}}}
	result, err := u.visioncollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("No visions found for update")
	}
	return nil
}


func (u *VisionServiceImpl) UpdateStageProcessFlags(vision *models.Vision, visionSerialNumber *string) error {
	// หาเอกสารที่ตรงกับเงื่อนไข
	filter := bson.D{primitive.E{Key: "serial_number", Value: *visionSerialNumber}}
	var existingVision models.Vision
	err := u.visioncollection.FindOne(u.ctx, filter).Decode(&existingVision)
	if err != nil {
		return err
	}

	// อัปเดตข้อมูลในเอกสารที่ตรงกับเงื่อนไข
	existingVision.Stage_1.Process_in = vision.Stage_1.Process_in
	existingVision.Stage_1.Process_out = vision.Stage_1.Process_out
	existingVision.Stage_2.Process_in = vision.Stage_2.Process_in
	existingVision.Stage_2.Process_out = vision.Stage_2.Process_out
	existingVision.Stage_3.Process_in = vision.Stage_3.Process_in
	existingVision.Stage_3.Process_out = vision.Stage_3.Process_out
	existingVision.Stage_4.Process_in = vision.Stage_4.Process_in
	existingVision.Stage_4.Process_out = vision.Stage_4.Process_out
	existingVision.Stage_5.Process_in = vision.Stage_5.Process_in
	existingVision.Stage_5.Process_out = vision.Stage_5.Process_out

	// อัปเดตข้อมูลใน MongoDB
	_, err = u.visioncollection.ReplaceOne(u.ctx, filter, existingVision)
	if err != nil {
		return err
	}

	return nil
}


func (u *VisionServiceImpl) DeleteVision(Serial_number *string) error {
	filter := bson.D{primitive.E{Key: "serial_number", Value: Serial_number}}
	result, _ := u.visioncollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("Now visions found for delete")
	}
	return nil
}
