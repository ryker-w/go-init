package process

import (
	"encoding/json"

	proxy "github.com/eclipse/paho.mqtt.golang"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
)

func OnConnectHandler(_ proxy.Client) {
	log.Info("[OnConnectHandler] Connect Mqtt success")
	var qos byte = 1
	topic := "event/+/up"
	err := app.GetMqtt().Subscribe(HandleMessage, qos, topic) // 订阅主题
	if err != nil {
		log.Info(err)
		return
	}
}

func OnLostHandler(_ proxy.Client, err error) {
	log.Info("[OnLostHandler] disconnect Mqtt, err: %s", err.Error())
}

// HandleMessage 处理上报数据
func HandleMessage(topic string, payload []byte) {
	log.Info("bridge receive : %s", string(payload))
	defer func() {
		if err := recover(); err != nil {
			log.Info(err)
		}
	}()
	// 数据格式转换
	var req map[string]interface{}
	err := json.Unmarshal(payload, &req)
	if err != nil {
		log.Info(err)
		log.Info("json.Unmarshal error: topic=%s,  payload=%s", topic, string(payload))
		return
	}
	// TODO handler

	// 如：Amqp 转发上报
	// var m = rabbit.Message{
	// 	Payload: req,
	// 	Router: rabbit.Route{
	// 		Exchange: rabbit.DefaultExchange,
	// 		Key:      "topic_B",
	// 	},
	// }
	// m.SetOption(rabbit.JsonEncodeOption, rabbit.UUIDMsgIdOption)
	// err = app.GetAmqp().Publish(m)
	// if err != nil {
	// 	log.Info(err)
	// 	return
	// }
}
