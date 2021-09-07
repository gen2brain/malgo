package io_api

import "github.com/gen2brain/malgo"

// StreamConfig describes the parameters for an audio stream.
// Default values will pick the defaults of the default device.
type StreamConfig struct {
	Format     malgo.Format
	Channels   int
	SampleRate int
}

func (config StreamConfig) asDeviceConfig(deviceType malgo.DeviceType) malgo.DeviceConfig {
	deviceConfig := malgo.NewDeviceConfig(deviceType)
	if config.Format != malgo.FormatUnknown {
		deviceConfig.Capture.Format = config.Format
		deviceConfig.Playback.Format = config.Format
	}
	if config.Channels != 0 {
		deviceConfig.Capture.Channels = config.Channels
		deviceConfig.Playback.Channels = config.Channels
	}
	if config.SampleRate != 0 {
		deviceConfig.SampleRate = config.SampleRate
	}
	return deviceConfig
}
