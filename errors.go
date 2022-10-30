package malgo

/*
#include "malgo.h"
*/
import "C"

// Result type.
type Result int32

func (self Result) Error() string {
	return errTag + C.GoString(C.ma_result_description(C.ma_result(self)))
}

// Errors.
const (
	errTag = "miniaudio: "
)

var (
	//  General errors.
	ErrGeneric                    = Result(C.MA_ERROR)
	ErrInvalidArgs                = Result(C.MA_INVALID_ARGS)
	ErrInvalidOperation           = Result(C.MA_INVALID_OPERATION)
	ErrOutOfMemory                = Result(C.MA_OUT_OF_MEMORY)
	ErrOutOfRange                 = Result(C.MA_OUT_OF_RANGE)
	ErrAccessDenied               = Result(C.MA_ACCESS_DENIED)
	ErrDoesNotExist               = Result(C.MA_DOES_NOT_EXIST)
	ErrAlreadyExists              = Result(C.MA_ALREADY_EXISTS)
	ErrTooManyOpenFiles           = Result(C.MA_TOO_MANY_OPEN_FILES)
	ErrInvalidFile                = Result(C.MA_INVALID_FILE)
	ErrTooBig                     = Result(C.MA_TOO_BIG)
	ErrPathTooLong                = Result(C.MA_PATH_TOO_LONG)
	ErrNameTooLong                = Result(C.MA_NAME_TOO_LONG)
	ErrNotDirectory               = Result(C.MA_NOT_DIRECTORY)
	ErrIsDirectory                = Result(C.MA_IS_DIRECTORY)
	ErrDirectoryNotEmpty          = Result(C.MA_DIRECTORY_NOT_EMPTY)
	ErrAtEnd                      = Result(C.MA_AT_END)
	ErrNoSpace                    = Result(C.MA_NO_SPACE)
	ErrBusy                       = Result(C.MA_BUSY)
	ErrIO                         = Result(C.MA_IO_ERROR)
	ErrInterrupt                  = Result(C.MA_INTERRUPT)
	ErrUnavailable                = Result(C.MA_UNAVAILABLE)
	ErrAlreadyInUse               = Result(C.MA_ALREADY_IN_USE)
	ErrBadAddress                 = Result(C.MA_BAD_ADDRESS)
	ErrBadSeek                    = Result(C.MA_BAD_SEEK)
	ErrBadPipe                    = Result(C.MA_BAD_PIPE)
	ErrDeadlock                   = Result(C.MA_DEADLOCK)
	ErrTooManyLinks               = Result(C.MA_TOO_MANY_LINKS)
	ErrNotImplemented             = Result(C.MA_NOT_IMPLEMENTED)
	ErrNoMessage                  = Result(C.MA_NO_MESSAGE)
	ErrBadMessage                 = Result(C.MA_BAD_MESSAGE)
	ErrNoDataAvailable            = Result(C.MA_NO_DATA_AVAILABLE)
	ErrInvalidData                = Result(C.MA_INVALID_DATA)
	ErrTimeout                    = Result(C.MA_TIMEOUT)
	ErrNetwork                    = Result(C.MA_NO_NETWORK)
	ErrNotUnique                  = Result(C.MA_NOT_UNIQUE)
	ErrNotSocket                  = Result(C.MA_NOT_SOCKET)
	ErrNoAddress                  = Result(C.MA_NO_ADDRESS)
	ErrBadProtocol                = Result(C.MA_BAD_PROTOCOL)
	ErrProtocolUnavailable        = Result(C.MA_PROTOCOL_UNAVAILABLE)
	ErrProtocolNotSupported       = Result(C.MA_PROTOCOL_NOT_SUPPORTED)
	ErrProtocolFamilyNotSupported = Result(C.MA_PROTOCOL_FAMILY_NOT_SUPPORTED)
	ErrAddressFamilyNotSupported  = Result(C.MA_ADDRESS_FAMILY_NOT_SUPPORTED)
	ErrSocketNotSupported         = Result(C.MA_SOCKET_NOT_SUPPORTED)
	ErrConnectionReset            = Result(C.MA_CONNECTION_RESET)
	ErrAlreadyConnected           = Result(C.MA_ALREADY_CONNECTED)
	ErrNotConnected               = Result(C.MA_NOT_CONNECTED)
	ErrConnectionRefused          = Result(C.MA_CONNECTION_REFUSED)
	ErrNoHost                     = Result(C.MA_NO_HOST)
	ErrInProgress                 = Result(C.MA_IN_PROGRESS)
	ErrCancelled                  = Result(C.MA_CANCELLED)
	ErrMemoryAlreadyMapped        = Result(C.MA_MEMORY_ALREADY_MAPPED)

	// General miniaudio-specific errors.
	ErrFormatNotSupported     = Result(C.MA_FORMAT_NOT_SUPPORTED)
	ErrDeviceTypeNotSupported = Result(C.MA_DEVICE_TYPE_NOT_SUPPORTED)
	ErrShareModeNotSupported  = Result(C.MA_SHARE_MODE_NOT_SUPPORTED)
	ErrNoBackend              = Result(C.MA_NO_BACKEND)
	ErrNoDevice               = Result(C.MA_NO_DEVICE)
	ErrAPINotFound            = Result(C.MA_API_NOT_FOUND)
	ErrInvalidDeviceConfig    = Result(C.MA_INVALID_DEVICE_CONFIG)
	ErrLoop                   = Result(C.MA_LOOP)

	// State errors.

	ErrDeviceNotInitialized     = Result(C.MA_DEVICE_NOT_INITIALIZED)
	ErrDeviceAlreadyInitialized = Result(C.MA_DEVICE_ALREADY_INITIALIZED)
	ErrDeviceNotStarted         = Result(C.MA_DEVICE_NOT_STARTED)
	ErrDeviceNotStopped         = Result(C.MA_DEVICE_NOT_STOPPED)

	// Operation errors.

	ErrFailedToInitBackend        = Result(C.MA_FAILED_TO_INIT_BACKEND)
	ErrFailedToOpenBackendDevice  = Result(C.MA_FAILED_TO_OPEN_BACKEND_DEVICE)
	ErrFailedToStartBackendDevice = Result(C.MA_FAILED_TO_START_BACKEND_DEVICE)
	ErrFailedToStopBackendDevice  = Result(C.MA_FAILED_TO_STOP_BACKEND_DEVICE)
)

// errorFromResult returns error for result code.
func errorFromResult(r C.ma_result) error {
	switch r {
	case C.MA_SUCCESS:
		return nil
	default:
		return Result(r)
	}
}
