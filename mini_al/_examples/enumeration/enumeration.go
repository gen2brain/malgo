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
	infos, err := device.Devices(mal.Playback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playback Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %s, %s\n", i, info.ID, info.Name)
	}

	fmt.Println()

	// Capture devices.
	infos, err = device.Devices(mal.Capture)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Capture Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %s, %s\n", i, info.ID, info.Name)
	}
}
