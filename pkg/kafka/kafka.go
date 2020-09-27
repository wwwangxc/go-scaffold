package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(string) string
	GetStringSlice(string) []string
	GetInt(string) int
	GetInt32(string) int32
	GetBool(string) bool
}

var (
	ConsumerGroupErrClosed = errors.New("[kafka] the consumer group is closed")
)

var (
	ConsumerErrClosed = errors.New("[kafka] the consumer is closed")
)

var (
	ProducerErrClosed = errors.New("[kafka] the producer is closed")
)

type (
	// ConsumerGroupHandlerSetup is run at the beginning of a new session, before ConsumeClaim.
	ConsumerGroupHandlerSetup func(sarama.ConsumerGroupSession)

	// ConsumerGroupHandlerConsumerClaim triggered every time when a message is received.
	ConsumerGroupHandlerConsumeClaim func(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim)

	// ConsumerGroupHandlerCleanup is run at the end of a session, once all ConsumeClaim goroutines have exited.
	ConsumerGroupHandlerCleanup func(sarama.ConsumerGroupSession)

	// ConsumerGroupHandlerError is run at the occurs of an error, when consume joins a cluster of consumers for a given list of topics
	ConsumerGroupHandlerError func(error)
)

type (
	// ConsumerHandlerError triggered every time when a error message is reveived.
	ConsumerHandlerError func(*sarama.ConsumerError)

	// ConsumerHandlerSuccess triggered every time when a success message is received.
	ConsumerHandlerSuccess func(*sarama.ConsumerMessage)
)

type (
	// ProducerHandlerError triggered every time when a error message is reveived.
	ProducerHandlerError func(*sarama.ProducerError)
)
