package malgo

// #include "malgo.h"
import "C"
import "unsafe"

// WasapiDeviceConfig type.
type WasapiDeviceConfig struct {
	NoAutoConvertSRC	uint32
	NoDefaultQualitySRC	uint32
}

// AlsaDeviceConfig type.
type AlsaDeviceConfig struct {
	NoMMap uint32
}

// PulseDeviceConfig type.
type PulseDeviceConfig struct {
	StreamNamePlayback *int8
	StreamNameCapture  *int8
}

// SubConfig type.
type SubConfig struct {
	DeviceID   *DeviceID
	Format     FormatType
	Channels   uint32
	ChannelMap [C.MA_MAX_CHANNELS]uint8
	ShareMode  ShareMode
	CgoPadding [4]byte
}

// DeviceConfig type.
type DeviceConfig struct {
	DeviceType               DeviceType
	SampleRate               uint32
	BufferSizeInFrames       uint32
	BufferSizeInMilliseconds uint32
	Periods                  uint32
	PerformanceProfile       PerformanceProfile
	DataCallback			 *[0]byte
	StopCallback			 *[0]byte
	PUserData				 *byte
	Playback                 SubConfig
	Capture                  SubConfig
	Wasapi                   WasapiDeviceConfig
	Alsa                     AlsaDeviceConfig
	Pulse                    PulseDeviceConfig
}

func (d *DeviceConfig) cptr() *C.ma_device_config {
	return (*C.ma_device_config)(unsafe.Pointer(d))
}

// DefaultDeviceConfig returns a default device config.
func DefaultDeviceConfig() DeviceConfig {
	return DeviceConfig{}
}
