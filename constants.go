package malgo

/*
#include "miniaudio.h"
*/
import "C"
// Backend type.
const simdAlignment = C.MA_SIMD_ALIGNMENT

type Backend int
// Backend enumeration.
const (
	BackendWasapi = C.ma_backend_wasapi
	BackendDsound = C.ma_backend_dsound
	BackendWinmm = C.ma_backend_winmm
	BackendCoreaudio = C.ma_backend_coreaudio
	BackendSndio = C.ma_backend_sndio
	BackendAudio4 = C.ma_backend_audio4
	BackendOss = C.ma_backend_oss
	BackendPulseaudio = C.ma_backend_pulseaudio
	BackendAlsa = C.ma_backend_alsa
	BackendJack = C.ma_backend_jack
	BackendAaudio = C.ma_backend_aaudio
	BackendOpensl = C.ma_backend_opensl
	BackendWebaudio = C.ma_backend_webaudio
	BackendNull = C.ma_backend_null
)
// returns the name of this backend
func (b Backend) String() string {
	return goString(C.ma_get_backend_name(C.ma_backend(b)))
}



type DeviceType int
// DeviceType enumeration.
const (
	Playback DeviceType = C.ma_device_type_playback
	Capture = C.ma_device_type_capture
	Duplex = C.ma_device_type_duplex
	Loopback = C.ma_device_type_loopback
)



type ShareMode int
// ShareMode enumeration.
const (
	Shared ShareMode = C.ma_share_mode_shared
	Exclusive = C.ma_share_mode_exclusive
)

type PerformanceProfile int
// PerformanceProfile enumeration.
const (
	LowLatency PerformanceProfile = C.ma_performance_profile_low_latency
	Conservative = C.ma_performance_profile_conservative
)


type Format int
// Format enumeration.
const (
	FormatUnknown Format = C.ma_format_unknown
	FormatU8 = C.ma_format_u8
	FormatS16 = C.ma_format_s16
	FormatS24 = C.ma_format_s24
	FormatS32 = C.ma_format_s32
	FormatF32 = C.ma_format_f32
)

func (f Format) String() string {
	return goString(C.ma_get_format_name(C.ma_format(f)))
}

type ThreadPriority int
// ThreadPriority enumeration.
const (
	ThreadPriorityIdle     ThreadPriority = C.ma_thread_priority_idle
	ThreadPriorityLowest   ThreadPriority = C.ma_thread_priority_lowest
	ThreadPriorityLow      ThreadPriority = C.ma_thread_priority_low
	ThreadPriorityNormal   ThreadPriority = C.ma_thread_priority_normal
	ThreadPriorityHigh     ThreadPriority = C.ma_thread_priority_high
	ThreadPriorityHighest  ThreadPriority = C.ma_thread_priority_highest
	ThreadPriorityRealtime ThreadPriority = C.ma_thread_priority_realtime

	ThreadPriorityDefault ThreadPriority = C.ma_thread_priority_default
)
// ResampleAlgorithm type.
type ResampleAlgorithm int

// ResampleAlgorithm enumeration.
const (
	ResampleAlgorithmLinear ResampleAlgorithm = C.ma_resample_algorithm_linear
	ResampleAlgorithmSpeex  ResampleAlgorithm = C.ma_resample_algorithm_speex
)

// IOSSessionCategory type.
type IOSSessionCategory int

// IOSSessionCategory enumeration.
const (
	IOSSessionCategoryDefault       IOSSessionCategory = C.ma_ios_session_category_default // AVAudioSessionCategoryPlayAndRecord with AVAudioSessionCategoryOptionDefaultToSpeaker.
	IOSSessionCategoryNone                                 =  C.ma_ios_session_category_none  // Leave the session category unchanged.
	IOSSessionCategoryAmbient                               = C.ma_ios_session_category_ambient  // AVAudioSessionCategoryAmbient
	IOSSessionCategorySoloAmbient                           = C.ma_ios_session_category_solo_ambient  // AVAudioSessionCategorySoloAmbient
	IOSSessionCategoryPlayback                              = C.ma_ios_session_category_playback  // AVAudioSessionCategoryPlayback
	IOSSessionCategoryRecord                                = C.ma_ios_session_category_record  // AVAudioSessionCategoryRecord
	IOSSessionCategoryPlayAndRecord                         = C.ma_ios_session_category_play_and_record  // AVAudioSessionCategoryPlayAndRecord
	IOSSessionCategoryMultiRoute                           = C.ma_ios_session_category_multi_route   // AVAudioSessionCategoryMultiRoute
)

// IOSSessionCategoryOptions type.
type IOSSessionCategoryOptions uint32

// IOSSessionCategoryOptions enumeration.
const (
	IOSSessionCategoryOptionMixWithOthers                        = C.ma_ios_session_category_option_mix_with_others // AVAudioSessionCategoryOptionMixWithOthers
	IOSSessionCategoryOptionDuckOthers                           = C.ma_ios_session_category_option_duck_others // AVAudioSessionCategoryOptionDuckOthers
	IOSSessionCategoryOptionAllowBluetooth                       = C.ma_ios_session_category_option_allow_bluetooth // AVAudioSessionCategoryOptionAllowBluetooth
	IOSSessionCategoryOptionDefaultToSpeaker                     = C.ma_ios_session_category_option_default_to_speaker // AVAudioSessionCategoryOptionDefaultToSpeaker
	IOSSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers = C.ma_ios_session_category_option_interrupt_spoken_audio_and_mix_with_others // AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers
	IOSSessionCategoryOptionAllowBluetoothA2dp                   = C.ma_ios_session_category_option_allow_bluetooth_a2dp // AVAudioSessionCategoryOptionAllowBluetoothA2DP
	IOSSessionCategoryOptionAllowAirPlay                         = C.ma_ios_session_category_option_allow_air_play // AVAudioSessionCategoryOptionAllowAirPlay
)
