package malgo

import "github.com/gen2brain/malgo/mini_al"

// StreamConfig describes the parameters for an audio stream.
// Default values will pick the defaults of the default device.
type StreamConfig struct {
	Format     mini_al.FormatType
	Channels   int
	SampleRate int
}

func (config StreamConfig) asDeviceConfig() mini_al.DeviceConfig {
	deviceConfig := mini_al.DefaultDeviceConfig()
	if config.Format != mini_al.FormatUnknown {
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
