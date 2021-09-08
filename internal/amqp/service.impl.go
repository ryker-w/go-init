package mq

import (
	"context"
	"github.com/streadway/amqp"
	"log"
)

type amqpService struct {
	ctx         context.Context
	upstream    *AmqpUpstream
	downStreams []Downstream

	conn *amqp.Connection
	chs  []*amqp.Channel

	amqpAddr string
}

func (s *amqpService) Run() (err error) {
	err = s.init()
	if err != nil {
		return
	}
	go s.run()
	return
}

func (s *amqpService) UpstreamHandler() Upstream {
	return s.upstream
}

func (s *amqpService) initUpstream() (err error) {
	log.Println("init upstream")
	ch, err := s.newCh()
	if err != nil {
		log.Println("upstream failed")
		log.Println(err)
		return
	}

	s.upstream = &AmqpUpstream{
		ch:       ch,
		exchange: defaultExchange, // TODO
	}
	return
}

func (s *amqpService) newCh() (ch *amqp.Channel, err error) {

	ch, err = s.conn.Channel()
	if err == nil {
		s.chs = append(s.chs, ch)
	} else {
		log.Println("create ch failed")
		log.Println(err)
	}
	return
}

func (s *amqpService) init() (err error) {
	log.Println("init service")
	log.Println("connect mq broker")
	conn, err := amqp.Dial(s.amqpAddr)
	if err != nil {
		log.Println("connect failed")
		log.Println(err)
		return
	}
	go s.releaseConn()

	s.conn = conn
	// init upstream handler
	err = s.initUpstream()
	if err != nil {
		return
	}

	// init downstream handler
	if len(s.downStreams) > 0 {
		for _, ds := range s.downStreams {
			go s.handleSub(ds)
		}
	}
	return
}

func (s *amqpService) releaseConn() {
	for {
		select {
		case <-s.ctx.Done():
			log.Println("conn down")
			for _, ch := range s.chs {
				_ = ch.Close()
			}
			_ = s.conn.Close()
			return
		}
	}
}

func (s *amqpService) handleSub(ds Downstream) {
	if ds == nil {
		return
	}
	ch, err := s.newCh()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("declare queue:" + ds.Topic())
	q, err := ch.QueueDeclare(
		ds.Topic(), // name
		true,       // durable
		true,       // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("queue failed")
		log.Println(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("consume failed")
		log.Println(err)
		return
	}
	err = ch.QueueBind(q.Name, q.Name, defaultExchange, false, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		select {
		case <-s.ctx.Done():
			return
		case m := <-msgs:
			s.handleMsg(ds, m)
		}
	}
}

func (s *amqpService) handleMsg(ds Downstream, ms amqp.Delivery) {
	log.Printf("received down stream [exchange:%s router:%s]", ms.Exchange, ms.RoutingKey)
	log.Printf(string(ms.Body))
	ds.Subscribe(ds.Topic(), ms.Body, s.upstream)
}

func (s *amqpService) run() {
	for {
		select {
		case <-s.ctx.Done():
			return
		}
	}
}
