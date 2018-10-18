package mini_al_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/gen2brain/malgo/mini_al"
)

var testenvWithHardware bool

func init() {
	flag.BoolVar(&testenvWithHardware, "malgo.hw", false, "Add flag to run tests expecting hardware")
	flag.Parse()
}

func TestCapturePlayback(t *testing.T) {
	onLog := func(message string) {
		fmt.Fprintf(ioutil.Discard, message)
	}

	ctx, err := mini_al.InitContext([]mini_al.Backend{mini_al.BackendNull}, mini_al.ContextConfig{}, onLog)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := mini_al.DefaultDeviceConfig()
	deviceConfig.Format = mini_al.FormatS16
	deviceConfig.Channels = 2
	deviceConfig.SampleRate = 48000
	deviceConfig.Alsa.NoMMap = 1

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	sizeInBytes := uint32(mini_al.SampleSizeInBytes(deviceConfig.Format))
	onRecvFrames := func(framecount uint32, pSamples []byte) {
		sampleCount := framecount * deviceConfig.Channels * sizeInBytes

		newCapturedSampleCount := capturedSampleCount + sampleCount

		pCapturedSamples = append(pCapturedSamples, pSamples...)

		capturedSampleCount = newCapturedSampleCount
	}

	captureCallbacks := mini_al.DeviceCallbacks{
		Recv: onRecvFrames,
	}
	device, err := mini_al.InitDevice(ctx.Context, mini_al.Capture, nil, deviceConfig, captureCallbacks)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != mini_al.Capture {
		t.Errorf("wrong device type")
	}

	if device.Format() != mini_al.FormatS16 {
		t.Errorf("wrong format")
	}

	if device.Channels() != 2 {
		t.Errorf("wrong number of channels")
	}

	if device.SampleRate() != 48000 {
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

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		samplesToRead := framecount * deviceConfig.Channels * sizeInBytes
		if samplesToRead > capturedSampleCount-playbackSampleCount {
			samplesToRead = capturedSampleCount - playbackSampleCount
		}

		copy(pSamples, pCapturedSamples[playbackSampleCount:playbackSampleCount+samplesToRead])

		playbackSampleCount += samplesToRead

		return samplesToRead / deviceConfig.Channels / sizeInBytes
	}

	playbackCallbacks := mini_al.DeviceCallbacks{
		Send: onSendFrames,
	}

	device, err = mini_al.InitDevice(ctx.Context, mini_al.Playback, nil, deviceConfig, playbackCallbacks)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != mini_al.Playback {
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

	_, err := mini_al.InitContext([]mini_al.Backend{mini_al.Backend(99)}, mini_al.ContextConfig{}, nil)
	if err == nil {
		t.Fatalf("context init with invalid backend")
	}

	ctx, err := mini_al.InitContext([]mini_al.Backend{mini_al.BackendNull}, mini_al.ContextConfig{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		return 0
	}

	deviceConfig := mini_al.DefaultDeviceConfig()
	deviceConfig.Format = mini_al.FormatType(99)
	deviceConfig.Channels = 99
	deviceConfig.SampleRate = 48000

	_, err = mini_al.InitDevice(ctx.Context, mini_al.Playback, nil, deviceConfig, mini_al.DeviceCallbacks{})
	if err == nil {
		t.Fatalf("device init with invalid config")
	}

	deviceConfig.Format = mini_al.FormatS16
	deviceConfig.Channels = 2
	deviceConfig.SampleRate = 48000

	dev, err := mini_al.InitDevice(ctx.Context, mini_al.Playback, nil, deviceConfig, mini_al.DeviceCallbacks{
		Send: onSendFrames,
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
