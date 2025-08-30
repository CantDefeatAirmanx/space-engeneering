package platform_mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mongo_options "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultConnectTimeout = 20 * time.Second
	defaultPingTimeout    = 5 * time.Second
)

func Connect(ctx context.Context, opts ...Option) (*mongo.Client, error) {
	options := &Options{
		ConnectTimeout: defaultConnectTimeout,
		PingTimeout:    defaultPingTimeout,
	}
	for _, opt := range opts {
		opt(options)
	}

	if options.URI == "" {
		return nil, ErrURIRequired
	}

	connectCtx, connectCtxCancel := context.WithTimeout(
		ctx,
		options.ConnectTimeout,
	)
	defer connectCtxCancel()

	clientOptions := mongo_options.Client().
		ApplyURI(options.URI).
		SetConnectTimeout(options.ConnectTimeout).
		SetServerSelectionTimeout(options.ConnectTimeout)

	client, err := mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		return nil, err
	}

	pingCtx, pingCtxCancel := context.WithTimeout(
		ctx,
		options.PingTimeout,
	)
	defer pingCtxCancel()

	err = client.Ping(pingCtx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
