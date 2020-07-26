package mock

import (
	"github.com/statistico/statistico-data/internal/app/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type CompetitionServer struct {
	mock.Mock
	grpc.ServerStream
}

func (c *CompetitionServer) Send(comp *proto.Competition) error {
	args := c.Called(comp)
	return args.Error(0)
}