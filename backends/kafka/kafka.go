package kafka

import (
	"bytes"

	"github.com/Shopify/sarama"
)

type kafka struct {
	Client sarama.AsyncProducer
	buffer *bytes.Buffer
}

func New(brokers []string) (*kafka, error) {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 5
	// The level of acknowledgement reliability needed from the broker.
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &kafka{
		Client: producer,
		buffer: new(bytes.Buffer),
	}, nil
}

func (k *kafka) Write(p []byte) (n int, err error) {
	return k.buffer.Write(p)
}

func (k *kafka) Flush() error {
	k.Client.Input() <- &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("test"),
		Value: sarama.StringEncoder(k.buffer.String()),
	}

	k.buffer.Reset()
	return nil
}

func (k *kafka) Close() error {
	k.Client.AsyncClose()
	return nil
}
