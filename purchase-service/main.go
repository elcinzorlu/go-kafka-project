package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

var purchaseTopic = "purchase-topic"

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
	// Satın alma işlemlerini yönetme işlemini gerçekleştirebilirsiniz.

	// Kafka'ya mesaj gönder
	sendMessage(purchaseTopic, "Purchase Service: Purchase Endpoint")

	// Örnek bir çıktı:
	fmt.Fprintf(w, "Purchase Service: Purchase Endpoint")
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

func main() {
	// Gorilla Mux router'ını oluşturun.
	router := mux.NewRouter()

	// SetPurchaseRoutes fonksiyonunu kullanarak rotaları ayarlayın.
	SetPurchaseRoutes(router)

	// HTTP sunucusunu başlatın.
	port := 8081
	fmt.Printf("Purchase Service is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

