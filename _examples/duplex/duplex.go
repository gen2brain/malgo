// This example simply captures data from your default microphone until you press Enter, after which it plays back the captured audio.
package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func main() {

	deviceConfig := malgo.NewDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatF32
	deviceConfig.Capture.Channels = 2
	deviceConfig.Playback.Format = malgo.FormatF32
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = true

	deviceConfig.DataCallback = func(_ malgo.Device, output, input *malgo.DataBuffer, framecount int) {
	copy(output.Bytes(), input.Bytes())
	}
	device, err := malgo.InitDevice(malgo.DefaultContext, deviceConfig)
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
}
