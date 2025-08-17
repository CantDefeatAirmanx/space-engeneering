package main

import (
	"context"
	"log"

	"github.com/CantDefeatAirmanx/space-engeneering/order/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
