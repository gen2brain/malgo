package malgo

import "fmt"

// Errors.
var (
	errTag = "miniaudio"

	ErrError                          = fmt.Errorf("%s: generic error", errTag)
	ErrInvalidArgs                    = fmt.Errorf("%s: invalid args", errTag)
	ErrInvalidPperation               = fmt.Errorf("%s: invalid pperation", errTag)
	ErrOutOfMemory                    = fmt.Errorf("%s: out o fmemory", errTag)
	ErrAccessDenied                   = fmt.Errorf("%s: access denied", errTag)
	ErrTooLarge                       = fmt.Errorf("%s: too large", errTag)
	ErrTimeout                        = fmt.Errorf("%s: timeout", errTag)
	ErrFormatNotSupported             = fmt.Errorf("%s: format not supported", errTag)
	ErrDeviceTypeNotSupported         = fmt.Errorf("%s: device type not supported", errTag)
	ErrShareModeNotSupported          = fmt.Errorf("%s: share mode not supported", errTag)
	ErrNoBackend                      = fmt.Errorf("%s: no backend", errTag)
	ErrNoDevice                       = fmt.Errorf("%s: no device", errTag)
	ErrAPINotFound                    = fmt.Errorf("%s: api not found", errTag)
	ErrInvalidDeviceConfig            = fmt.Errorf("%s: invalid device config", errTag)
	ErrDeviceBusy                     = fmt.Errorf("%s: device busy", errTag)
	ErrDeviceNotInitialized           = fmt.Errorf("%s: device not initialized", errTag)
	ErrDeviceNotStarted               = fmt.Errorf("%s: device not started", errTag)
	ErrDeviceUnavailable              = fmt.Errorf("%s: device unavailable", errTag)
	ErrFailedToMapDeviceBuffer        = fmt.Errorf("%s: failed to map device buffer", errTag)
	ErrFailedToUnmapDeviceBuffer      = fmt.Errorf("%s: failed to unmap device buffer", errTag)
	ErrFailedToInitBackend            = fmt.Errorf("%s: failed to init backend", errTag)
	ErrFailedToReadDataFromDevice     = fmt.Errorf("%s: failed to read data from device", errTag)
	ErrFailedToReadDataFromClient     = fmt.Errorf("%s: failed to read data from client", errTag)
	ErrFailedToSendDataToClient       = fmt.Errorf("%s: failed to send data to client", errTag)
	ErrFailedToSendDataToDevice       = fmt.Errorf("%s: failed to send data to device", errTag)
	ErrFailedToOpenBackendDevice      = fmt.Errorf("%s: failed to open backend device", errTag)
	ErrFailedToStartBackendDevice     = fmt.Errorf("%s: failed to start backend device", errTag)
	ErrFailedToStopBackendDevice      = fmt.Errorf("%s: failed to stop backend device", errTag)
	ErrFailedToConfigureBackendDevice = fmt.Errorf("%s: failed to configure backend device", errTag)
	ErrFailedToCreateMutex            = fmt.Errorf("%s: failed to create mutex", errTag)
	ErrFailedToCreateEvent            = fmt.Errorf("%s: failed to create event", errTag)
	ErrFailedToCreateThread           = fmt.Errorf("%s: failed to create thread", errTag)
)

// errorFromResult returns error for result code.
func errorFromResult(r Result) error {
	switch r {
	case Success:
		return nil
	case Error:
		return ErrError
	case InvalidArgs:
		return ErrInvalidArgs
	case InvalidPperation:
		return ErrInvalidPperation
	case OutOfMemory:
		return ErrOutOfMemory
	case AccessDenied:
		return ErrAccessDenied
	case TooLarge:
		return ErrTooLarge
	case Timeout:
		return ErrTimeout
	case FormatNotSupported:
		return ErrFormatNotSupported
	case DeviceTypeNotSupported:
		return ErrDeviceTypeNotSupported
	case ShareModeNotSupported:
		return ErrShareModeNotSupported
	case NoBackend:
		return ErrNoBackend
	case NoDevice:
		return ErrNoDevice
	case APINotFound:
		return ErrAPINotFound
	case InvalidDeviceConfig:
		return ErrInvalidDeviceConfig
	case DeviceBusy:
		return ErrDeviceBusy
	case DeviceNotInitialized:
		return ErrDeviceNotInitialized
	case DeviceNotStarted:
		return ErrDeviceNotStarted
	case DeviceUnavailable:
		return ErrDeviceUnavailable
	case FailedToMapDeviceBuffer:
		return ErrFailedToMapDeviceBuffer
	case FailedToUnmapDeviceBuffer:
		return ErrFailedToUnmapDeviceBuffer
	case FailedToInitBackend:
		return ErrFailedToInitBackend
	case FailedToReadDataFromDevice:
		return ErrFailedToReadDataFromDevice
	case FailedToReadDataFromClient:
		return ErrFailedToReadDataFromClient
	case FailedToSendDataToClient:
		return ErrFailedToSendDataToClient
	case FailedToSendDataToDevice:
		return ErrFailedToSendDataToDevice
	case FailedToOpenBackendDevice:
		return ErrFailedToOpenBackendDevice
	case FailedToStartBackendDevice:
		return ErrFailedToStartBackendDevice
	case FailedToStopBackendDevice:
		return ErrFailedToStopBackendDevice
	case FailedToConfigureBackendDevice:
		return ErrFailedToConfigureBackendDevice
	case FailedToCreateMutex:
		return ErrFailedToCreateMutex
	case FailedToCreateEvent:
		return ErrFailedToCreateEvent
	case FailedToCreateThread:
		return ErrFailedToCreateThread
	default:
		return ErrError
	}
}
