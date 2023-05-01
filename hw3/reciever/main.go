package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"strconv"

	color "github.com/fatih/color"
	kafka "github.com/segmentio/kafka-go"
)

type message struct {
	Type string `json:"type"`
}

var (
	types   = [2]string{"message", "error"}
	brokers = []string{"kafka0:9092", "kafka1:9093", "kafka2:9094"}
)

const (
	readTopic  = "main"
	writeTopic = "DeadLetter"
)

func main() {
	color.NoColor = false

	reader := setupKafkaReader(brokers, readTopic)
	writer, err := setupKafkaWriter(brokers, writeTopic)
	if err != nil {
		log.Fatalln(color.RedString("Can't setup kafka: %v", err))
	}

	log.Println("Start kafka consumer!")
	for {
		msg, err := readMessage(reader)
		if err != nil {
			log.Println(color.RedString("Failed to read message: %v", err))
			continue
		}

		log.Printf("Succesfuly read message of type: %v", msg.Type)

		if msg.Type == types[0] {
			log.Println(color.GreenString("Done!"))

		} else {
			if err = writeMessage(writer, msg); err != nil {
				log.Println(color.RedString("Failed to send message: %v", err))
			}
			log.Printf("Succesfuly send message of type %v to %v topic", msg.Type, writeTopic)
		}
	}

}

func readMessage(reader *kafka.Reader) (message, error) {
	msgEncoded, err := reader.ReadMessage(context.Background())
	if err != nil {
		return message{}, fmt.Errorf("can't read message from kafka: %v", err)
	}
	var msg message
	if err = json.Unmarshal(msgEncoded.Value, &msg); err != nil {
		return message{}, fmt.Errorf("can't unmarshal message: %v", err)
	}
	return msg, err
}

func writeMessage(writer *kafka.Writer, msg message) error {

	msgEncoded, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("can't marshal message: %v", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: msgEncoded,
	})
	if err != nil {
		return fmt.Errorf("can't send msg to kafka: %v", err)
	}
	return nil
}

func setupKafkaWriter(brokers []string, topic string) (*kafka.Writer, error) {
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

func setupKafkaReader(brokers []string, topic string) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:   topic,
		Brokers: brokers,
	})
	return reader
}
