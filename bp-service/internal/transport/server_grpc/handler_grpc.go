package server_grpc

import (
	pr "bp-service/internal/pb/grpc-pressure/proto"
	tg "bp-service/internal/pb/grpc-tag/proto"
	"bp-service/internal/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HandlerGRPC struct {
	*PressureGRPC
	*TagGRPC
}

func NewHandlerGRPC(svc *service.Service, grpcServer *grpc.Server) *HandlerGRPC {

	pressureGRPC := NewPressureGRPC(svc)
	tagGRPC := NewTagGRPC(svc)
	pr.RegisterBloodPressureServiceServer(grpcServer, pressureGRPC)
	tg.RegisterTagServiceServer(grpcServer, tagGRPC)

	return &HandlerGRPC{
		PressureGRPC: pressureGRPC,
		TagGRPC:      tagGRPC,
	}
}

func tokenValid(ctx context.Context) int {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	tokenApi := md.Get("token_api")[0]
	fmt.Println("token:", tokenApi)

	//if tokenApi != os.Getenv("SERVICE_TOKEN_API") {
	//	return 0
	//}
	return 1
}
