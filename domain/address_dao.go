package domain

import (
	"fmt"

	// Postgres driver
	_ "github.com/lib/pq"

	"github.com/dh258/go-integration-demo/utils"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

var (
	AddressRepo addressRepoInterface = &addressRepo{}
)

const (
	addressTable = "address"
)

type addressRepoInterface interface {
	Initialize(host, port, username, password, database string) *xorm.Engine
	Create(*Address) (*Address, utils.MessageErr)
	Get(addressID int64) (*Address, utils.MessageErr)
	GetAll() ([]*Address, utils.MessageErr)
}

type addressRepo struct {
	db *xorm.Engine
}

func NewAddressRepository(db *xorm.Engine) addressRepoInterface {
	return &addressRepo{db: db}
}

func (ar *addressRepo) Initialize(host, port, username, password, database string) *xorm.Engine {
	var err error

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	log.Printf("dsn: %s", dsn)
	// Setup DB
	ar.db, err = xorm.NewEngine("postgres", dsn)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	return ar.db
}

func (ar *addressRepo) Create(address *Address) (*Address, utils.MessageErr) {
	_, err := ar.db.Table(addressTable).Insert(address)
	if err != nil {
		log.Err(err).Msg("Create address failed")
		return nil, utils.NewInternalServerError("Error when trying to save address")
	}

	return address, nil
}

func (ar *addressRepo) Get(addressID int64) (*Address, utils.MessageErr) {
	result := &Address{}

	log.Printf("address ID: %v", addressID)
	ar.db.ShowSQL(true)
	_, err := ar.db.Table(addressTable).ID(addressID).Get(result)
	if err != nil {
		log.Err(err).Msg("Get address failed")
		return nil, utils.NewInternalServerError("Error fetching address")
	}

	return result, nil
}

func (ar *addressRepo) GetAll() ([]*Address, utils.MessageErr) {
	var result []*Address

	err := ar.db.Table(addressTable).Find(&result)
	if err != nil {
		log.Err(err).Msg("Get all address failed")
		return nil, utils.NewInternalServerError("Error fetching addresses")
	}

	return result, nil
}
