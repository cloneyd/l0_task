package stanconn

import (
	"github.com/nats-io/stan.go"
	"log"
)

const (
	clusterID = "test-cluster"
	clientID  = "l0-internship-nats-service"
)

type Conn struct {
	conn stan.Conn
}

func NewConn() stan.Conn {
	conn, err := stan.Connect(clusterID, clientID, stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
		log.Fatalf("Connection lost, reason: %v", reason)
	}), stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func (c Conn) Subscribe(subject string, messageHandler stan.MsgHandler) (stan.Subscription, error) {
	return c.conn.Subscribe(subject, messageHandler)
}
