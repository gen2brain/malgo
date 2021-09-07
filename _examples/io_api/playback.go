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
	deviceConfig := config.asDeviceConfig(malgo.Playback)
	abortChan := make(chan error)
	defer close(abortChan)
	aborted := false

		deviceConfig.DataCallback = func(_ malgo.Device, outputSamples, inputSamples *malgo.DataBuffer, frameCount int) {
			if aborted {
				return
			}
			if frameCount == 0 {
				return
			}

			read, err := io.ReadFull(r, outputSamples.Bytes())
			if read <= 0 {
				if err != nil {
					aborted = true
					abortChan <- err
				}
				return
			}
		}

	return stream(ctx, abortChan, deviceConfig)
}
