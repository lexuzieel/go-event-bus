package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type MessageChannel <-chan *message.Message

var bus = gochannel.NewGoChannel(
	gochannel.Config{},
	watermill.NewStdLogger(false, false),
)

func Listen(event string) (MessageChannel, error) {
	return bus.Subscribe(context.Background(), event)
}

type EmitEventChain struct {
	event   string
	message *message.Message
}

func New(name string) *EmitEventChain {
	return &EmitEventChain{
		event:   name,
		message: message.NewMessage(watermill.NewShortUUID(), nil),
	}
}

func (e *EmitEventChain) With(key string, value string) *EmitEventChain {
	e.message.Metadata.Set(key, value)

	return e
}

func (e *EmitEventChain) Payload(payload []byte) *EmitEventChain {
	e.message.Payload = payload

	return e
}

func (e *EmitEventChain) Emit() error {
	return bus.Publish(e.event, e.message)
}

func WaitFor(channel MessageChannel) *message.Message {
	message := <-channel
	message.Ack()
	return message
}
