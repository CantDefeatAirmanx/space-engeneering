package platform_kafka_consumer

import (
	"time"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/IBM/sarama"
)

var _ platform_kafka.Consumer = (*KafkaConsumerImpl)(nil)

type KafkaConsumerImpl struct {
	consumerGroup      sarama.ConsumerGroup
	consumeErrHandlers []func(err error)
}

func NewKafkaConsumer(
	brokers []string,
	groupID string,
	opts ...ConsumerConfigOption,
) (*KafkaConsumerImpl, error) {
	cfg := NewConsumerConfig(brokers, groupID, opts...)

	saramaCfg := sarama.NewConfig()

	saramaCfg.Version = platform_kafka.KafkaVersion
	saramaCfg.Metadata.AllowAutoTopicCreation = false

	// Auto commit настройки
	saramaCfg.Consumer.Offsets.AutoCommit.Enable = cfg.EnableAutoCommit
	if cfg.EnableAutoCommit {
		saramaCfg.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	}

	// Offset reset стратегия
	saramaCfg.Consumer.Offsets.Initial = getInitialOffsetStart(cfg.InitialOffsetStart)

	// Network & Connection настройки
	saramaCfg.Net.MaxOpenRequests = cfg.MaxOpenRequests
	saramaCfg.Net.ReadTimeout = cfg.RequestTimeout
	saramaCfg.Net.WriteTimeout = cfg.RequestTimeout

	// Consumer Group настройки
	saramaCfg.Consumer.Group.Session.Timeout = cfg.SessionTimeout
	saramaCfg.Consumer.Group.Heartbeat.Interval = cfg.HeartbeatInterval
	saramaCfg.Consumer.Group.Rebalance.Timeout = cfg.RebalanceTimeout

	// Partition assignment strategy
	switch cfg.PartitionStrategy {
	case "roundrobin":
		saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "sticky":
		saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	default: // range
		saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	}

	// Fetch настройки
	saramaCfg.Consumer.Fetch.Min = int32(cfg.FetchMinBytes)
	saramaCfg.Consumer.Fetch.Max = int32(cfg.FetchMaxBytes)
	saramaCfg.Consumer.MaxWaitTime = cfg.FetchMaxWait

	// Retry настройки
	saramaCfg.Consumer.Retry.Backoff = cfg.RetryBackoff
	saramaCfg.Metadata.Retry.Max = cfg.RetryMax
	saramaCfg.Metadata.Retry.Backoff = cfg.RetryBackoff

	consumerGroup, err := sarama.NewConsumerGroup(
		cfg.Brokers,
		cfg.GroupID,
		saramaCfg,
	)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumerImpl{
		consumerGroup:      consumerGroup,
		consumeErrHandlers: cfg.ConsumeErrHandlers,
	}, nil
}

func getInitialOffsetStart(initialOffsetStart InitialOffsetStart) int64 {
	switch initialOffsetStart {
	case "earliest":
		return sarama.OffsetOldest
	case "latest":
		return sarama.OffsetNewest
	default:
		return sarama.OffsetOldest
	}
}
