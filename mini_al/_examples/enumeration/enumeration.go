package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo/mini_al"
)

func main() {
	device := mini_al.NewDevice()

	err := device.ContextInit(nil, mini_al.ContextConfig{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer device.ContextUninit()

	// Playback devices.
	infos, err := device.Devices(mini_al.Playback)
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
	infos, err = device.Devices(mini_al.Capture)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Capture Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %s, %s\n", i, info.ID, info.Name)
	}
}
