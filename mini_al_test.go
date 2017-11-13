package mal

import (
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

	err := device.ContextInit([]Backend{BackendNull}, true)
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

	device.Stop()
	device.Uninit()
}

func TestEnumerate(t *testing.T) {
	device := NewDevice()

	err := device.ContextInit([]Backend{BackendNull}, false)
	if err != nil {
		t.Fatal(err)
	}

	defer device.ContextUninit()

	infos, err := device.EnumerateDevices(Playback)
	if err != nil {
		t.Error(err)
	}

	if len(infos) == 0 {
		t.Errorf("empty playback device info")
	}

	infos, err = device.EnumerateDevices(Capture)
	if err != nil {
		t.Error(err)
	}

	if len(infos) == 0 {
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

	err := device.ContextInit([]Backend{BackendNull}, false)
	if err != nil {
		t.Fatal(err)
	}

	defer device.ContextUninit()

	config := device.ConfigInit(FormatS16, 2, 48000, nil, nil)

	err = device.Init(Capture, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	device.SetRecvCallback(onRecvFrames)
	device.SetSendCallback(onSendFrames)
	device.SetStopCallback(onStop)

	device.Uninit()
}
