package services

import (
	"context"
	"errors"

	"github.com/Bright2704/battery-tracking/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VisionServiceImpl struct {
	visioncollection *mongo.Collection
	ctx  			 context.Context
}

func NewVisionService(visioncollection *mongo.Collection, ctx context.Context) VisionService {
	return &VisionServiceImpl{
		visioncollection: visioncollection,
		ctx: 			  ctx,
	}
}

func (u *VisionServiceImpl) CreateVision(vision *models.Vision) error {
	_, err := u.visioncollection.InsertOne(u.ctx, vision)
	return err
}

func (u *VisionServiceImpl) GetVision(Related *string) (*models.Vision, error) {
	var vision *models.Vision
	query := bson.D{bson.E{Key: "related", Value: Related}}
	err := u.visioncollection.FindOne(u.ctx, query).Decode(&vision)
	return vision, err
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

func (u *VisionServiceImpl) UpdateVision(vision *models.Vision) error {
	filter := bson.D{primitive.E{Key: "related", Value: vision.Related}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "related", Value: vision.Related},
																		primitive.E{Key: "stage_1", Value: vision.Stage_1}, 
																		primitive.E{Key: "stage_2", Value: vision.Stage_2}, 
																		primitive.E{Key: "stage_3", Value: vision.Stage_3}, 
																		primitive.E{Key: "stage_4", Value: vision.Stage_4}, 
																		primitive.E{Key: "stage_5", Value: vision.Stage_5}}}}
	result, _ := u.visioncollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No visions found for update")
	}
	return nil
}

func (u *VisionServiceImpl) DeleteVision(Related *string) error {
	filter := bson.D{primitive.E{Key: "related", Value: Related}}
	result, _  := u.visioncollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("Now visions found for delete")
	}
	return nil
}