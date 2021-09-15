package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func main() {
	context, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	// Playback devices.
	infos, err := context.Devices(malgo.Playback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playback Devices")
	for i, info := range infos {
		e := "ok"
		full, err := context.DeviceInfo(malgo.Playback, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}
		fmt.Printf("    %d: %v, %s, [%s], channels: %d-%d, samplerate: %d-%d\n",
			i, info.ID, info.Name, e, full.MinChannels, full.MaxChannels, full.MinSampleRate, full.MaxSampleRate)
	}

	fmt.Println()

	// Capture devices.
	infos, err = context.Devices(malgo.Capture)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Capture Devices")
	for i, info := range infos {
		e := "ok"
		full, err := context.DeviceInfo(malgo.Capture, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}
		fmt.Printf("    %d: %v, %s, [%s], channels: %d-%d, samplerate: %d-%d\n",
			i, info.ID, info.Name, e, full.MinChannels, full.MaxChannels, full.MinSampleRate, full.MaxSampleRate)
	}
}
