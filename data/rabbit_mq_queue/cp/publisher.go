package cp

import (
	mqLog "data/rabbit_mq_queue/cp/mqL"
	"fmt"
	"github.com/streadway/amqp"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Publish() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		mqLog.Log("conn", err)
	}
	defer conn.Close()

	netAddr := conn.LocalAddr()
	fmt.Println("Publish: ", netAddr.String())

	ch, err := conn.Channel()
	if err != nil {
		mqLog.Log("channel", err)
	}
	// 失败关闭，优雅关闭，释放资源
	defer ch.Close()

	// 指明发送队列
	// auto delete ,发完就删除
	queue, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil {
		mqLog.Log("declare queue", err)
	}

	rand.Seed(time.Millisecond.Microseconds())
	id := rand.Int63()

	// 定义消息体
	publishMessage := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("userID: " + strconv.FormatInt(id, 10) + "; This is first rabbitmq message from hello queue. " + time.Now().Format(time.RFC3339)),
	}

	signalCh := make(chan os.Signal, 1)
	doneCh := make(chan bool, 1)
	doneCh <- true
	go signalOut(signalCh)
	for <-doneCh {
		select {
		case <-signalCh:
			defer func() {
				fmt.Println("ch close, conn close")
			}()
			fmt.Println("Publish signal out Crtl + C")
			doneCh <- false
			close(signalCh)
		default:
			id = rand.Int63()
			publishMessage.Body =
				[]byte("userID " + strconv.FormatInt(id, 10) + ";This is first rabbitmq message from hello queue. " + time.Now().Format(time.RFC3339))

			err = ch.Publish("", queue.Name, false, false, publishMessage)

			if err != nil {
				mqLog.Log("publish", err)
			}
			doneCh <- true
			time.Sleep(time.Millisecond * 500)
		}
	}

	fmt.Println("Publish end")

	go func() {
		var notifyReturnCh chan amqp.Return = make(chan amqp.Return)
		ch.NotifyReturn(notifyReturnCh)
		notifyReturnMessage := <-notifyReturnCh
		fmt.Println("notify Return")

		fmt.Println(notifyReturnMessage.RoutingKey)
		fmt.Println(notifyReturnMessage.Exchange)
		fmt.Println(notifyReturnMessage.ContentType)
		fmt.Println(notifyReturnMessage.Body)
	}()

}

// 注册信号，退出
func signalOut(signalCh chan os.Signal) {
	signal.Notify(signalCh, syscall.SIGINT)
}
