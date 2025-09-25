package config_kafka

type KafkaConfigData struct {
	Brokers       []string `env:"brokers,required"`
	OrderTopic    string   `env:"orderTopic,required"`
	AssemblyTopic string   `env:"assemblyTopic,required"`
}
