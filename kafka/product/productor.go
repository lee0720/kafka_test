package product

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

// var enqueued, producerErrors int

// Productor ...
func Productor(topic, id string, obj interface{}, producer sarama.AsyncProducer) {

	log.Println("productor")

	meg, err := json.Marshal(obj)
	if err != nil {
		log.Println("[ERROR] ", err)
		return
	}

	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(id), Value: sarama.StringEncoder(string(meg))}

	select {
	case err := <-producer.Errors():
		log.Println(err.Err)
	default:
		log.Println("Successfully")
	}
}
