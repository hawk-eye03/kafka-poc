package consumers

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hawk-eye03/kafka-poc/lib/config"
	"go.uber.org/zap"
)

type BaseConsumer struct {
	C    Consumer
	conf config.ConfigMap
}

func NewBaseConsumer(consumer Consumer, conf config.ConfigMap) *BaseConsumer {
	return &BaseConsumer{
		C:    consumer,
		conf: conf,
	}
}

func (b *BaseConsumer) StartConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": b.conf.Kafka.Host,    // Kafka broker address
		"group.id":          b.conf.Kafka.GroupID, // Consumer group ID
		"auto.offset.reset": "earliest",           // Start consuming from the beginning of the topic
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		zap.L().Info("Error failed in creating a new consumer %v\n", zap.Error(err))
		return
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err = consumer.SubscribeTopics(b.C.GetTopicName(), nil)
	if err != nil {
		zap.L().Info("Failed to subscribe to topic: %v\n", zap.Error(err))
		return
	}

	go func(c *kafka.Consumer) {
		defer c.Close()

		for {
			select {
			case sig := <-sigchan:
				zap.L().Info("Caught signal %v: terminating\n", zap.String("signal", sig.String()))
				return
			default:
				ev := c.Poll(500)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					zap.L().Info("Consumer received message: %s\n", zap.String("Message", string(e.Value)))
					b.C.ProcessMessage(string(e.Value))
					// Handle the message here
				case kafka.Error:
					zap.L().Info("Consumer error: %v\n", zap.String("Message", string(e.Code().String())))
				}
			}
		}
	}(consumer)
	zap.L().Info("Consumers started. Press Ctrl+C to exit.")
	<-sigchan
	zap.L().Info("Exiting Consumer.")
}
