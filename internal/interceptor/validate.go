package interceptor

import (
	"context"
	"flag"
	"fmt"
	"log"

	descAccess "github.com/BelyaevEI/microservices_chat/pkg/access_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var accessToken = flag.String("a", "", "access token")

const servicePort = 50051

type validator interface {
	Validate() error
}

// ValidateInterceptor interceptor for proto validate
func ValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.NewClient(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)
	_, err = cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: info.FullMethod,
	})
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
