package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"nats_service/internal/cache"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run(storage cache.Cache) {
	getRoutes(storage)
	router.Use(cors.Default())
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes(storage cache.Cache) {
	api := router.Group("/api")
	AddOrderRoutes(api, storage)
}
