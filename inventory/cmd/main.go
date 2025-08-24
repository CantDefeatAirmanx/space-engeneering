package main

import (
	"context"
	"fmt"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/app"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

func main() {
	ctx := context.Background()
	closer, done := closer.NewCloser(ctx)

	defer func() {
		go func() {
			if err := closer.CloseAll(ctx); err != nil {
				logger.Logger().Error(fmt.Sprintf("Failed to close app: %v", err))
			}
		}()
		<-done
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic in main goroutine, closing. %v\n", r)
		}
	}()

	app, err := app.NewApp(ctx, closer)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize app: %v", err))
	}

	if err := app.Run(ctx); err != nil {
		panic(fmt.Sprintf("Failed to run app: %v", err))
	}
}
