package platform_kafka_converter

import (
	"errors"

	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

func SaramaMessageToPlatformMessage(saramaMessage *sarama.ConsumerMessage) platform_kafka.Message {
	return platform_kafka.Message{
		Headers:        ConvertPointerHeaders(saramaMessage.Headers),
		Timestamp:      saramaMessage.Timestamp,
		BlockTimestamp: saramaMessage.BlockTimestamp,
		Topic:          saramaMessage.Topic,
		Key:            saramaMessage.Key,
		Value:          saramaMessage.Value,
		Partition:      saramaMessage.Partition,
		Offset:         saramaMessage.Offset,
	}
}

func ConvertPointerHeaders(headers []*sarama.RecordHeader) platform_kafka.ProducerHeaders {
	producerHeaders := make(platform_kafka.ProducerHeaders)

	for idx := range headers {
		header := headers[idx]
		producerHeaders[string(header.Key)] = header.Value
	}

	return producerHeaders
}

func ConvertHeaders(headers []sarama.RecordHeader) platform_kafka.ProducerHeaders {
	producerHeaders := make(platform_kafka.ProducerHeaders)

	for idx := range headers {
		header := headers[idx]
		producerHeaders[string(header.Key)] = header.Value
	}

	return producerHeaders
}

func SaramaMessageToProducerMessage(
	saramaMessage *sarama.ProducerMessage,
) (*platform_kafka.ProducerMessage, error) {
	key, err := saramaMessage.Key.Encode()
	if err != nil {
		return nil, err
	}

	value, err := saramaMessage.Value.Encode()
	if err != nil {
		return nil, err
	}

	mess := platform_kafka.NewProducerMessage(
		saramaMessage.Topic,
		key,
		value,
		platform_kafka.WithHeaders(ConvertHeaders(saramaMessage.Headers)),
	)

	return &mess, nil
}

//nolint:cyclop
func ConvertSaramaError(err error) platform_kafka.KafkaError {
	switch {
	// Consumer message handler errors
	case errors.Is(err, platform_kafka.ErrConsumerMessageHandler): // Ошибка обработчика сообщений
		return err

	// Network errors
	case errors.Is(err, sarama.ErrOutOfBrokers): // Все брокеры недоступны
	case errors.Is(err, sarama.ErrNotConnected): // Соединение разорвано
	case errors.Is(err, sarama.ErrRequestTimedOut): // Таймаут запроса
	case errors.Is(err, sarama.ErrBrokerNotFound): // Брокер не найден
	case errors.Is(err, sarama.ErrBrokerNotAvailable): // Брокер не доступен
		return errors.Join(platform_kafka.ErrNetworkError, err)

	// Coordination errors
	case errors.Is(err, sarama.ErrConsumerCoordinatorNotAvailable): // Координатор недоступен
	case errors.Is(err, sarama.ErrNotCoordinatorForConsumer): // Неверный координатор
	case errors.Is(err, sarama.ErrRebalanceInProgress): // Идет ребалансировка
	case errors.Is(err, sarama.ErrIllegalGeneration): // Неверное поколение группы
	case errors.Is(err, sarama.ErrUnknownMemberId): // Неизвестный член группы
		return errors.Join(platform_kafka.ErrCoordination, err)

	// Partition errors
	case errors.Is(err, sarama.ErrLeaderNotAvailable): // Лидер партиции недоступен
	case errors.Is(err, sarama.ErrNotLeaderForPartition): // Брокер не лидер партиции
	case errors.Is(err, sarama.ErrReplicaNotAvailable): // Реплика недоступна
		return errors.Join(platform_kafka.ErrPartition, err)

	// Configuration errors
	case errors.Is(err, sarama.ErrInvalidConfig): // Неверная конфигурация
	case errors.Is(err, sarama.ErrUnsupportedVersion): // Неподдерживаемая версия протокола
	case errors.Is(err, sarama.ErrInconsistentGroupProtocol): // Несовместимый протокол группы
		return errors.Join(platform_kafka.ErrConfiguration, err)

	// Offsets errors
	case errors.Is(err, sarama.ErrOffsetOutOfRange): // Offset вне диапазона
	case errors.Is(err, sarama.ErrInvalidCommitOffsetSize): // Неверный размер offset
	case errors.Is(err, sarama.ErrOffsetMetadataTooLarge): // Слишком большие metadata offset'а
		return errors.Join(platform_kafka.ErrOffsets, err)

	// Messages errors
	case errors.Is(err, sarama.ErrMessageTooLarge): // Сообщение превышает лимит
	case errors.Is(err, sarama.ErrInvalidMessage): // Некорректное сообщение
	case errors.Is(err, sarama.ErrInvalidMessageSize): // Неверный размер сообщения
	case errors.Is(err, sarama.ErrConsumerOffsetNotAdvanced): // Offset не продвинулся
		return errors.Join(platform_kafka.ErrMessages, err)

	// Consumer group errors
	case errors.Is(err, sarama.ErrUnknownTopicOrPartition): // Топик/партиция не существует
	case errors.Is(err, sarama.ErrInvalidGroupId): // Неверный ID группы
	case errors.Is(err, sarama.ErrInvalidSessionTimeout): // Неверный session timeout
		return errors.Join(platform_kafka.ErrConsumerGroup, err)

	// Authorization errors
	case errors.Is(err, sarama.ErrSASLAuthenticationFailed): // Ошибка SASL аутентификации
	case errors.Is(err, sarama.ErrTopicAuthorizationFailed): // Нет прав на топик
	case errors.Is(err, sarama.ErrGroupAuthorizationFailed): // Нет прав на группу
	case errors.Is(err, sarama.ErrClusterAuthorizationFailed): // Нет прав на кластер
		return errors.Join(platform_kafka.ErrAuthorization, err)

	// Lifecycle errors
	case errors.Is(err, sarama.ErrClosedConsumerGroup): // Группа консьюмеров закрыта
	case errors.Is(err, sarama.ErrClosedClient): // Клиент закрыт
	case errors.Is(err, sarama.ErrShuttingDown): // Происходит shutdown
	case errors.Is(err, sarama.ErrInvalidPartition): // Неверная партиция
		return errors.Join(platform_kafka.ErrLifecycle, err)

	// Encode/decode errors
	case errors.Is(err, sarama.PacketDecodingError{}): // Ошибка кодирования пакета
		return errors.Join(platform_kafka.ErrEncodeDecode, err)
	}

	return errors.Join(platform_kafka.ErrUnknownError, err)
}
