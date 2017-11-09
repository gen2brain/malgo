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

	config := device.ConfigInit(FormatS16, 2, 48000, onRecvFrames, onSendFrames)

	err = device.Init(Capture, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	device.Uninit()

	err = device.Init(Playback, nil, &config)
	if err != nil {
		t.Fatal(err)
	}

	err = device.Start()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	device.Uninit()
}
