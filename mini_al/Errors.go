package mini_al

import "fmt"

// Errors.
var (
	errTag = "mini_al"

	ErrError                          = fmt.Errorf("%s: generic error", errTag)
	ErrInvalidArgs                    = fmt.Errorf("%s: invalid args", errTag)
	ErrInvalidOperation               = fmt.Errorf("%s: invalid operation", errTag)
	ErrOutOfMemory                    = fmt.Errorf("%s: out of memory", errTag)
	ErrFormatNotSupported             = fmt.Errorf("%s: format not supported", errTag)
	ErrNoBackend                      = fmt.Errorf("%s: no backend", errTag)
	ErrNoDevice                       = fmt.Errorf("%s: no device", errTag)
	ErrAPINotFound                    = fmt.Errorf("%s: api not found", errTag)
	ErrDeviceBusy                     = fmt.Errorf("%s: device busy", errTag)
	ErrDeviceNotInitialized           = fmt.Errorf("%s: device busy", errTag)
	ErrDeviceNotStarted               = fmt.Errorf("%s: device not started", errTag)
	ErrDeviceNotStopped               = fmt.Errorf("%s: device not stopped", errTag)
	ErrDeviceAlreadyStarted           = fmt.Errorf("%s: device already started", errTag)
	ErrDeviceAlreadyStarting          = fmt.Errorf("%s: device already starting", errTag)
	ErrDeviceAlreadyStopped           = fmt.Errorf("%s: device already stopped", errTag)
	ErrDeviceAlreadyStopping          = fmt.Errorf("%s: device already stopping", errTag)
	ErrFailedToMapDeviceBuffer        = fmt.Errorf("%s: failed to map device buffer", errTag)
	ErrFailedToUnmapDeviceBuffer      = fmt.Errorf("%s: failed to unmap device buffer", errTag)
	ErrFailedToInitBackend            = fmt.Errorf("%s: failed to init backend", errTag)
	ErrFailedToReadDataFromClient     = fmt.Errorf("%s: failed to read data from client", errTag)
	ErrFailedToReadDataFromDevice     = fmt.Errorf("%s: failed to read data from device", errTag)
	ErrFailedToSendDataToClient       = fmt.Errorf("%s: failed to send data to client", errTag)
	ErrFailedToSendDataToDevice       = fmt.Errorf("%s: failed to send data to device", errTag)
	ErrFailedToOpenBackendDevice      = fmt.Errorf("%s: failed to open backend device", errTag)
	ErrFailedToStartBackendDevice     = fmt.Errorf("%s: failed to start backend device", errTag)
	ErrFailedToStopBackendDevice      = fmt.Errorf("%s: failed to stop backend device", errTag)
	ErrFailedToConfigureBackendDevice = fmt.Errorf("%s: failed to configure backend device", errTag)
	ErrFailedToCreateMutex            = fmt.Errorf("%s: failed to create mutex", errTag)
	ErrFailedToCreateEvent            = fmt.Errorf("%s: failed to create event", errTag)
	ErrFailedToCreateThread           = fmt.Errorf("%s: failed to create thread", errTag)
	ErrInvalidDeviceConfig            = fmt.Errorf("%s: invalid device config", errTag)
	ErrAccessDenied                   = fmt.Errorf("%s: access denied", errTag)
	ErrTooLarge                       = fmt.Errorf("%s: too large", errTag)
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
	case InvalidOperation:
		return ErrInvalidOperation
	case OutOfMemory:
		return ErrOutOfMemory
	case FormatNotSupported:
		return ErrFormatNotSupported
	case NoBackend:
		return ErrNoBackend
	case NoDevice:
		return ErrNoDevice
	case APINotFound:
		return ErrAPINotFound
	case DeviceBusy:
		return ErrDeviceBusy
	case DeviceNotInitialized:
		return ErrDeviceNotInitialized
	case DeviceNotStarted:
		return ErrDeviceNotStarted
	case DeviceNotStopped:
		return ErrDeviceNotStopped
	case DeviceAlreadyStarted:
		return ErrDeviceAlreadyStarted
	case DeviceAlreadyStarting:
		return ErrDeviceAlreadyStarting
	case DeviceAlreadyStopped:
		return ErrDeviceAlreadyStopped
	case DeviceAlreadyStopping:
		return ErrDeviceAlreadyStopping
	case FailedToMapDeviceBuffer:
		return ErrFailedToMapDeviceBuffer
	case FailedToUnmapDeviceBuffer:
		return ErrFailedToUnmapDeviceBuffer
	case FailedToInitBackend:
		return ErrFailedToInitBackend
	case FailedToReadDataFromClient:
		return ErrFailedToReadDataFromClient
	case FailedToReadDataFromDevice:
		return ErrFailedToReadDataFromDevice
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
	case InvalidDeviceConfig:
		return ErrInvalidDeviceConfig
	case AccessDenied:
		return ErrAccessDenied
	case TooLarge:
		return ErrTooLarge
	default:
		return ErrError
	}
}
