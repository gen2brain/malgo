package malgo

// #include "malgo.h"
import "C"
import (
	"unsafe"
)

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

	var deviceConfig DeviceConfig

	deviceConfig.DeviceType = DeviceType(config.deviceType)
	deviceConfig.SampleRate = uint32(config.sampleRate)
	deviceConfig.PeriodSizeInFrames = uint32(config.periodSizeInFrames)
	deviceConfig.PeriodSizeInMilliseconds = uint32(config.periodSizeInMilliseconds)
	deviceConfig.Periods = uint32(config.periodSizeInFrames)
	deviceConfig.PerformanceProfile = PerformanceProfile(config.performanceProfile)
	deviceConfig.NoPreZeroedOutputBuffer = uint32(config.noPreZeroedOutputBuffer)
	deviceConfig.NoClip = uint32(config.noClip)
	deviceConfig.DataCallback = config.dataCallback
	deviceConfig.StopCallback = config.stopCallback
	deviceConfig.PUserData = (*byte)(config.pUserData)

	deviceConfig.Resampling.Algorithm = ResampleAlgorithm(config.resampling.algorithm)
	deviceConfig.Resampling.Linear.LpfOrder = uint32(config.resampling.linear.lpfOrder)
	deviceConfig.Resampling.Speex.Quality = int(config.resampling.speex.quality)

	deviceConfig.Playback.DeviceID = unsafe.Pointer(config.playback.pDeviceID)
	deviceConfig.Playback.Format = FormatType(config.playback.format)
	deviceConfig.Playback.Channels = uint32(config.playback.channels)
	for i := 0; i < len(config.playback.channelMap); i++ {
		deviceConfig.Playback.ChannelMap[i] = uint8(config.playback.channelMap[i])
	}
	deviceConfig.Playback.ShareMode = ShareMode(config.playback.shareMode)

	deviceConfig.Capture.DeviceID = unsafe.Pointer(config.capture.pDeviceID)
	deviceConfig.Capture.Format = FormatType(config.capture.format)
	deviceConfig.Capture.Channels = uint32(config.capture.channels)
	for i := 0; i < len(config.capture.channelMap); i++ {
		deviceConfig.Capture.ChannelMap[i] = uint8(config.capture.channelMap[i])
	}
	deviceConfig.Capture.ShareMode = ShareMode(config.capture.shareMode)

	deviceConfig.Wasapi.NoAutoConvertSRC = uint32(config.wasapi.noHardwareOffloading)
	deviceConfig.Wasapi.NoDefaultQualitySRC = uint32(config.wasapi.noDefaultQualitySRC)
	deviceConfig.Wasapi.NoAutoStreamRouting = uint32(config.wasapi.noAutoStreamRouting)
	deviceConfig.Wasapi.NoHardwareOffloading = uint32(config.wasapi.noHardwareOffloading)

	deviceConfig.Alsa.NoMMap = uint32(config.alsa.noMMap)
	deviceConfig.Alsa.NoAutoFormat = uint32(config.alsa.noAutoFormat)
	deviceConfig.Alsa.NoAutoChannels = uint32(config.alsa.noAutoChannels)
	deviceConfig.Alsa.NoAutoResample = uint32(config.alsa.noAutoResample)

	if config.pulse.pStreamNameCapture != nil {
		deviceConfig.Pulse.StreamNameCapture = C.GoString(config.pulse.pStreamNameCapture)
	}
	if config.pulse.pStreamNamePlayback != nil {
		deviceConfig.Pulse.StreamNamePlayback = C.GoString(config.pulse.pStreamNamePlayback)
	}

	return deviceConfig
}

func (d *DeviceConfig) toC() (C.ma_device_config, func()) {
	deviceConfig := C.ma_device_config_init(C.ma_device_type(d.DeviceType))

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

	deviceConfig.playback.pDeviceID = (*C.ma_device_id)(d.Playback.DeviceID)
	deviceConfig.playback.format = C.ma_format(d.Playback.Format)
	deviceConfig.playback.channels = C.uint(d.Playback.Channels)
	for i := 0; i < len(deviceConfig.playback.channelMap); i++ {
		deviceConfig.playback.channelMap[i] = (C.uchar)(d.Playback.ChannelMap[i])
	}
	deviceConfig.playback.shareMode = C.ma_share_mode(d.Playback.ShareMode)

	deviceConfig.capture.pDeviceID = (*C.ma_device_id)(d.Capture.DeviceID)
	deviceConfig.capture.format = C.ma_format(d.Capture.Format)
	deviceConfig.capture.channels = C.uint(d.Capture.Channels)
	for i := 0; i < len(deviceConfig.capture.channelMap); i++ {
		deviceConfig.capture.channelMap[i] = (C.uchar)(d.Capture.ChannelMap[i])
	}
	deviceConfig.capture.shareMode = C.ma_share_mode(d.Capture.ShareMode)

	deviceConfig.wasapi.noAutoConvertSRC = C.uchar(d.Wasapi.NoHardwareOffloading)
	deviceConfig.wasapi.noDefaultQualitySRC = C.uchar(d.Wasapi.NoDefaultQualitySRC)
	deviceConfig.wasapi.noAutoStreamRouting = C.uchar(d.Wasapi.NoAutoStreamRouting)
	deviceConfig.wasapi.noHardwareOffloading = C.uchar(d.Wasapi.NoHardwareOffloading)

	deviceConfig.alsa.noMMap = C.uint(d.Alsa.NoMMap)
	deviceConfig.alsa.noAutoFormat = C.uint(d.Alsa.NoAutoFormat)
	deviceConfig.alsa.noAutoChannels = C.uint(d.Alsa.NoAutoChannels)
	deviceConfig.alsa.noAutoResample = C.uint(d.Alsa.NoAutoResample)

	var releasers []func()
	if d.Pulse.StreamNameCapture != "" {
		streamNameCapturePtr := C.CString(d.Pulse.StreamNameCapture)
		deviceConfig.pulse.pStreamNameCapture = streamNameCapturePtr
		releasers = append(releasers, func() {
			C.ma_free(unsafe.Pointer(streamNameCapturePtr), nil)
		})
	}
	if d.Pulse.StreamNamePlayback != "" {
		streamNamePlaybackPtr := C.CString(d.Pulse.StreamNamePlayback)
		deviceConfig.pulse.pStreamNamePlayback = streamNamePlaybackPtr
		releasers = append(releasers, func() {
			C.ma_free(unsafe.Pointer(streamNamePlaybackPtr), nil)
		})
	}

	return deviceConfig, func() {
		for _, release := range releasers {
			defer release()
		}
	}
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
	NoAutoChannels uint32
	NoAutoResample uint32
}

// PulseDeviceConfig type.
type PulseDeviceConfig struct {
	StreamNamePlayback string
	StreamNameCapture  string
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
