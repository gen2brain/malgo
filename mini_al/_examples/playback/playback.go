package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"

	"github.com/gen2brain/malgo/mini_al"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input file.")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32

	switch strings.ToLower(filepath.Ext(os.Args[1])) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())
	default:
		fmt.Println("Not a valid file.")
		os.Exit(1)
	}

	device := mini_al.NewDevice()

	// This is the function that's used for sending more data to the device for playback.
	onSendSamples := func(framecount uint32, psamples []byte) uint32 {
		n, err := reader.Read(psamples)
		if err == io.EOF {
			return 0
		}

		return uint32(n) / device.Channels() / device.SampleSizeInBytes(device.Format())
	}

	err = device.ContextInit(nil, mini_al.ContextConfig{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer device.ContextUninit()

	config := device.ConfigInitPlayback(mini_al.FormatS16, channels, sampleRate, onSendSamples)
	config.Alsa.NoMMap = 1

	err = device.Init(mini_al.Playback, nil, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer device.Uninit()

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	os.Exit(0)
}
