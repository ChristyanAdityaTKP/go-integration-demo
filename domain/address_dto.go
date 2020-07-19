package domain

import (
	"strings"
	"time"

	"github.com/dh258/go-integration-demo/utils"
)

type Address struct {
	ID        int64     `json:"id" xorm:"id pk autoincr"`
	Name      string    `json:"name" xorm:"name"`
	Country   string    `json:"country" xorm:"country"`
	CreatedAt time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated_at"`
}

func (a *Address) Validate() utils.MessageErr {
	a.Name = strings.TrimSpace(a.Name)
	a.Country = strings.TrimSpace(a.Country)

	if a.Name == "" {
		return utils.NewUnprocessibleEntityError("Please enter a valid name")
	}
	if a.Country == "" {
		return utils.NewUnprocessibleEntityError("Please enter a valid country")
	}
	return nil
}
