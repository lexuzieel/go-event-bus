package events

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventEmitAndListen(t *testing.T) {
	event := "sample-event"

	channel, err := Listen(event)
	if err != nil {
		t.Error(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		message := WaitFor(channel)
		require.Equal(t, message.Metadata["value"], "metadata")
		require.Equal(t, string(message.Payload), "hello")
		wg.Done()
	}()

	err = New(event).
		With("value", "metadata").
		Payload([]byte("hello")).
		Emit()
	if err != nil {
		t.Error(err)
	}

	wg.Wait()
}
