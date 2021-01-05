package consumer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

// Consumer ...
func Consumer(topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewClient([]string{"192.168.89.71:9092", "192.168.89.71:9093", "192.168.89.71:9094"}, config)
	if err != nil {
		log.Println("--------------------->")
		panic(err)
	}
	defer client.Close()
	consumer, err := sarama.NewConsumerFromClient(client)

	defer consumer.Close()
	if err != nil {
		panic(err)
	}

	// get partitionId list
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for i, partitionId := range partitions {
		// create partitionConsumer for every partitionId
		partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		log.Printf("Topic: %v, Partition: %v, consumer running ...", topic, i)
		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			// block
			for message := range (*pc).Messages() {
				value := string(message.Value)
				log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
			}
		}(&partitionConsumer)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}
