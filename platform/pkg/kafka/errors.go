package platform_kafka

import "errors"

type KafkaError error

var (
	ErrConsumerMessageHandler KafkaError = errors.New("kafka consumer message handler error")
	ErrAuthorization          KafkaError = errors.New("kafka authorization error")
	ErrNetworkError           KafkaError = errors.New("kafka network error")
	ErrCoordination           KafkaError = errors.New("kafka coordination error")
	ErrPartition              KafkaError = errors.New("kafka partition error")
	ErrConfiguration          KafkaError = errors.New("kafka configuration error")
	ErrOffsets                KafkaError = errors.New("kafka offsets error")
	ErrMessages               KafkaError = errors.New("kafka messages error")
	ErrConsumerGroup          KafkaError = errors.New("kafka consumer group error")
	ErrLifecycle              KafkaError = errors.New("kafka lifecycle error")
	ErrEncodeDecode           KafkaError = errors.New("kafka encode/decode error")
	ErrUnknownError           KafkaError = errors.New("kafka unknown error")
)
