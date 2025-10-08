package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
)

type (
	SessionUUIDCtxKeyType string
	UserInfoCtxKeyType    string
)

const (
	SessionUUIDMetadataKey                       = "session-uuid"
	SessionUUIDCtxKey      SessionUUIDCtxKeyType = "session-uuid"
	UserInfoCtxKey         UserInfoCtxKeyType    = "user-info"
)

var (
	ErrInvalidArguments = business_errors.NewBadRequestError(errors.New("session uuid is required"))
	ErrUnauthorized     = business_errors.NewUnauthorizedError(errors.New("session uuid is required"))

	ErrIncomingCtxNotFoundValue = business_errors.NewBadRequestError(errors.New("incoming ctx not found value"))
)

func AuthInterceptor(
	authClient auth_v1.AuthServiceClient,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		sessionUUID, err := getSessionUUIDFromMetadata(ctx)
		if err != nil {
			return nil, ErrInvalidArguments
		}
		res, err := authClient.WhoAmI(ctx, &auth_v1.WhoAmIRequest{
			SessionUuid: sessionUUID,
		})
		if err != nil {
			return nil, ErrUnauthorized
		}

		withSessionUUIDCtx := GetCtxWithSessionUUID(ctx, sessionUUID)
		withUserInfoCtx := GetCtxWithUserInfo(withSessionUUIDCtx, res.User)

		enrichedCtx := withUserInfoCtx

		return handler(enrichedCtx, req)
	}
}

func ForwardSessionUUIDToOutgoingCtx(ctx context.Context) context.Context {
	sessionUUID, err := GetSessionUUIDFromIncomingCtx(ctx)
	if err != nil {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, SessionUUIDMetadataKey, sessionUUID)
}

func GetSessionUUIDFromIncomingCtx(ctx context.Context) (string, error) {
	sessionUUID, ok := ctx.Value(SessionUUIDCtxKey).(string)
	if !ok {
		return "", ErrIncomingCtxNotFoundValue
	}
	return sessionUUID, nil
}

func GetUserInfoFromIncomingCtx(ctx context.Context) (*common_v1.User, error) {
	userInfo, ok := ctx.Value(UserInfoCtxKey).(*common_v1.User)
	if !ok {
		return nil, ErrIncomingCtxNotFoundValue
	}
	return userInfo, nil
}

func GetCtxWithUserInfo(ctx context.Context, user *common_v1.User) context.Context {
	return context.WithValue(ctx, UserInfoCtxKey, user)
}

func GetCtxWithSessionUUID(ctx context.Context, sessionUUID string) context.Context {
	return context.WithValue(ctx, SessionUUIDCtxKey, sessionUUID)
}

func getSessionUUIDFromMetadata(reqCtx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(reqCtx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "metadata is required")
	}

	sessionUUIDValues := meta.Get(SessionUUIDMetadataKey)
	if len(sessionUUIDValues) == 0 {
		return "", status.Errorf(codes.InvalidArgument, "session_uuid is required")
	}

	sessionUUID := sessionUUIDValues[0]
	if sessionUUID == "" {
		return "", status.Errorf(codes.InvalidArgument, "session_uuid is required")
	}

	return sessionUUID, nil
}
