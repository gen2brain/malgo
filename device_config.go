package malgo

// #include "malgo.h"
import "C"
import (
	"unsafe"
	"time"
)

type DeviceConfig struct {
	DeviceType              DeviceType
	SampleRate              int
	PeriodSizeInFrames      int
	PeriodSizeInDuration    time.Duration
	Periods                 int
	PerformanceProfile      int
	NoPreZeroedOutputBuffer bool
	NoClip                  bool
	Playback                struct {
		DeviceID       *DeviceID
		Format         FormatType
		Channels       int
		ChannelMap     [C.MA_MAX_CHANNELS]uint8
		ShareMode      int
		ChannelMixMode int
	}
	Capture struct {
		DeviceID       *DeviceID
		Format         FormatType
		Channels       int
		ChannelMap     [C.MA_MAX_CHANNELS]uint8
		ShareMode      int
		ChannelMixMode int
	}
	Wasapi struct {
		NoAutoConvertSRC     bool
		NoDefaultQualitySRC  bool
		NoAutoStreamRouting  bool
		NoHardwareOffloading bool
	}
	Alsa struct {
		NoMMap         bool
		NoAutoFormat   bool
		NoAutoChannels bool
		NoAutoResample bool
	}
	Pulse struct {
		StreamNamePlayback string
		StreamNameCapture  string
	}
	CoreAudio struct {
		AllowNominalSampleRateChange bool
	}
	Opensl struct {
		StreamType      int
		RecordingPreset int
	}
	Aaudio struct {
		Usage       int
		ContentType int
		InputPreset int
	}
	Resampling struct {
		Algorithm int
		Linear    struct {
			LPFOrder int
		}
		Speex struct {
			Quality int
		}
	}
}
// DefaultDeviceConfig returns a default device config.
func DefaultDeviceConfig(deviceType DeviceType) DeviceConfig {
	cConfig := C.ma_device_config_init(C.ma_device_type(deviceType))
	config := DeviceConfig{}
	config.DeviceType = DeviceType(cConfig.deviceType)
	config.SampleRate = int(cConfig.sampleRate)
	config.PeriodSizeInFrames = int(cConfig.periodSizeInFrames)
	config.PeriodSizeInDuration = time.Millisecond * time.Duration(cConfig.periodSizeInMilliseconds)
	config.Periods = int(cConfig.periods)
	config.PerformanceProfile = int(cConfig.performanceProfile)
	config.NoPreZeroedOutputBuffer = int8ToBool(cConfig.noPreZeroedOutputBuffer)
	config.NoClip = int8ToBool(cConfig.noClip)
	config.Resampling.Algorithm = int(cConfig.resampling.algorithm)
	config.Resampling.Linear.LPFOrder = int(cConfig.resampling.linear.lpfOrder)
	config.Resampling.Speex.Quality = int(cConfig.resampling.speex.quality)
	config.Wasapi.NoAutoConvertSRC = int8ToBool(cConfig.wasapi.noAutoConvertSRC)
	config.Wasapi.NoDefaultQualitySRC = int8ToBool(cConfig.wasapi.noDefaultQualitySRC)
	config.Wasapi.NoAutoStreamRouting = int8ToBool(cConfig.wasapi.noAutoStreamRouting)
	config.Wasapi.NoHardwareOffloading = int8ToBool(cConfig.wasapi.noHardwareOffloading)
	config.Alsa.NoMMap = intToBool(cConfig.alsa.noMMap)
	config.Alsa.NoAutoFormat = intToBool(cConfig.alsa.noAutoFormat)
	config.Alsa.NoAutoChannels = intToBool(cConfig.alsa.noAutoChannels)
	config.Alsa.NoAutoResample = intToBool(cConfig.alsa.noAutoResample)
	config.Pulse.StreamNameCapture = goString(cConfig.pulse.pStreamNameCapture)
	config.Pulse.StreamNamePlayback = goString(cConfig.pulse.pStreamNamePlayback)
	config.CoreAudio.AllowNominalSampleRateChange = intToBool(cConfig.coreaudio.allowNominalSampleRateChange)
	config.Opensl.StreamType = int(cConfig.opensl.streamType)
	config.Opensl.RecordingPreset = int(cConfig.opensl.recordingPreset)
	config.Aaudio.ContentType = int(cConfig.aaudio.contentType)
	config.Aaudio.InputPreset = int(cConfig.aaudio.inputPreset)

	return config
}

