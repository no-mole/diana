package bar

import (
	"context"
	"fmt"

	"diana/model/bar"
	pb "github.com/no-mole/neptune/protos/bar"
	"github.com/no-mole/neptune/registry"
)

type Service struct {
	*registry.Metadata
	pb.UnimplementedServiceServer
}

func New() *Service {
	return &Service{
		Metadata: pb.Metadata(),
	}
}

func (s *Service) SayHelly(ctx context.Context, in *pb.SayHelloRequest) (ret *pb.SayHelloResponse, err error) {
	defer func() {
		model := bar.New(ctx)
		_, err = model.Log(in.Say, ret.Reply)
	}()
	return &pb.SayHelloResponse{
		Reply: fmt.Sprintf("your say %s,i say %s too!", in.Say, in.Say),
	}, nil
}
