package main

import (
	"context"
	"fmt"
	"log"
	
	"os"

	controllers "golang/battery-tracking/controller"
	
	"golang/battery-tracking/routes"
	"golang/battery-tracking/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	server  	*gin.Engine
	us 			services.VisionService
	uc      	controllers.VisionController
	
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
	uc    = controllers.New(us)
	server= gin.Default()
}

func main() {

// Load .env file
err := godotenv.Load(".env")
if err != nil {
	log.Fatal("Error loading .env file")
}

// Create a new Gin engine
router := gin.New()
router.Use(gin.Logger())

// Register Vision routes on the chosen engine
basepath := router.Group("/v1")
uc.RegisterVisionRoutes(basepath)

// Register Auth and User routes on the same engine
routes.AuthRoutes(router)
routes.UserRoutes(router)

// Additional routes if needed
router.GET("/api-1", func(c *gin.Context){
	c.JSON(200, gin.H{"seccess":"Access granted for api-1"})
})

router.GET("/api-2", func(c *gin.Context){
	c.JSON(200, gin.H{"success":"Access granted for api-2"})
})

// Get port from .env (default to 8080 if not set)
port := os.Getenv("PORT")
if port == "" {
	port = "9000"
}

// Run the consolidated server on the desired port
err = router.Run(":" + port)
if err != nil {
	log.Fatalf("Failed to run server on port %s: %v", port, err)
}

// Clean up resources when done
defer func() {
	err := mongoclient.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}()

}

func DBinstance() {
	panic("unimplemented")
}
