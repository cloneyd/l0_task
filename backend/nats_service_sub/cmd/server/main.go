package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"nats_service/internal/cache"
	"nats_service/internal/database/postgres"
	"nats_service/internal/models"
	"nats_service/internal/routes"
	"nats_service/internal/stanconn"
	"os"
	"os/signal"
)

func main() {
	conn := stanconn.NewConn()
	storage := cache.NewCache()

	db := postgres.NewConn()
	defer db.Close()

	storage.InitFromDB(db)

	sub, err := conn.Subscribe("foo", func(msg *stan.Msg) {
		var order models.Order

		orderUid := order.OrderUid
		jsonData := msg.Data

		if err := json.Unmarshal(jsonData, &order); err != nil {
			log.Println(msg.Data)
			log.Printf("Error unmarshalling incoming order JSON: %s\n", err)
			return
		}

		db.Exec(context.Background(), "CALL insert_order($1, $2)", orderUid, jsonData)

		if err := storage.Set(order.OrderUid, msg.Data); err != nil {
			log.Printf("Error setting cache item: %s\n", err)
			return
		}
	})
	if err != nil {
		log.Fatalln(err)
	}

	go routes.Run(storage)

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			sub.Unsubscribe()

			conn.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
