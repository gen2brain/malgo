package io_api

import "github.com/gen2brain/malgo"

// StreamConfig describes the parameters for an audio stream.
// Default values will pick the defaults of the default device.
type StreamConfig struct {
	Format     malgo.FormatType
	Channels   int
	SampleRate int
}

func (config StreamConfig) asDeviceConfig() malgo.DeviceConfig {
	deviceConfig := malgo.DefaultDeviceConfig()
	if config.Format != malgo.FormatUnknown {
		deviceConfig.Format = config.Format
	}
	if config.Channels != 0 {
		deviceConfig.Channels = uint32(config.Channels)
	}
	if config.SampleRate != 0 {
		deviceConfig.SampleRate = uint32(config.SampleRate)
	}
	return deviceConfig
}
