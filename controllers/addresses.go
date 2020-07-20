package controllers

import (
	"net/http"
	"strconv"

	"github.com/dh258/go-integration-demo/domain"
	"github.com/dh258/go-integration-demo/usecase"
	"github.com/dh258/go-integration-demo/utils"
	"github.com/gin-gonic/gin"
)

// CreateAddress for creating new address
func CreateAddress(c *gin.Context) {
	var address domain.Address

	err := c.ShouldBindJSON(&address)
	if err != nil {
		errResp := utils.NewUnprocessibleEntityError("JSON cannot be processed")
		c.JSON(errResp.Status(), errResp)
		return
	}

	result, errResult := usecase.AddressUsecase.CreateAddress(&address)
	if errResult != nil {
		c.JSON(errResult.Status(), errResult)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAllAddresses fetches all addresses in the database
func GetAllAddresses(c *gin.Context) {
	result, err := usecase.AddressUsecase.GetAllAddresses()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAddressByID fetches address by its ID
func GetAddressByID(c *gin.Context) {
	id := c.Param("id")
	sanitizedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errMsg := utils.NewBadRequestError("Address ID should be a number")
		c.JSON(errMsg.Status(), errMsg)
		return
	}

	address, getErr := usecase.AddressUsecase.GetByID(sanitizedID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, address)
}
