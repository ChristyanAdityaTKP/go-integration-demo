package domain

import (
	"fmt"

	"github.com/dh258/go-integration-demo/utils"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

var (
	AddressRepo addressRepoInterface = &addressRepo{}
)

type addressRepoInterface interface {
	Initialize(host, port, username, password, database string) *xorm.Engine
	Create(*Address) (*Address, utils.MessageErr)
}

type addressRepo struct {
	db *xorm.Engine
}

func NewAddressRepository(db *xorm.Engine) addressRepoInterface {
	return &addressRepo{db: db}
}

func (ar *addressRepo) Initialize(host, port, username, password, database string) *xorm.Engine {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	// Setup DB
	ar.db, err = xorm.NewEngine("postgres", dsn)
	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}

	return ar.db
}

func (ar *addressRepo) Create(address *Address) (*Address, utils.MessageErr) {
	affected, err := ar.db.Table("address").Insert(address)
	if err != nil {
		log.Err(err).Msgf("Create address failed: %s", err)
		return nil, utils.NewInternalServerError("Error when trying to save address")
	}
	log.Printf("affected row: %s\n", affected)

	return address, nil
}
