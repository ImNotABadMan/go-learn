package cp

import (
	mqLog "data/rabbit_mq_queue/cp/mqL"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func Customer() {
	url := "amqp://guest:guest@127.0.0.1:5672"
	fmt.Println("Start Receive", url)
	// 连接rabbitmq
	conn, err := amqp.Dial(url)
	defer conn.Close()
	if err != nil {
		mqLog.Log("conn", err)
	}

	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		mqLog.Log("channel", err)
	}

	// 声明队列
	queue, err := ch.QueueDeclare("hello", true, true, false, false, nil)
	if err != nil {
		mqLog.Log("queue", err)
	}

	// 接受消息
	deliveryCh, err :=
		ch.Consume(queue.Name, "", false, false, false, false, nil)

	//sigCh := make(chan os.Signal, 1)
	//go signalOut(sigCh)

	for delivery := range deliveryCh {
		fmt.Println("Receive: ")
		fmt.Println(delivery.Headers)
		fmt.Println(delivery.RoutingKey)
		fmt.Println(string(delivery.Body))
		if err := delivery.Ack(false); err != nil {
			//mqLog.Log("ack", err)
			log.Println("ack", err)
		}
		time.Sleep(time.Millisecond * 900)
	}

	fmt.Println("Customer end")

}
