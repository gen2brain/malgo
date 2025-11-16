package malgo

// #include "malgo.h"
import "C"
import (
	"unsafe"
)

// DeviceConfig type.
type DeviceConfig struct {
	DeviceType                DeviceType
	SampleRate                uint32
	PeriodSizeInFrames        uint32
	PeriodSizeInMilliseconds  uint32
	Periods                   uint32
	PerformanceProfile        PerformanceProfile
	NoPreSilencedOutputBuffer uint32
	NoClip                    uint32
	NoDisableDenormals        uint32
	NoFixedSizedCallback      uint32
	DataCallback              *[0]byte
	NotificationCallback      *[0]byte
	StopCallback              *[0]byte
	PUserData                 *byte
	Resampling                ResampleConfig
	Playback                  SubConfig
	Capture                   SubConfig
	Wasapi                    WasapiDeviceConfig
	Alsa                      AlsaDeviceConfig
	Pulse                     PulseDeviceConfig
	AAudio                    AAudioDeviceConfig
	// TODO: Add support for coreaudio, opensl, aaudio
}

// DefaultDeviceConfig returns a default device config.
func DefaultDeviceConfig(deviceType DeviceType) DeviceConfig {
	config := C.ma_device_config_init(C.ma_device_type(deviceType))

	var deviceConfig DeviceConfig

	deviceConfig.DeviceType = DeviceType(config.deviceType)
	deviceConfig.SampleRate = uint32(config.sampleRate)
	deviceConfig.PeriodSizeInFrames = uint32(config.periodSizeInFrames)
	deviceConfig.PeriodSizeInMilliseconds = uint32(config.periodSizeInMilliseconds)
	deviceConfig.Periods = uint32(config.periods)
	deviceConfig.PerformanceProfile = PerformanceProfile(config.performanceProfile)
	deviceConfig.NoPreSilencedOutputBuffer = uint32(config.noPreSilencedOutputBuffer)
	deviceConfig.NoClip = uint32(config.noClip)
	deviceConfig.NoDisableDenormals = uint32(config.noDisableDenormals)
	deviceConfig.NoFixedSizedCallback = uint32(config.noFixedSizedCallback)
	deviceConfig.DataCallback = config.dataCallback
	deviceConfig.NotificationCallback = config.notificationCallback
	deviceConfig.StopCallback = config.stopCallback
	deviceConfig.PUserData = (*byte)(config.pUserData)

	deviceConfig.Resampling.Algorithm = ResampleAlgorithm(config.resampling.algorithm)
	deviceConfig.Resampling.Linear.LpfOrder = uint32(config.resampling.linear.lpfOrder)

	deviceConfig.Playback.DeviceID = unsafe.Pointer(config.playback.pDeviceID)
	deviceConfig.Playback.Format = FormatType(config.playback.format)
	deviceConfig.Playback.Channels = uint32(config.playback.channels)
	deviceConfig.Playback.ChannelMap = unsafe.Pointer(config.playback.pChannelMap)
	deviceConfig.Playback.ShareMode = ShareMode(config.playback.shareMode)

	deviceConfig.Capture.DeviceID = unsafe.Pointer(config.capture.pDeviceID)
	deviceConfig.Capture.Format = FormatType(config.capture.format)
	deviceConfig.Capture.Channels = uint32(config.capture.channels)
	deviceConfig.Capture.ChannelMap = unsafe.Pointer(config.capture.pChannelMap)
	deviceConfig.Capture.ShareMode = ShareMode(config.capture.shareMode)

	deviceConfig.Wasapi.NoAutoConvertSRC = uint32(config.wasapi.noAutoConvertSRC)
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

	deviceConfig.AAudio.Usage = AAudioUsage(config.aaudio.usage)
	deviceConfig.AAudio.ContentType = AAudioContentType(config.aaudio.contentType)
	deviceConfig.AAudio.InputPreset = AAudioInputPreset(config.aaudio.inputPreset)
	deviceConfig.AAudio.AllowedCapturePolicy = AAudioAllowedCapturePolicy(config.aaudio.allowedCapturePolicy)
	deviceConfig.AAudio.NoAutoStartAfterReroute = uint32(config.aaudio.noAutoStartAfterReroute)
	deviceConfig.AAudio.EnableCompatibilityWorkarounds = uint32(config.aaudio.enableCompatibilityWorkarounds)

	return deviceConfig
}

