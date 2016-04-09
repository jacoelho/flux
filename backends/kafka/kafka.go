package kafka

import "github.com/Shopify/sarama"

type kafka struct {
	Client sarama.AsyncProducer
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

	return &kafka{Client: producer}, nil
}

func (k *kafka) Serialize(msg string) error {
	msgEncoder := sarama.StringEncoder(msg)

	k.Client.Input() <- &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("test"),
		Value: msgEncoder,
	}
	return nil
}

func (k *kafka) Close() error {
	k.Client.AsyncClose()
	return nil
}
