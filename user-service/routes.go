package main

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

var userTopic = "user-topic"

// SetUserRoutes, user servisi için HTTP rotalarını tanımlar.
func SetUserRoutes(router *mux.Router) {
    router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        // Kullanıcıları getirme işlemleri burada yapılır.
        // Bu fonksiyon, gelen HTTP isteğine yanıt olarak çalışır.
        // Örnek olarak, w.Write() ile JSON yanıtı gönderebilirsiniz.

        // Kullanıcıları getirdikten sonra Kafka'ya mesaj gönderme örneği:
        sendMessage("user-topic", "New user created!")

        fmt.Fprintln(w, "GetUser Endpoint - Placeholder")
    }).Methods("GET")
}

func sendMessage(topic, message string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}
	defer producer.Close()

	// Mesajı Kafka'ya gönder
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
	}
}
