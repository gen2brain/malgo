package malgo

import (
	"context"
	"io"

	"github.com/gen2brain/malgo/mini_al"
)

// Capture records incoming samples into the provided writer.
// The function initializes a capture device in the default context using
// provide stream configuration.
// Capturing will commence writing the samples to the writer until either the
// writer returns an error, or the context signals done.
func Capture(ctx context.Context, w io.Writer, config StreamConfig) error {
	deviceConfig := config.asDeviceConfig()
	abortChan := make(chan error)
	defer close(abortChan)
	aborted := false

	deviceCallbacks := mini_al.DeviceCallbacks{
		Recv: func(frameCount uint32, samples []byte) {
			if aborted {
				return
			}

			_, err := w.Write(samples)
			if err != nil {
				aborted = true
				abortChan <- err
			}
		},
	}

	return stream(ctx, abortChan, mini_al.Capture, deviceConfig, deviceCallbacks)
}
