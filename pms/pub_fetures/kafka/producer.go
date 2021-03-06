package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

type Kafka struct {
	AsyncProducer sarama.AsyncProducer
}

const addr = "192.168.10.113:9092"
const TopicShopifyPublish = "shopify-publish"

func (kafka *Kafka) Producer() *Kafka {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	kafkaProducer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		log.Fatal(err)
	}
	kafka.AsyncProducer = kafkaProducer
	return kafka
}

func (kafka *Kafka) Produce(topic string, messages []string) {
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt)

	var enqueued int

	for _, message := range messages {
		// 生产消息
		produceMessage := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		//select {
		kafka.AsyncProducer.Input() <- produceMessage
		enqueued++
		//case <-signals:
		// 命令行接受断开信号
	}
}

func (kafka *Kafka) Close() {
	kafka.AsyncProducer.AsyncClose()
}

func (kafka *Kafka) WaitProduce() {
	var (
		wg              sync.WaitGroup
		success, errors int
	)

	wg.Add(2)

	log.Printf("Prepare produce: %d; errors: %d\n", success, errors)

	// goroutine 发送成功
	go func() {
		defer wg.Done()
		for range kafka.AsyncProducer.Successes() {
			success++
		}
	}()

	// goroutine 发送失败
	go func() {
		defer wg.Done()
		for range kafka.AsyncProducer.Errors() {
			errors++
		}
	}()

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", success, errors)
}
