package bootstrap

import (
	"context"
	"google.golang.org/grpc"

	"github.com/no-mole/neptune/grpc_pool"
	"github.com/no-mole/neptune/middlewares"
	"github.com/no-mole/neptune/protos/bar"
	"github.com/no-mole/neptune/registry"
)

func DiscoveryService(ctx context.Context) error {
	option := grpc_pool.WithOptions(
		grpc_pool.WithDialOptions(
			grpc.WithChainUnaryInterceptor(
				middleware.TracingClientInterceptor(),
			),
		),
	)
	grpc_pool.Init(ctx,
		option,
		bar.Metadata(),
	)
	return nil
}

func RegistrationService(ctx context.Context) error {
	return registry.Registry(ctx,
		bar.Metadata(),
	)
}
