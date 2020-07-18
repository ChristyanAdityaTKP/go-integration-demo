package app

import "github.com/dh258/go-integration-demo/controllers"

func routes() {
	router.GET("/", controllers.Healthcheck)
}
