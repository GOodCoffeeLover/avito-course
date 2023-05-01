package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"

	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type message struct {
	Type string `json:"type"`
}

var types = [2]string{"message", "error"}

func main() {
	writer, err := setupKafka([]string{"kafka0:9092", "kafka1:9093", "kafka2:9094"}, "main")
	if err != nil {
		log.Fatalf("Can't setup kafka: %v", err)
	}
	for {
		i := rand.Int63n(2)
		msg, _ := json.Marshal(message{
			Type: types[i],
		})
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: msg,
		})
		if err != nil {
			log.Printf("Can't send msg to kafka: %v", err)
		}
		log.Printf("Successfuly sent message  with type %v", types[i])
		time.Sleep(time.Millisecond * 500) // 0.5 s
	}

}

func setupKafka(brokers []string, topic string) (*kafka.Writer, error) {
	for _, broker := range brokers {
		if err := createTopic(topic, broker); err != nil {
			return nil, err
		}
	}

	writer := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: topic,
	}
	return writer, nil
}

func createTopic(topic, broker string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return fmt.Errorf("can't connect to kafka due to: %v", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("can't connect to kafka controller due to: %v", err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("can't connect to kafka controller due to: %v", err)
	}
	defer controllerConn.Close()

	topicConfigs := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 3,
	}

	err = controllerConn.CreateTopics(topicConfigs)
	if err != nil {
		return fmt.Errorf("can't create topic %v in kafka(%v) due to: %v", topic, broker, err)
	}
	return nil
}
