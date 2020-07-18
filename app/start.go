package app

import (
	"fmt"
	"os"

	// PSQL driver
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	log.Info().Msgf("DSN string: %s", dsn)
	// Setup DB
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}
	engine.Ping()
}
