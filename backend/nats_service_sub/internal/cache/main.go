package cache

import (
	"encoding/json"
	"log"
	"nats_service/internal/database/postgres"
	"nats_service/internal/models"
	"time"

	"github.com/allegro/bigcache/v3"
)

const (
	eviction = time.Minute * 10
)

type Cache struct {
	cache *bigcache.BigCache
}

func NewCache() Cache {
	cache, err := bigcache.NewBigCache(
		bigcache.DefaultConfig(eviction),
	)
	if err != nil {
		log.Fatalln(err)
	}

	return Cache{cache: cache}
}

func (c *Cache) InitFromDB(db *postgres.Database) {
	orders, err := models.GetAllOrders(db)
	if err != nil {
		log.Println("Error getting all orders:", err)
		return
	}

	for _, order := range orders {
		orderByte, err := json.Marshal(order)
		if err != nil {
			log.Println("Error marshalling struct:", err)
			return
		}
		err = c.Set(order.OrderUid, orderByte)
		if err != nil {
			log.Printf("Error setting cache item with order UID %s: %s\n", order.OrderUid, err)
			return
		}
	}
}

func (c *Cache) Set(key string, entry []byte) error {
	return c.cache.Set(key, entry)
}

func (c *Cache) Get(key string) ([]byte, error) {
	return c.cache.Get(key)
}
