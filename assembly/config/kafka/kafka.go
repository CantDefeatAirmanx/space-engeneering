package config_kafka

type KafkaConfigData struct {
	Brokers       []string `env:"brokers,required"`
	AssemblyTopic string   `env:"assemblyTopic,required"`
	OrderTopic    string   `env:"orderTopic,required"`
}
