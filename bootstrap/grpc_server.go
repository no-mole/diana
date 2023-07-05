package bootstrap

import (
	"context"

	"diana/service/bar"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/no-mole/neptune/app"
	"github.com/no-mole/neptune/middlewares"
	barPb "github.com/no-mole/neptune/protos/bar"
)

func InitGrpcServer(_ context.Context) error {
	s := app.NewGrpcServer(
		grpc_middleware.WithUnaryServerChain(
			middleware.TracingServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			middleware.TracingServerStreamInterceptor(),
		),
	)
	s.RegisterService(&barPb.Metadata().ServiceDesc, bar.New())
	return nil
}
