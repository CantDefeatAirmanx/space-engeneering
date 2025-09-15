# ============================
# Stage 1: Build stage
# ============================
ARG GO_VERSION=1.24.4
ARG BUILD_OUTPUT_INVENTORY=app-inventory

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Копируем go.work и его контрольную сумму — это позволяет Go правильно связать модули

COPY go.work go.work.sum ./
# Копируем go.mod и go.sum всех зависимых модулей — нужно для кэширования go mod download
COPY platform/go.mod platform/go.sum ./platform/
COPY shared/go.mod shared/go.sum ./shared/
COPY order/go.mod order/go.sum ./order/
COPY inventory/go.mod inventory/go.sum ./inventory/
COPY payment/go.mod payment/go.sum ./payment/

# Загружаем все зависимости, указанные в модульных файлах
RUN go mod download

# Копируем исходный код всех модулей — включая зависимости, т.к. они нужны при компиляции
COPY platform ./platform
COPY shared ./shared
COPY order ./order
COPY inventory ./inventory
COPY payment ./payment

# ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 ./grpc-health-probe
# RUN chmod +x grpc-health-probe

# Собираем бинарный файл Inventory-сервиса для Linux-архитектуры без CGO
WORKDIR /app/inventory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app-inventory ./cmd


# ============================
# Stage 2: Final image
# ============================

FROM alpine:3.21.3

# Создаём системного пользователя без root-прав — best practice для безопасности
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
USER appuser

# Копируем скомпилированный бинарник из стадии builder
COPY --from=builder /app/inventory/app-inventory .

# Копируем бинарь grpc-health-probe в /bin
# COPY --from=builder /app/grpc-health-probe /bin/grpc-health-probe

# Экспонируем порт gRPC-сервиса Inventory
EXPOSE 50051

# Устанавливаем команду запуска — запускаем наш бинарь
ENTRYPOINT ["./app-inventory"]