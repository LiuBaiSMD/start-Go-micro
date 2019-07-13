package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	namespace "namespace/proto/namespace"
)

type Namespace struct{}

func (e *Namespace) Handle(ctx context.Context, msg *namespace.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *namespace.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
