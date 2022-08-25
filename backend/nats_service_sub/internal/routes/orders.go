package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"nats_service/internal/cache"
	"nats_service/internal/models"
	"net/http"
)

func AddOrderRoutes(rg *gin.RouterGroup, cache cache.Cache) {
	orders := rg.Group("/orders")

	orders.GET("/:uid", func(c *gin.Context) {
		var order models.Order
		uid := c.Param("uid")

		orderData, err := cache.Get(uid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no orders with this uid",
			})
			return
		}

		if err := json.Unmarshal(orderData, &order); err != nil {
			log.Println("unable to unmarshall to JSON", orderData)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "unable to unmarshall to JSON",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"order": order,
		})
	})
}
