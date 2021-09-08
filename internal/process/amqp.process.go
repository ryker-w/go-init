package process

import (
	"context"
	"github.com/lishimeng/go-log"
	mq "github.com/ryker-w/go-init/internal/amqp"
	"github.com/ryker-w/go-init/internal/etc"
)

var MqService mq.Service

type Ds struct {
}

func (ds *Ds) Subscribe(topic string, v interface{}, upstream mq.Upstream) {
	log.Info("ds receive:" + topic)
	log.Info("data:")
	s := v.([]byte)
	log.Info(string(s))
}

func (ds *Ds) Topic() string {
	return "MQTopicLoraToInflux"
}

// AmqpStart 启动-运行amqp
func AmqpStart(ctx context.Context) (err error) {

	ds := Ds{}
	MqService, err = mq.Start(ctx, mq.Connector{Conn: etc.Config.Mq.MQConn}, &ds)
	if err != nil {
		log.Info("Amqp start err ")
		log.Info(err)
		return
	}

	return
}
