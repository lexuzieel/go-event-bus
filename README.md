# Go Event Bus

Simple event bus API based on [Watermill](https://github.com/ThreeDotsLabs/watermill) based on Go channel Pub/Sub adapter.

## Usage

```bash
go get github.com/MajorLettuce/go-event-bus
```

## Example

```go
package main

import bus "github.com/MajorLettuce/go-event-bus"

func main() {
	event := "sample-event"

	channel, err := Listen(event)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for incoming message asynchronously
	go func() {
		message := WaitFor(channel)
		log.Printf("Got message, ID: %s", message.UUID)
		log.Printf("Payload: %s", message.Payload)
		for key, value := range message.Metadata {
			log.Printf("Metadata %s=%s", key, value)
		}
	}()

	// Emit a new event "With" metadata and a "Payload"
	New(event).
		With("color", "green").
		With("font-size", "huge").
		Payload([]byte("Sample text")).Emit()

	// Wait a bit so the listening goroutine caches up
	time.Sleep(time.Second)
}
```
