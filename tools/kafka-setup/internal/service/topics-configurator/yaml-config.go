package service_topics_configurator

type KafkaConfigYaml struct {
	DefaultTopicConfig DefaultTopicConfig `yaml:"default-topic-config"`
	Topics             []TopicConfig      `yaml:"topics"`
}

type DefaultTopicConfig struct {
	Partitions        int    `yaml:"partitions"`
	ReplicationFactor int    `yaml:"replication-factor"`
	MinInSyncReplicas int    `yaml:"min-insync-replicas"`
	RetentionMs       int64  `yaml:"retention-ms"`
	CleanupPolicy     string `yaml:"cleanup-policy"`
}

type TopicConfig struct {
	Name              string `yaml:"name"`
	Partitions        int    `yaml:"partitions,omitempty"`
	ReplicationFactor int    `yaml:"replication-factor,omitempty"`
}
