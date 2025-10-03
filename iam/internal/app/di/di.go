package di

import (
	"context"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/iam/config"
	api_auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/api/auth/v1"
	api_user_v1 "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/api/user/v1"
	repository_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method"
	repository_notification_method_postgres "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method/postgres"
	repository_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/session"
	repository_session_redis "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/session/redis"
	repository_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user"
	repository_user_postgres "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user/postgres"
	service_auth "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/service/auth"
	service_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/service/user"
	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
	platform_redis_redisgo "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis/go-redis_impl"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	platform_pwdhasher "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/pwd-hasher"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
	platform_transaction_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction/postgres"
)

type DiContainer struct {
	closer closer.Closer

	userV1API   *api_user_v1.Api
	userService service_user.UserService

	authV1API   *api_auth_v1.Api
	authService service_auth.AuthService

	userRepository               repository_user.UserRepository
	notificationMethodRepository repository_notification_method.NotificationMethodRepository
	sessionRepository            repository_session.SessionRepository

	pwdHasher platform_pwdhasher.PwdHasher
	txManager platform_transaction.TxManager[platform_transaction.Executor]
	db        *pgxpool.Pool
	redis     platform_redis.RedisCache
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetUserV1API(ctx context.Context) *api_user_v1.Api {
	if d.userV1API != nil {
		return d.userV1API
	}

	d.userV1API = api_user_v1.NewApi(
		d.GetUserService(ctx),
	)

	return d.userV1API
}

func (d *DiContainer) GetAuthV1API(ctx context.Context) *api_auth_v1.Api {
	if d.authV1API != nil {
		return d.authV1API
	}

	d.authV1API = api_auth_v1.NewApi(
		d.GetAuthService(ctx),
	)

	return d.authV1API
}

func (d *DiContainer) GetUserService(ctx context.Context) service_user.UserService {
	if d.userService != nil {
		return d.userService
	}

	userService := service_user.NewUserServiceImpl(
		d.GetUserRepository(ctx),
		d.GetNotificationMethodRepository(ctx),
		d.GetTxManager(ctx),
		d.GetPwdHasher(ctx),
	)
	d.userService = userService

	return userService
}

func (d *DiContainer) GetAuthService(ctx context.Context) service_auth.AuthService {
	if d.authService != nil {
		return d.authService
	}

	authService := service_auth.NewAuthService(
		d.GetUserRepository(ctx),
		d.GetNotificationMethodRepository(ctx),
		d.GetSessionRepository(ctx),
		d.GetPwdHasher(ctx),
	)
	d.authService = authService

	return authService
}

func (d *DiContainer) GetUserRepository(ctx context.Context) repository_user.UserRepository {
	if d.userRepository != nil {
		return d.userRepository
	}

	repository := repository_user_postgres.NewUserRepositoryPostgres(
		d.getPostgres(ctx),
	)
	d.userRepository = repository

	return repository
}

func (d *DiContainer) GetNotificationMethodRepository(
	ctx context.Context,
) repository_notification_method.NotificationMethodRepository {
	if d.notificationMethodRepository != nil {
		return d.notificationMethodRepository
	}

	repository := repository_notification_method_postgres.NewNotificationMethodRepositoryPostgres(
		d.getPostgres(ctx),
	)
	d.notificationMethodRepository = repository

	return repository
}

func (d *DiContainer) GetSessionRepository(ctx context.Context) repository_session.SessionRepository {
	if d.sessionRepository != nil {
		return d.sessionRepository
	}

	repository := repository_session_redis.NewSessionRepositoryRedisImpl(
		d.GetRedis(ctx),
	)
	d.sessionRepository = repository

	return repository
}

func (d *DiContainer) GetPwdHasher(ctx context.Context) platform_pwdhasher.PwdHasher {
	if d.pwdHasher != nil {
		return d.pwdHasher
	}

	pwdHasher := platform_pwdhasher.NewPwdHasherImpl()
	d.pwdHasher = pwdHasher

	return pwdHasher
}

func (d *DiContainer) getPostgres(
	ctx context.Context,
) *pgxpool.Pool {
	if d.db != nil {
		return d.db
	}

	db, err := pgxpool.New(
		ctx,
		config.Config.Postgres().Uri(),
	)
	if err != nil {
		logger.Logger().Error("Failed to create postgres", zap.Error(err))
		panic(err)
	}
	d.db = db
	d.closer.AddNamed("Postgres", func(ctx context.Context) error {
		db.Close()
		return nil
	})

	return db
}

func (d *DiContainer) GetRedis(ctx context.Context) platform_redis.RedisCache {
	if d.redis != nil {
		return d.redis
	}

	redis, err := platform_redis_redisgo.NewSingleNodeClient(
		net.JoinHostPort(
			config.Config.Redis().Host(),
			strconv.Itoa(config.Config.Redis().ExternalPort()),
		),
	)
	if err != nil {
		logger.Logger().Error("Failed to create redis", zap.Error(err))
		panic(err)
	}
	d.closer.AddNamed("Redis", func(ctx context.Context) error {
		return redis.Close()
	})
	d.redis = redis

	return redis
}

func (d *DiContainer) GetTxManager(ctx context.Context) platform_transaction.TxManager[platform_transaction.Executor] {
	if d.txManager != nil {
		return d.txManager
	}

	txManager := platform_transaction_postgres.NewTransactionManagerPostgres(
		d.getPostgres(ctx),
	)
	d.txManager = txManager

	return txManager
}
