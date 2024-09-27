package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (t *TreeProducer) ReportEvents(p *kafka.Producer) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				t.failingToDeliver = true
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				t.failingToDeliver = false
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}
