package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventEmitAndListen(t *testing.T) {
	eventName := "sample-event"

	channel, err := Listen(eventName)
	if err != nil {
		t.Error(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		message := WaitFor(channel)
		assert.Equal(t, message.Metadata["field"], "metadata")
		assert.Equal(t, string(message.Payload), "hello")
		wg.Done()
	}()

	err = Event(eventName).
		With("field", "metadata").
		Payload([]byte("hello")).
		Emit()
	if err != nil {
		t.Error(err)
	}

	wg.Wait()
}
