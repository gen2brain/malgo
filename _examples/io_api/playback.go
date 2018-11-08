package io_api

import (
	"context"
	"io"

	"github.com/gen2brain/malgo"
)

// Playback streams samples from a reader to the sound device.
// The function initializes a playback device in the default context using
// provide stream configuration.
// Playback will commence playing the samples provided from the reader until either the
// reader returns an error, or the context signals done.
func Playback(ctx context.Context, r io.Reader, config StreamConfig) error {
	deviceConfig := config.asDeviceConfig()
	abortChan := make(chan error)
	defer close(abortChan)
	aborted := false

	deviceCallbacks := malgo.DeviceCallbacks{
		Send: func(frameCount uint32, samples []byte) uint32 {
			if aborted {
				return 0
			}
			if frameCount == 0 {
				return 0
			}

			frameSize := len(samples) / int(frameCount)
			read, err := r.Read(samples)
			if read <= 0 {
				if err != nil {
					aborted = true
					abortChan <- err
				}
				return 0
			}
			return uint32(read / frameSize)
		},
	}

	return stream(ctx, abortChan, malgo.Playback, deviceConfig, deviceCallbacks)
}
