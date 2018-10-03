package mini_al

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestCapturePlayback(t *testing.T) {
	device := NewDevice()

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	onRecvFrames := func(framecount uint32, pSamples []byte) {
		sizeInBytes := device.SampleSizeInBytes(device.Format())
		sampleCount := framecount * device.Channels() * sizeInBytes

		newCapturedSampleCount := capturedSampleCount + sampleCount

		pCapturedSamples = append(pCapturedSamples, pSamples...)

		capturedSampleCount = newCapturedSampleCount
	}

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		sizeInBytes := device.SampleSizeInBytes(device.Format())
		samplesToRead := framecount * device.Channels() * sizeInBytes
		if samplesToRead > capturedSampleCount-playbackSampleCount {
			samplesToRead = capturedSampleCount - playbackSampleCount
		}

		copy(pSamples, pCapturedSamples[playbackSampleCount:playbackSampleCount+samplesToRead])

		playbackSampleCount += samplesToRead

		return samplesToRead / device.Channels() / sizeInBytes
	}

	onLog := func(message string) {
		fmt.Fprintf(ioutil.Discard, message)
	}

	contextConfig := device.ContextConfigInit(onLog)

	err := device.ContextInit([]Backend{BackendNull}, contextConfig)
	if err != nil {
		t.Fatal(err)
	}

	defer device.ContextUninit()

	config := device.ConfigInitCapture(FormatS16, 2, 48000, onRecvFrames)

	err = device.Init(Capture, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != Capture {
		t.Errorf("wrong device type")
	}

	if device.Format() != FormatS16 {
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

	config = device.ConfigInitPlayback(FormatS16, 2, 48000, onSendFrames)

	err = device.Init(Playback, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	if device.Type() != Playback {
		t.Errorf("wrong device type")
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	device.Uninit()
}

func TestDevices(t *testing.T) {
	device := NewDevice()

	config := device.ContextConfigInit(nil)
	config.Alsa.UseVerboseDeviceEnumeration = 1

	err := device.ContextInit([]Backend{BackendNull}, ContextConfig{})
	if err != nil {
		t.Fatal(err)
	}

	defer device.ContextUninit()

	infosPlayback, err := device.Devices(Playback)
	if err != nil {
		t.Error(err)
	}

	if len(infosPlayback) == 0 {
		t.Errorf("empty playback device info")
	}

	for _, i := range infosPlayback {
		fmt.Fprintf(ioutil.Discard, i.String())
	}

	infosCapture, err := device.Devices(Capture)
	if err != nil {
		t.Error(err)
	}

	if len(infosCapture) == 0 {
		t.Errorf("empty capture device info")
	}
}

func TestConfigInit(t *testing.T) {
	device := NewDevice()

	onRecvFrames := func(framecount uint32, pSamples []byte) {
	}

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		return 0
	}

	onStop := func() {
	}

	err := device.ContextInit([]Backend{BackendNull}, ContextConfig{})
	if err != nil {
		t.Fatal(err)
	}

	defer device.ContextUninit()

	config := device.ConfigInit(FormatS16, 2, 48000, onRecvFrames, onSendFrames)

	err = device.Init(Playback, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	device.SetStopCallback(onStop)

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	if !device.IsStarted() {
		t.Fatalf("device not started")
	}

	err = device.Stop()
	if err != nil {
		t.Fatal(err)
	}

	device.Uninit()
}

func TestErrors(t *testing.T) {
	device := NewDevice()

	err := device.ContextInit([]Backend{Backend(99)}, ContextConfig{})
	if err == nil {
		t.Fatalf("context init with invalid backend")
	}

	err = device.ContextInit([]Backend{BackendNull}, ContextConfig{})
	if err != nil {
		t.Fatal(err)
	}

	onSendFrames := func(framecount uint32, pSamples []byte) uint32 {
		return 0
	}

	config := device.ConfigInitPlayback(FormatType(99), 99, 48000, nil)

	err = device.Init(Playback, nil, &config)
	if err == nil {
		t.Fatalf("device init with invalid config")
	}

	config = device.ConfigInitPlayback(FormatS16, 2, 48000, onSendFrames)

	err = device.Init(Playback, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = device.Start()
	if err == nil {
		t.Fatalf("device start but already started")
	}

	time.Sleep(1 * time.Second)

	err = device.Stop()
	if err != nil {
		t.Fatal(err)
	}

	err = device.Stop()
	if err == nil {
		t.Fatalf("device stop but already stopped")
	}

	device.ContextUninit()

	device.Uninit()
}
