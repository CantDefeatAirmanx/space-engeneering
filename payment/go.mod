module github.com/CantDefeatAirmanx/space-engeneering/payment

go 1.24.4

require (
	github.com/CantDefeatAirmanx/space-engeneering/platform v0.0.0-00010101000000-000000000000
	github.com/CantDefeatAirmanx/space-engeneering/shared v0.0.0-20250814190328-0af679cc83ba
	github.com/caarlos0/env/v11 v11.3.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.74.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/protobuf v1.36.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/CantDefeatAirmanx/space-engeneering/platform => ../platform

replace github.com/CantDefeatAirmanx/space-engeneering/shared => ../shared
