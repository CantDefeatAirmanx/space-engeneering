package service_topics_configurator

type KafkaConfigYaml struct {
	DefaultTopicConfig DefaultTopicConfig `yaml:"default-topic-config" json:"default-topic-config"`
	Topics             []TopicConfig      `yaml:"topics" json:"topics"`
}

type DefaultTopicConfig struct {
	Partitions        int    `yaml:"partitions" json:"partitions"`
	ReplicationFactor int    `yaml:"replication-factor" json:"replication-factor"`
	MinInSyncReplicas int    `yaml:"min-insync-replicas" json:"min-insync-replicas"`
	RetentionMs       int64  `yaml:"retention-ms" json:"retention-ms"`
	CleanupPolicy     string `yaml:"cleanup-policy" json:"cleanup-policy"`
}

type TopicConfig struct {
	Name              string `yaml:"name" json:"name"`
	Partitions        int    `yaml:"partitions,omitempty" json:"partitions,omitempty"`
	ReplicationFactor int    `yaml:"replication-factor,omitempty" json:"replication-factor,omitempty"`
}
