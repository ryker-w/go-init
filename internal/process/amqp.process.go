package process

import (
	"encoding/json"

	"github.com/lishimeng/app-starter/amqp/rabbit"
	"github.com/lishimeng/go-log"
)

type AmqpConsumer struct {
}

func (c *AmqpConsumer) Subscribe(payload interface{}, txHandler rabbit.TxHandler, serverCtx rabbit.ServerContext) {
	log.Info("ds receive from:" + serverCtx.Queue)
	s := payload.([]byte)
	log.Info("data:%s", string(s))

	// Unmarshal struct
	var req map[string]interface{}
	err := json.Unmarshal(s, &req)
	if err != nil {
		log.Info("Unmarshal req err")
		log.Info(err)
		return
	}
	//TODO func handler
}

func (c *AmqpConsumer) Router() rabbit.Route {
	return rabbit.Route{
		Queue: "topic_A",
		Key:   "topic_A_1",
	}
}
