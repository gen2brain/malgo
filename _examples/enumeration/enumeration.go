package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func main() {
	context, err := malgo.InitContext(nil, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		 context.Uninit()
	}()
	fmt.Printf("Using backend %v, loopback supported %v\n", context.Backend(), context.IsLoopbackSupported())
	// Playback devices.
	playback, capture, err := context.AllDevices()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playback Devices")
	for i, info := range playback {
		e := "ok"
		full, err := context.DeviceInfo(malgo.Playback, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}
		fmt.Printf("    %d: %s, [%s], channels: %d-%d, samplerate: %d-%d, %v\n",
			i, info.Name, e, full.MinChannels, full.MaxChannels, full.MinSampleRate, full.MaxSampleRate, full.Formats)
	}

	fmt.Println()

	// Capture devices.
	fmt.Println("Capture Devices")
	for i, info := range capture {
		e := "ok"
		full, err := context.DeviceInfo(malgo.Capture, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}
		fmt.Printf("    %d: %s, [%s], channels: %d-%d, samplerate: %d-%d, formats: %v\n",
			i, info.Name, e, full.MinChannels, full.MaxChannels, full.MinSampleRate, full.MaxSampleRate, full.Formats)
	}
}
