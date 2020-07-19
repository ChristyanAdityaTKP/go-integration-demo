package app

import (
	"os"

	// PSQL driver
	"github.com/dh258/go-integration-demo/domain"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	router = gin.Default()
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(".env file not found")
	}
}

// Start the application
func Start() {
	setDB()
	routes()
	router.Run(":8080")
}

func setDB() {
	// Fetch DB credentials
	username := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	domain.AddressRepo.Initialize(host, port, username, password, database)
}
