package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func main() {
	device := mal.NewDevice()

	err := device.ContextInit(nil, mal.ContextConfig{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer device.ContextUninit()

	// Playback devices.
	infos, err := device.EnumerateDevices(mal.Playback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playback Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %s\n", i, info)
	}

	fmt.Println()

	// Capture devices.
	infos, err = device.EnumerateDevices(mal.Capture)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Capture Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %s\n", i, info)
	}
}
