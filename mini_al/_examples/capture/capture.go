// This example simply captures data from your default microphone until you press Enter, after which it plays back the captured audio.
package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func main() {
	device := mal.NewDevice()

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

	err := device.ContextInit(nil, mal.ContextConfig{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer device.ContextUninit()

	config := device.ConfigInit(mal.FormatS16, 2, 48000, onRecvFrames, onSendFrames)
	config.Alsa.NoMMap = 1

	fmt.Println("Recording...")
	err = device.Init(mal.Capture, nil, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to stop recording...")
	fmt.Scanln()

	device.Uninit()

	fmt.Println("Playing...")
	err = device.Init(mal.Playback, nil, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	device.Uninit()

	os.Exit(0)
}
