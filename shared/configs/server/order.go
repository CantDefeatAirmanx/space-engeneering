package configs_order

import "time"

const (
	Port              = 8080
	Timeout           = 10 * time.Second
	ReadHeaderTimeout = 5 * time.Second
	ShutdownTimeout   = 10 * time.Second
)
