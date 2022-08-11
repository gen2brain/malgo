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

// caller must ma_free
func (d *DeviceConfig) cptrClone() (*C.ma_device_config, error) {
	deviceConfigPtr := C.ma_malloc(C.sizeof_ma_device_config, nil)
	if uintptr(deviceConfigPtr) == uintptr(0) {
		return nil, ErrOutOfMemory
	}
	deviceConfig := (*C.ma_device_config)(deviceConfigPtr)

	deviceConfig.deviceType = C.ma_device_type(d.DeviceType)
	deviceConfig.sampleRate = C.uint(d.SampleRate)
	deviceConfig.periodSizeInFrames = C.uint(d.PeriodSizeInFrames)
	deviceConfig.periodSizeInMilliseconds = C.uint(d.PeriodSizeInMilliseconds)
	deviceConfig.periods = C.uint(d.Periods)
	deviceConfig.performanceProfile = C.ma_performance_profile(d.PerformanceProfile)
	deviceConfig.noPreZeroedOutputBuffer = C.uchar(d.NoPreZeroedOutputBuffer)
	deviceConfig.noClip = C.uchar(d.NoClip)
	deviceConfig.dataCallback = d.DataCallback
	deviceConfig.stopCallback = d.StopCallback
	deviceConfig.pUserData = unsafe.Pointer(d.PUserData)
	deviceConfig.resampling.algorithm = C.ma_resample_algorithm(d.Resampling.Algorithm)
	deviceConfig.resampling.linear.lpfOrder = C.uint(d.Resampling.Linear.LpfOrder)
	deviceConfig.resampling.speex.quality = C.int(d.Resampling.Speex.Quality)
	convertSubConfig := func(to *C.struct___95, from SubConfig) {
		to.pDeviceID = (*C.ma_device_id)(from.DeviceID)
		to.format = C.ma_format(from.Format)
		to.channels = C.uint(from.Channels)
		for i := 0; i < len(to.channelMap); i++ {
			to.channelMap[i] = (C.uchar)(from.ChannelMap[i])
		}
		to.shareMode = C.ma_share_mode(from.ShareMode)
	}
	convertSubConfig(&deviceConfig.playback, d.Playback)
	convertSubConfig(&deviceConfig.capture, d.Capture)
	deviceConfig.wasapi.noAutoConvertSRC = C.uchar(d.Wasapi.NoHardwareOffloading)
	deviceConfig.wasapi.noDefaultQualitySRC = C.uchar(d.Wasapi.NoDefaultQualitySRC)
	deviceConfig.wasapi.noAutoStreamRouting = C.uchar(d.Wasapi.NoAutoStreamRouting)
	deviceConfig.wasapi.noHardwareOffloading = C.uchar(d.Wasapi.NoHardwareOffloading)
	deviceConfig.alsa.noMMap = C.uint(d.Alsa.NoMMap)
	deviceConfig.alsa.noAutoFormat = C.uint(d.Alsa.NoAutoFormat)
	deviceConfig.alsa.noAutoChannels = C.uint(d.Alsa.NoAutoChannles)
	deviceConfig.alsa.noAutoResample = C.uint(d.Alsa.NoAutoResample)
	deviceConfig.pulse.pStreamNameCapture = (*C.char)(d.Pulse.StreamNameCapture)
	deviceConfig.pulse.pStreamNamePlayback = (*C.char)(d.Pulse.StreamNamePlayback)

	return deviceConfig, nil
}

// SubConfig type.
type SubConfig struct {
	DeviceID   unsafe.Pointer
	Format     FormatType
	Channels   uint32
	ChannelMap [C.MA_MAX_CHANNELS]uint8
	ShareMode  ShareMode
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

// ResampleConfig type.
type ResampleConfig struct {
	Algorithm ResampleAlgorithm
	Linear    ResampleLinearConfig
	Speex     ResampleSpeexConfig
}

// ResampleLinearConfig type.
type ResampleLinearConfig struct {
	LpfOrder uint32
}

// ResampleSpeexConfig type.
type ResampleSpeexConfig struct {
	Quality int
}
