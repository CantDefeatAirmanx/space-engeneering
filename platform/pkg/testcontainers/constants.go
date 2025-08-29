package testcontainers

import (
	"fmt"
	"math/rand"
	"time"
)

// MongoDB constants
const (
	// MongoDB container constants
	MongoContainerNameBase = "mongo"
	MongoPort              = "27017"

	// MongoDB environment variables
	MongoImageNameKey = "MONGO_IMAGE_NAME"
	MongoHostKey      = "MONGO_HOST"
	MongoPortKey      = "MONGO_PORT"
	MongoDatabaseKey  = "MONGO_DATABASE"
	MongoUsernameKey  = "MONGO_INITDB_ROOT_USERNAME"
	MongoPasswordKey  = "MONGO_INITDB_ROOT_PASSWORD" //nolint:gosec
	MongoAuthDBKey    = "MONGO_AUTH_DB"
)

func GenerateMongoContainerName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s-%d", MongoContainerNameBase, rand.Intn(100000))
}
