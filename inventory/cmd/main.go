package main

import (
	"context"
	"log"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	if err := app.Run(ctx); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
