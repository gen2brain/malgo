package mini_al

// Backend type.
type Backend uint32

// Backend enumeration.
const (
	BackendNull Backend = iota
	BackendWasapi
	BackendDsound
	BackendWinmm
	BackendAlsa
	BackendPulseAudio
	BackendJack
	BackendCoreAudio
	BackendSndio
	BackendAudio4
	BackendOss
	BackendOpensl
	BackendOpenal
	BackendSdl
)

// DeviceType type.
type DeviceType uint32

// DeviceType enumeration.
const (
	Playback DeviceType = iota
	Capture
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
	Success                        = 0
	Error                          = -1
	InvalidArgs                    = -2
	InvalidOperation               = -3
	OutOfMemory                    = -4
	FormatNotSupported             = -5
	NoBackend                      = -6
	NoDevice                       = -7
	APINotFound                    = -8
	DeviceBusy                     = -9
	DeviceNotInitialized           = -10
	DeviceNotStarted               = -11
	DeviceNotStopped               = -12
	DeviceAlreadyStarted           = -13
	DeviceAlreadyStarting          = -14
	DeviceAlreadyStopped           = -15
	DeviceAlreadyStopping          = -16
	FailedToMapDeviceBuffer        = -17
	FailedToUnmapDeviceBuffer      = -18
	FailedToInitBackend            = -19
	FailedToReadDataFromClient     = -20
	FailedToReadDataFromDevice     = -21
	FailedToSendDataToClient       = -22
	FailedToSendDataToDevice       = -23
	FailedToOpenBackendDevice      = -24
	FailedToStartBackendDevice     = -25
	FailedToStopBackendDevice      = -26
	FailedToConfigureBackendDevice = -27
	FailedToCreateMutex            = -28
	FailedToCreateEvent            = -29
	FailedToCreateThread           = -30
	InvalidDeviceConfig            = -31
	AccessDenied                   = -32
	TooLarge                       = -33
	DeviceUnavailable              = -34
	Timeout                        = -35
)
