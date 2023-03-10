package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xcodeme21/go-test-project/controllers"
	"github.com/xcodeme21/go-test-project/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        c.Writer.Status(),
		"data":          nil,
		"error_message": "Lost in space",
	})
}

func createDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME_1"))
	DB.Exec(createDatabaseCommand)

	createDatabaseCommand2 := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME_2"))
	DB.Exec(createDatabaseCommand2)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Yayyyy I'am Gin Gonic",
		})
	})

	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE", "PATCH"}
	r.Use(cors.New(config))

	createDatabase()

	// Initializes databaseSource
	db, _ := database.Initialize()
	r.Use(database.Inject(db))

	// Initializes database Destination
	dbTwo, _ := database.InitializeTwo()
	r.Use(database.Inject(dbTwo))

	//Seeder
	database.SourceSeeder()
	database.DestinationSeeder()

	//Connection
	cn, _ := database.Connect()
	cnTwo, _ := database.ConnectTwo()

	sourceController := controllers.ProductsController{DB: cn}
	r.GET("/list-source-products", sourceController.ListSourceProduct)

	destinationController := controllers.ProductsController{DB: cnTwo}
	r.GET("/list-destination-products", destinationController.ListDestinationProduct)

	controller := controllers.ProductsController{DB: cn, DBTwo: cnTwo}
	r.GET("/update-destination-products", controller.UpdateDestinationProduct)

	port := os.Getenv("PORT")
	r.NoRoute(lostInSpce)
	r.Run(":" + port)
}
