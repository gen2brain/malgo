package malgo_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/gen2brain/malgo"
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

	ctx, err := malgo.InitContext([]malgo.Backend{malgo.BackendNull}, malgo.ContextConfig{}, onLog)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig()
	deviceConfig.Format = malgo.FormatS16
	deviceConfig.Channels = 2
	deviceConfig.SampleRate = 48000
	deviceConfig.Alsa.NoMMap = 1

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Format))
	onRecvFrames := func(framecount uint32, pSamples []byte) {
		sampleCount := framecount * deviceConfig.Channels * sizeInBytes

		newCapturedSampleCount := capturedSampleCount + sampleCount

		pCapturedSamples = append(pCapturedSamples, pSamples...)

		capturedSampleCount = newCapturedSampleCount
	}

	captureCallbacks := malgo.DeviceCallbacks{
		Recv: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, malgo.Capture, nil, deviceConfig, captureCallbacks)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != malgo.Capture {
		t.Errorf("wrong device type")
	}

	if device.Format() != malgo.FormatS16 {
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

	playbackCallbacks := malgo.DeviceCallbacks{
		Send: onSendFrames,
	}

	device, err = malgo.InitDevice(ctx.Context, malgo.Playback, nil, deviceConfig, playbackCallbacks)
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

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		return 0
	}

	deviceConfig := malgo.DefaultDeviceConfig()
	deviceConfig.Format = malgo.FormatType(99)
	deviceConfig.Channels = 99
	deviceConfig.SampleRate = 48000

	_, err = malgo.InitDevice(ctx.Context, malgo.Playback, nil, deviceConfig, malgo.DeviceCallbacks{})
	if err == nil {
		t.Fatalf("device init with invalid config")
	}

	deviceConfig.Format = malgo.FormatS16
	deviceConfig.Channels = 2
	deviceConfig.SampleRate = 48000

	dev, err := malgo.InitDevice(ctx.Context, malgo.Playback, nil, deviceConfig, malgo.DeviceCallbacks{
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
