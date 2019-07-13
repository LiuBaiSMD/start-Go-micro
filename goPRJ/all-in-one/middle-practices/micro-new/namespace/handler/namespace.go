package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	namespace "namespace/proto/namespace"
)

type Namespace struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Namespace) Call(ctx context.Context, req *namespace.Request, rsp *namespace.Response) error {
	log.Log("Received Namespace.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Namespace) Stream(ctx context.Context, req *namespace.StreamingRequest, stream namespace.Namespace_StreamStream) error {
	log.Logf("Received Namespace.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&namespace.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Namespace) PingPong(ctx context.Context, stream namespace.Namespace_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&namespace.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
