package malgo_test

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/gen2brain/malgo"
)

func TestCapturePlayback(t *testing.T) {
	onLog := func(message string) {
		fmt.Fprintf(ioutil.Discard, message)
	}

	ctx, err := malgo.InitContext([]malgo.Backend{malgo.BackendNull}, malgo.ContextConfig{}, onLog)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig()
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 2
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Playback.Format))
	onRecvFrames := func(outpuSamples, inputSamples []byte, framecount uint32) {
		sampleCount := framecount * deviceConfig.Playback.Channels * sizeInBytes

		newCapturedSampleCount := capturedSampleCount + sampleCount

		pCapturedSamples = append(pCapturedSamples, inputSamples...)

		capturedSampleCount = newCapturedSampleCount
	}

	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	captureDeviceConfig := deviceConfig
	captureDeviceConfig.DeviceType = malgo.Capture
	device, err := malgo.InitDevice(ctx.Context, captureDeviceConfig, captureCallbacks)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != malgo.Capture {
		t.Errorf("wrong device type")
	}

	if device.PlaybackFormat() != malgo.FormatS16 {
		t.Errorf("wrong format")
	}

	if device.PlaybackChannels() != 2 {
		t.Errorf("wrong number of channels")
	}

	if device.SampleRate() != 44100 {
		t.Errorf("wrong samplerate")
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	if !device.IsStarted() {
		t.Fatalf("device not started")
	}

	time.Sleep(1 * time.Second)

	device.Uninit()

	onSendFrames := func(outputSamples, inputSamples []byte, framecount uint32) {
		samplesToRead := framecount * deviceConfig.Playback.Channels * sizeInBytes
		if samplesToRead > capturedSampleCount-playbackSampleCount {
			samplesToRead = capturedSampleCount - playbackSampleCount
		}

		copy(outputSamples, pCapturedSamples[playbackSampleCount:playbackSampleCount+samplesToRead])

		playbackSampleCount += samplesToRead
	}

	playbackCallbacks := malgo.DeviceCallbacks{
		Data: onSendFrames,
	}

	playbackDeviceConfig := deviceConfig
	playbackDeviceConfig.DeviceType = malgo.Playback
	device, err = malgo.InitDevice(ctx.Context, playbackDeviceConfig, playbackCallbacks)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != malgo.Playback {
		t.Errorf("wrong device type")
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	device.Uninit()
}

func TestErrors(t *testing.T) {
	_, err := malgo.InitContext([]malgo.Backend{malgo.Backend(99)}, malgo.ContextConfig{}, nil)
	if err == nil {
		t.Fatalf("context init with invalid backend")
	}

	ctx, err := malgo.InitContext([]malgo.Backend{malgo.BackendNull}, malgo.ContextConfig{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	onSendFrames := func(outputSamples, inputSamples []byte, framecount uint32) {
	}

	deviceConfig := malgo.DefaultDeviceConfig()
	deviceConfig.DeviceType = malgo.Playback
	deviceConfig.Playback.Format = malgo.FormatType(99)
	deviceConfig.Playback.Channels = 99
	deviceConfig.SampleRate = 44100

	_, err = malgo.InitDevice(ctx.Context, deviceConfig, malgo.DeviceCallbacks{})
	if err == nil {
		t.Fatalf("device init with invalid config")
	}

	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = 44100

	dev, err := malgo.InitDevice(ctx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: onSendFrames,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dev.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = dev.Start()
	if err == nil {
		t.Fatalf("device start but already started")
	}

	time.Sleep(1 * time.Second)

	err = dev.Stop()
	if err != nil {
		t.Fatal(err)
	}

	err = dev.Stop()
	if err == nil {
		t.Fatalf("device stop but already stopped")
	}

	dev.Uninit()
}
