package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Bright2704/battery-tracking/controller"
	"github.com/Bright2704/battery-tracking/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	server  	*gin.Engine
	us 			services.VisionService
	uc      	controller.VisionController
	ctx     	context.Context
	visionc 	*mongo.Collection
	mongoclient *mongo.Client
	err     	error
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://username:password@119.59.99.194:27017").SetAuth(options.Credential{
		AuthSource: "battery",
		Username:   "dev",
		Password:   "123456789",
	})
	// Auto Migrate the struct
	
	

	// Connect to MongoDB server.
	mongoclient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err = mongoclient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB server.")

	visionc = mongoclient.Database("battery").Collection("Battery-Tracking")
	us    = services.NewVisionService(visionc, ctx)
	uc    = controller.New(us)
	server= gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterVisionRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}