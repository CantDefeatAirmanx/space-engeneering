package platform_kafka_client

import "github.com/IBM/sarama"

func (kc *KafkaClient) DescribeCluster() (
	brokers []*sarama.Broker,
	controllerID int32,
	err error,
) {
	return kc.admin.DescribeCluster()
}
