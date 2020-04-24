package server

import (
	"context"

	"distributed/ch04"
	api "distributed/ch04/api/v1"

	_ "google.golang.org/grpc"
)

var _ api.LogServer = (*grpcServer)(nil)

type Config struct {
	CommitLog log.CommitLog
}

type grpcServer struct {
	*Config
}

func newgrpcServer(config *Config) (*grpcServer, error) {
	s := &grpcServer{
		Config: config,
	}
	return s, nil
}

func (s *grpcServer) Produce(ctx context.Context, req *api.ProduceRequest) (*api.ProduceResponse, error) {
	offset, err := s.CommitLog.Append(req.Record)
	if err != nil {
		return nil, err
	}
	return &api.ProduceResponse{Offset: offset}, nil
}

func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	record, err := s.CommitLog.Read(req.Offset)
	if err != nil {
		return nil, err
	}
	return &api.ConsumeResponse{Record: record}, nil
}

func (s *grpcServer) ProduceStream(
	stream api.Log_ProduceStreamServer,
) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		resp, err := s.Produce(stream.Context(), req)
		if err != nil {
			return err
		}
		if err = stream.Send(resp); err != nil {
			return err
		}
	}
}

func (s *grpcServer) ConsumeStream(
	req *api.ConsumeRequest,
	stream api.Log_ConsumeStreamServer,
) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			resp, err := s.Consume(stream.Context(), req)
			switch err.(type) {
			case nil:
			case api.ErrOffsetOutOfRange:
				continue
			default:
				return err
			}
			if err = stream.Send(resp); err != nil {
				return err
			}
			req.Offset++
		}
	}
}
