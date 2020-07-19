package controllers

import (
	"net/http"

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
