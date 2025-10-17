package main

import (
	"core/controller"
	"core/db"
	"core/generated"
	"core/logger"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := logger.Init()
	if err != nil {
		panic(err)
	}
	log := logger.Logger
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	//connectDB()
	if err := db.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to the database: ", zap.Error(err))
	}
	//MigrateSchema()
	db.MigrateSchema()

	// elasticsearch.Init()

	router := DataManagementServerRouter()

	router.Run(":" + os.Getenv("SERVER_PORT"))

}

func DataManagementServerRouter() *gin.Engine {

	routes := generated.ApiHandleFunctions{}
	routes.MigrationAPIAPI = controller.NewMigrationApiController()
	routes.SearchAPIAPI = controller.NewSearchApiController()

	router := gin.Default()

	router = generated.NewRouterWithGinEngine(router, routes)

	router.Static("/swagger-ui", "./swagger-ui")

	return router
}
