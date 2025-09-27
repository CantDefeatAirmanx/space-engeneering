package config_kafka

type KafkaConfig struct {
	Brokers        []string `env:"brokers,required"`
	ExternalPort   int      `env:"externalPort,required"`
	InternalPort   int      `env:"internalPort,required"`
	ControllerPort int      `env:"controllerPort,required"`
	UiPort         int      `env:"uiPort,required"`
}