func (d *DeviceConfig) toC() (C.ma_device_config, func()) {
	deviceConfig := C.ma_device_config_init(C.ma_device_type(d.DeviceType))

	deviceConfig.sampleRate = C.uint(d.SampleRate)
	deviceConfig.periodSizeInFrames = C.uint(d.PeriodSizeInFrames)
	deviceConfig.periodSizeInMilliseconds = C.uint(d.PeriodSizeInMilliseconds)
	deviceConfig.periods = C.uint(d.Periods)
	deviceConfig.performanceProfile = C.ma_performance_profile(d.PerformanceProfile)
	deviceConfig.noPreSilencedOutputBuffer = C.uchar(d.NoPreSilencedOutputBuffer)
	deviceConfig.noClip = C.uchar(d.NoClip)
	deviceConfig.noDisableDenormals = C.uchar(d.NoDisableDenormals)
	deviceConfig.noFixedSizedCallback = C.uchar(d.NoFixedSizedCallback)

	deviceConfig.dataCallback = d.DataCallback
	deviceConfig.notificationCallback = d.NotificationCallback
	deviceConfig.stopCallback = d.StopCallback
	deviceConfig.pUserData = unsafe.Pointer(d.PUserData)

	deviceConfig.resampling.algorithm = C.ma_resample_algorithm(d.Resampling.Algorithm)
	deviceConfig.resampling.linear.lpfOrder = C.uint(d.Resampling.Linear.LpfOrder)

	deviceConfig.playback.pDeviceID = (*C.ma_device_id)(d.Playback.DeviceID)
	deviceConfig.playback.format = C.ma_format(d.Playback.Format)
	deviceConfig.playback.channels = C.uint(d.Playback.Channels)
	deviceConfig.playback.pChannelMap = (*C.ma_channel)(d.Playback.ChannelMap)
	deviceConfig.playback.shareMode = C.ma_share_mode(d.Playback.ShareMode)

	deviceConfig.capture.pDeviceID = (*C.ma_device_id)(d.Capture.DeviceID)
	deviceConfig.capture.format = C.ma_format(d.Capture.Format)
	deviceConfig.capture.channels = C.uint(d.Capture.Channels)
	deviceConfig.capture.pChannelMap = (*C.ma_channel)(d.Capture.ChannelMap)
	deviceConfig.capture.shareMode = C.ma_share_mode(d.Capture.ShareMode)

	deviceConfig.wasapi.noAutoConvertSRC = C.uchar(d.Wasapi.NoAutoConvertSRC)
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

	deviceConfig.aaudio.usage = C.ma_aaudio_usage(d.AAudio.Usage)
	deviceConfig.aaudio.contentType = C.ma_aaudio_content_type(d.AAudio.ContentType)
	deviceConfig.aaudio.inputPreset = C.ma_aaudio_input_preset(d.AAudio.InputPreset)
	deviceConfig.aaudio.allowedCapturePolicy = C.ma_aaudio_allowed_capture_policy(d.AAudio.AllowedCapturePolicy)
	deviceConfig.aaudio.noAutoStartAfterReroute = C.ma_bool32(d.AAudio.NoAutoStartAfterReroute)
	deviceConfig.aaudio.enableCompatibilityWorkarounds = C.ma_bool32(d.AAudio.EnableCompatibilityWorkarounds)

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
	ChannelMap unsafe.Pointer
	ShareMode  ShareMode

	// Unexposed: channelMixMode, calculateLFEFromSpatialChannels
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

	// Unexposed: format, channels, sampleRateIn, sampleRateOut, pBackendVTable, pBackendUserData
}

// ResampleLinearConfig type.
type ResampleLinearConfig struct {
	LpfOrder uint32
}

// ResampleSpeexConfig type.
type ResampleSpeexConfig struct {
	Quality int
}

// AAudioDeviceConfig type.
type AAudioDeviceConfig struct {
	Usage                          AAudioUsage
	ContentType                    AAudioContentType
	InputPreset                    AAudioInputPreset
	AllowedCapturePolicy           AAudioAllowedCapturePolicy
	NoAutoStartAfterReroute        uint32
	EnableCompatibilityWorkarounds uint32
}
