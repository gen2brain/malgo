package malgo

// #include "malgo.h"
import "C"
import "unsafe"

// DeviceConfig type.
type DeviceConfig struct {
	DeviceType               DeviceType
	SampleRate               uint32
	PeriodSizeInFrames       uint32
	PeriodSizeInMilliseconds uint32
	Periods                  uint32
	PerformanceProfile       PerformanceProfile
	NoPreZeroedOutputBuffer  uint32
	NoClip                   uint32
	DataCallback             *[0]byte
	StopCallback             *[0]byte
	PUserData                *byte
	Resampling               ResampleConfig
	Playback                 SubConfig
	Capture                  SubConfig
	Wasapi                   WasapiDeviceConfig
	Alsa                     AlsaDeviceConfig
	Pulse                    PulseDeviceConfig
}

// DefaultDeviceConfig returns a default device config.
func DefaultDeviceConfig(deviceType DeviceType) DeviceConfig {
	config := C.ma_device_config_init(C.ma_device_type(deviceType))
	return *(*DeviceConfig)(unsafe.Pointer(&config))
}

func (d *DeviceConfig) cptr() *C.ma_device_config {
	return (*C.ma_device_config)(unsafe.Pointer(d))
}

// SubConfig type.
type SubConfig struct {
	DeviceID   *DeviceID
	Format     FormatType
	Channels   uint32
	ChannelMap [C.MA_MAX_CHANNELS]uint8
	ShareMode  ShareMode
	_          [4]byte // cgo padding
}

// WasapiDeviceConfig type.
type WasapiDeviceConfig struct {
	NoAutoConvertSRC     uint32
	NoDefaultQualitySRC  uint32
	NoAutoStreamRouting  uint32
	NoHardwareOffloading uint32
}

// AlsaDeviceConfig type.
type AlsaDeviceConfig struct {
	NoMMap         uint32
	NoAutoFormat   uint32
	NoAutoChannles uint32
	NoAutoResample uint32
}

// PulseDeviceConfig type.
type PulseDeviceConfig struct {
	StreamNamePlayback *int8
	StreamNameCapture  *int8
}

type ResampleConfig struct {
	Algorithm ResampleAlgorithm
	Linear    ResampleLinearConfig
	Speex     ResampleSpeexConfig
}

type ResampleLinearConfig struct {
	LpfOrder uint32
}

type ResampleSpeexConfig struct {
	Quality int
}
