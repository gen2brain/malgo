// This example simply captures data from your default microphone until you press Enter, after which it plays back the captured audio.
package main

import (
	"fmt"
	"github.com/gen2brain/malgo"
	"bytes"
)

func main() {
	ctx, err := malgo.InitContext(nil, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		 ctx.Uninit()
	}()

	deviceConfig := malgo.NewDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatF32
	deviceConfig.Capture.Channels = 2
	deviceConfig.Playback.Format = malgo.FormatF32
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = true
	capturedSamples := make([]byte, 10000)
	deviceConfig.DataCallback = func(_ malgo.Device, pSample2, pSample *malgo.DataBuffer, framecount int) {
		capturedSamples = append(capturedSamples, pSample.Bytes()...)
	}
	fmt.Println("Recording...")
	device, err := malgo.InitDevice(ctx, deviceConfig)
	if err != nil {
		panic(err)
	}

	err = device.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println("Press Enter to stop recording...")
	fmt.Scanln()
	device.Uninit()
	r := bytes.NewReader(capturedSamples)
	deviceConfig.DataCallback = func(device malgo.Device, pSample, pInput *malgo.DataBuffer, frameCount int) {
		outputBytes := pSample.Bytes()
		amount, _ := r.Read(outputBytes)
		if amount < len(outputBytes) {
			// We didn't completely fill the buffer, which means the reader has run out of data. Loop by making a new reader that will start at the beginning again
			r = bytes.NewReader(capturedSamples)
			r.Read(outputBytes[amount:])
		}
	}
	fmt.Println("Playing...")
	deviceConfig.DeviceType = malgo.Duplex
	device, err = malgo.InitDevice(ctx, deviceConfig)
	if err != nil {
		panic(err)
	}
	err = device.Start()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Playing %v bytes of data at sample rate %v, channels %v, format %v, length %v\n", len(capturedSamples), device.SampleRate(), device.CaptureChannels(), device.CaptureFormat(), malgo.BytesToDuration(len(capturedSamples), device.SampleRate(), device.CaptureChannels(), device.CaptureFormat()))
	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	device.Uninit()
}
