package integrationtest

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/dh258/go-integration-demo/domain"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

const (
	addressTable              = "address"
	queryTruncateAddressTable = "TRUNCATE TABLE address;"
)

var (
	db *xorm.Engine
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../.env.test"))
	if err != nil {
		log.Fatal("Failed to get env file.")
	}
	os.Exit(m.Run())
}

func setupDB() {
	// Fetch DB credentials
	username := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	db = domain.AddressRepo.Initialize(host, port, username, password, database)
}

func clearTable() error {
	_, err := db.Exec(queryTruncateAddressTable)
	if err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}

	return nil
}

func seedOneAddress() (*domain.Address, error) {
	address := &domain.Address{
		Name:      "Jalan Purwokerto",
		Country:   "Indonesia",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := db.Table(addressTable).Insert(&address)
	if err != nil {
		log.Fatalf("Failed to insert table: %v", err)
	}

	return address, nil
}

func seedAddresses() ([]*domain.Address, error) {
	addresses := []*domain.Address{
		{
			Name:      "Jalan Purwokerto",
			Country:   "Indonesia",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jalan Tanjung Duren",
			Country:   "Malaysia",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jalan Melayang",
			Country:   "Singapura",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, address := range addresses {
		_, err := db.Table(addressTable).Insert(&address)
		if err != nil {
			log.Fatalf("Failed to insert table: %v", err)
		}

	}

	return addresses, nil
}
