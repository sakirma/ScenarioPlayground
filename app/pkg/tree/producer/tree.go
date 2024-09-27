package producer

import (
	"ScenarioPlayground/contract"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

type TreeProducerConfig struct {
	interval time.Duration
}

type TreeProducer struct {
	producer *kafka.Producer

	TreeProducerConfig

	failingToDeliver bool
}

func New(config TreeProducerConfig) (contract.TreeProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	treeProducer := &TreeProducer{
		producer:           p,
		TreeProducerConfig: config,
	}

	go treeProducer.ReportEvents(p)

	return treeProducer, nil
}

func (t *TreeProducer) ProduceLeaf() error {
	go func() {
		interval := t.TreeProducerConfig.interval
		topic := "Tree"

		timer := time.NewTicker(interval)
		defer timer.Stop()

		for range timer.C {
			err := t.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte("leaf"),
			}, nil)

			if err != nil {
				fmt.Printf("Failed to produce message: %v\n", err)
			}
		}
	}()
}

func (t *TreeProducer) ShutDown() {
	p := t.producer

	// Flush and close the producer and the events channel
	for p.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}
}

func (t *TreeProducer) Healthy() bool {
	return true
}

func (t *TreeProducer) Ready() bool {
	return true
}