func (config *DeviceConfig) toCRepr() (C.ma_device_config, pointerList) {
	// even if we forget to initialize some fields, telling Miniaudio to reinitialize the config ensures it doesn't break anything
	// a list of pointers that holds any memory allocated by the config like strings and deviceIDs ETC. This is returned alongside the new cConfig, is stored in the DeviceCallbacks struct and freed when Device.Uninit() is called
	var memory pointerList
	cConfig := C.ma_device_config_init(C.ma_device_type(config.DeviceType))
	cConfig.periodSizeInFrames = C.ma_uint32(config.PeriodSizeInFrames)
	cConfig.periodSizeInMilliseconds = C.ma_uint32(config.PeriodSizeInDuration.Milliseconds())
	cConfig.periods = C.ma_uint32(config.Periods)
	cConfig.performanceProfile = C.ma_performance_profile(config.PerformanceProfile)
	cConfig.noPreZeroedOutputBuffer = boolToInt8(config.NoPreZeroedOutputBuffer)
	cConfig.noClip = boolToInt8(config.NoClip)
	cConfig.resampling.algorithm = C.ma_resample_algorithm(config.Resampling.Algorithm)

	cConfig.resampling.linear.lpfOrder = C.ma_uint32(config.Resampling.Linear.LPFOrder)

	cConfig.resampling.speex.quality = C.int(config.Resampling.Speex.Quality)
	if config.Playback.DeviceID != nil {
		// the below device ID is stored in memory, whos pointers are freed when the user calls Device.Free.
		cConfig.playback.pDeviceID = (*C.ma_device_id)(unsafe.Pointer(memory.cbytes(config.Playback.DeviceID[:])))
	}
	cConfig.playback.format = C.ma_format(config.Playback.Format)
	cConfig.playback.channels = C.ma_uint32(config.Playback.Channels)
	for i, v := range config.Playback.ChannelMap {
		cConfig.playback.channelMap[i] = C.ma_channel(v)
	}
	cConfig.playback.channelMixMode = C.ma_channel_mix_mode(config.Playback.ChannelMixMode)
	cConfig.playback.shareMode = C.ma_share_mode(config.Playback.ShareMode)

	if config.Capture.DeviceID != nil {
		cConfig.capture.pDeviceID = (*C.ma_device_id)(unsafe.Pointer(memory.cbytes(config.Capture.DeviceID[:])))
	}
	cConfig.capture.format = C.ma_format(config.Capture.Format)
	cConfig.capture.channels = C.ma_uint32(config.Capture.Channels)
	for i, v := range config.Capture.ChannelMap {
		cConfig.capture.channelMap[i] = C.ma_channel(v)
	}
	cConfig.capture.channelMixMode = C.ma_channel_mix_mode(config.Capture.ChannelMixMode)
	cConfig.capture.shareMode = C.ma_share_mode(config.Capture.ShareMode)

	cConfig.wasapi.noAutoConvertSRC = boolToInt8(config.Wasapi.NoAutoConvertSRC)
	cConfig.wasapi.noDefaultQualitySRC = boolToInt8(config.Wasapi.NoDefaultQualitySRC)
	cConfig.wasapi.noAutoStreamRouting = boolToInt8(config.Wasapi.NoAutoStreamRouting)
	cConfig.wasapi.noHardwareOffloading = boolToInt8(config.Wasapi.NoHardwareOffloading)

	cConfig.alsa.noMMap = boolToInt(config.Alsa.NoMMap)
	cConfig.alsa.noAutoFormat = boolToInt(config.Alsa.NoAutoFormat)
	cConfig.alsa.noAutoChannels = boolToInt(config.Alsa.NoAutoChannels)
	cConfig.alsa.noAutoResample = boolToInt(config.Alsa.NoAutoResample)

	cConfig.pulse.pStreamNamePlayback = memory.cString(config.Pulse.StreamNamePlayback)
	cConfig.pulse.pStreamNameCapture = memory.cString(config.Pulse.StreamNameCapture)

	cConfig.coreaudio.allowNominalSampleRateChange = boolToInt(config.CoreAudio.AllowNominalSampleRateChange)

	cConfig.opensl.streamType = C.ma_opensl_stream_type(config.Opensl.StreamType)
	cConfig.opensl.recordingPreset = C.ma_opensl_recording_preset(config.Opensl.RecordingPreset)

	cConfig.aaudio.usage = C.ma_aaudio_usage(config.Aaudio.Usage)
	cConfig.aaudio.contentType = C.ma_aaudio_content_type(config.Aaudio.ContentType)
	cConfig.aaudio.inputPreset = C.ma_aaudio_input_preset(config.Aaudio.InputPreset)
	return cConfig, memory
}
