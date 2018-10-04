package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo/mini_al"
)

func main() {
	context, err := mini_al.InitContext(nil, mini_al.ContextConfig{}, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	// Playback devices.
	infos, err := context.Devices(mini_al.Playback)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playback Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %v, %s\n", i, info.ID, info.Name())
	}

	fmt.Println()

	// Capture devices.
	infos, err = context.Devices(mini_al.Capture)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Capture Devices")
	for i, info := range infos {
		fmt.Printf("    %d: %v, %s\n", i, info.ID, info.Name())
	}
}
