package main

import (
	"context"
	"fmt"

	"github.com/CantDefeatAirmanx/space-engeneering/order/internal/app"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

func main() {
	ctx := context.Background()
	closer, done := closer.NewCloser(ctx)

	defer func() {
		go closer.CloseAll(ctx)
		<-done
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic in main goroutine, closing. %v\n", r)
		}
	}()

	app, err := app.NewApp(ctx, closer)
	if err != nil {
		panic(fmt.Sprintf("failed to create app: %v", err))
	}

	logger.Logger().Info("Starting app")
	err = app.Run(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to run app: %v", err))
	}
}
