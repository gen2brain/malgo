package mini_al

// #include "malgo.h"
import "C"
import "unsafe"

// AlsaDeviceConfig type.
type AlsaDeviceConfig struct {
	NoMMap uint32
}

// PulseDeviceConfig type.
type PulseDeviceConfig struct {
	StreamName *byte
}

// DeviceConfig type.
type DeviceConfig struct {
	Format                   FormatType
	Channels                 uint32
	SampleRate               uint32
	ChannelMap               [32]byte
	BufferSizeInFrames       uint32
	BufferSizeInMilliseconds uint32
	Periods                  uint32
	ShareMode                ShareMode
	PerformanceProfile       PerformanceProfile
	_                        uintptr
	_                        uintptr
	_                        uintptr
	Alsa                     AlsaDeviceConfig
	_                        [4]byte
	Pulse                    PulseDeviceConfig
}

func (d *DeviceConfig) cptr() *C.mal_device_config {
	return (*C.mal_device_config)(unsafe.Pointer(d))
}

// DefaultDeviceConfig returns a default device config.
func DefaultDeviceConfig() DeviceConfig {
	return DeviceConfig{}
}
