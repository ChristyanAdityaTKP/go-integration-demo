package app

import "github.com/dh258/go-integration-demo/controllers"

func routes() {
	router.GET("/", controllers.Healthcheck)
	router.POST("/addresses", controllers.CreateAddress)
	router.GET("/addresses", controllers.GetAllAddresses)
	router.GET("/addresses/:id", controllers.GetAddressByID)
}
