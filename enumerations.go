package malgo

// Backend type.
type Backend uint32

// Backend enumeration.
const (
	BackendWasapi = iota
	BackendDsound
	BackendWinmm
	BackendCoreaudio
	BackendSndio
	BackendAudio4
	BackendOss
	BackendPulseaudio
	BackendAlsa
	BackendJack
	BackendAaudio
	BackendOpensl
	BackendWebaudio
	BackendNull
)

// DeviceType type.
type DeviceType uint32

// DeviceType enumeration.
const (
	Playback DeviceType = iota + 1
	Capture
	Duplex
	Loopback
)

// ShareMode type.
type ShareMode uint32

// ShareMode enumeration.
const (
	Shared ShareMode = iota
	Exclusive
)

// PerformanceProfile type.
type PerformanceProfile uint32

// PerformanceProfile enumeration.
const (
	LowLatency PerformanceProfile = iota
	Conservative
)

// FormatType type.
type FormatType uint32

// Format enumeration.
const (
	FormatUnknown FormatType = iota
	FormatU8
	FormatS16
	FormatS24
	FormatS32
	FormatF32
)

// ThreadPriority type.
type ThreadPriority int32

// ThreadPriority enumeration.
const (
	ThreadPriorityIdle     ThreadPriority = -5
	ThreadPriorityLowest   ThreadPriority = -4
	ThreadPriorityLow      ThreadPriority = -3
	ThreadPriorityNormal   ThreadPriority = -2
	ThreadPriorityHigh     ThreadPriority = -1
	ThreadPriorityHighest  ThreadPriority = 0
	ThreadPriorityRealtime ThreadPriority = 1

	ThreadPriorityDefault ThreadPriority = 0
)

// ResampleAlgorithm type.
type ResampleAlgorithm uint32

// ResampleAlgorithm enumeration.
const (
	ResampleAlgorithmLinear ResampleAlgorithm = 0
	ResampleAlgorithmSpeex  ResampleAlgorithm = 1
)

// IOSSessionCategory type.
type IOSSessionCategory uint32

// IOSSessionCategory enumeration.
const (
	IOSSessionCategoryDefault       IOSSessionCategory = iota // AVAudioSessionCategoryPlayAndRecord with AVAudioSessionCategoryOptionDefaultToSpeaker.
	IOSSessionCategoryNone                                    // Leave the session category unchanged.
	IOSSessionCategoryAmbient                                 // AVAudioSessionCategoryAmbient
	IOSSessionCategorySoloAmbient                             // AVAudioSessionCategorySoloAmbient
	IOSSessionCategoryPlayback                                // AVAudioSessionCategoryPlayback
	IOSSessionCategoryRecord                                  // AVAudioSessionCategoryRecord
	IOSSessionCategoryPlayAndRecord                           // AVAudioSessionCategoryPlayAndRecord
	IOSSessionCategoryMultiRoute                              // AVAudioSessionCategoryMultiRoute
)

// IOSSessionCategoryOptions type.
type IOSSessionCategoryOptions uint32

// IOSSessionCategoryOptions enumeration.
const (
	IOSSessionCategoryOptionMixWithOthers                        = 0x01 // AVAudioSessionCategoryOptionMixWithOthers
	IOSSessionCategoryOptionDuckOthers                           = 0x02 // AVAudioSessionCategoryOptionDuckOthers
	IOSSessionCategoryOptionAllowBluetooth                       = 0x04 // AVAudioSessionCategoryOptionAllowBluetooth
	IOSSessionCategoryOptionDefaultToSpeaker                     = 0x08 // AVAudioSessionCategoryOptionDefaultToSpeaker
	IOSSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers = 0x11 // AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers
	IOSSessionCategoryOptionAllowBluetoothA2dp                   = 0x20 // AVAudioSessionCategoryOptionAllowBluetoothA2DP
	IOSSessionCategoryOptionAllowAirPlay                         = 0x40 // AVAudioSessionCategoryOptionAllowAirPlay
)

// DataFormatFlag type.
type DataFormatFlag uint32

const (
	DataFormatFlagNone          DataFormatFlag = 0
	DataFormatFlagExclusiveMode DataFormatFlag = 1 << 1
)

var flagNames = map[DataFormatFlag]string{
	DataFormatFlagExclusiveMode: "ExclusiveMode",
}

func (f *DataFormatFlag) Set(flag DataFormatFlag) {
	*f |= flag
}

func (f *DataFormatFlag) Clear(flag DataFormatFlag) {
	*f &^= flag
}

func (f DataFormatFlag) Has(flag DataFormatFlag) bool {
	return f&flag != 0
}

func (f DataFormatFlag) List() []string {
	var active []string
	for flag, name := range flagNames {
		if f.Has(flag) {
			active = append(active, name)
		}
	}
	return active
}
// AAudioUsage type.
type AAudioUsage uint32

// AAudioUsages enumeration.
const (
	AAudioUsageDefault AAudioUsage = iota
	AAudioUsageMedia
	AAudioUsageVoiceCommunication
	AAudioUsageVoiceCommunicationSignalling
	AAudioUsageAlarm
	AAudioUsageNotification
	AAudioUsageNotificationRingtone
	AAudioUsageNotificationEvent
	AAudioUsageAssistanceAccessibility
	AAudioUsageAssistanceNavigationGuidance
	AAudioUsageAssistanceSonification
	AAudioUsageGame
	AAudioUsageAssitant
	AAudioUsageEmergency
	AAudioUsageSafety
	AAudioUsageVehicleStatus
	AAudioUsageAnnouncement
)

// AAudioContentType type.
type AAudioContentType uint32

// AAudioContentTypes enumeration.
const (
	AAudioContentTypeDefault AAudioContentType = iota
	AAudioContentTypeSpeech
	AAudioContentTypeMusic
	AAudioContentTypeMovie
	AAudioContentTypeSonification
)

// AAudioInputPreset type.
type AAudioInputPreset uint32

// AAudioInputPresets enumeration.
const (
	AAudioInputPresetDefault AAudioInputPreset = iota
	AAudioInputPresetGeneric
	AAudioInputPresetCamcorder
	AAudioInputPresetVoiceRecognition
	AAudioInputPresetVoiceCommunication
	AAudioInputPresetUnprocessed
	AAudioInputPresetVoicePerformance
)

// AAudioAllowedCapturePolicy type.
type AAudioAllowedCapturePolicy uint32

// AAudioAllowedCapturePolicies enumeration.
const (
	AAudioAllowCapturePolicyDefault AAudioAllowedCapturePolicy = iota
	AAudioAllowCapturePolicyByAll
	AAudioAllowCapturePolicyBySystem
	AAudioAllowCapturePolicyByNone
)
