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

// Result type.
type Result int32

// Return codes.
const (
	Success = 0

	//  General errors.
	Error            = -1 // a generic error
	InvalidArgs      = -2
	InvalidPperation = -3
	OutOfMemory      = -4
	AccessDenied     = -5
	TooLarge         = -6
	Timeout          = -7

	// General miniaudio-specific errors.
	FormatNotSupported     = -100
	DeviceTypeNotSupported = -101
	ShareModeNotSupported  = -102
	NoBackend              = -103
	NoDevice               = -104
	APINotFound            = -105
	InvalidDeviceConfig    = -106

	// State errors.
	DeviceBusy           = -200
	DeviceNotInitialized = -201
	DeviceNotStarted     = -202
	DeviceUnavailable    = -203

	// Operation errors.
	FailedToMapDeviceBuffer        = -300
	FailedToUnmapDeviceBuffer      = -301
	FailedToInitBackend            = -302
	FailedToReadDataFromDevice     = -304
	FailedToReadDataFromClient     = -303
	FailedToSendDataToClient       = -305
	FailedToSendDataToDevice       = -306
	FailedToOpenBackendDevice      = -307
	FailedToStartBackendDevice     = -308
	FailedToStopBackendDevice      = -309
	FailedToConfigureBackendDevice = -310
	FailedToCreateMutex            = -311
	FailedToCreateEvent            = -312
	FailedToCreateThread           = -313
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
