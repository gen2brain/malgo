package malgo

import (

)
/*
#include "miniaudio.h"
*/
import "C"



const (
	errorTag = "miniaudio: "
)
// The error type used by all Miniaudio functions. It's comparable to all the below constants
type Error int

const (
	success = C.MA_SUCCESS

	//  General errors.
	ErrGeneric            = Error(C.MA_ERROR)
	ErrInvalidArgs      = Error(C.MA_INVALID_ARGS)
	ErrInvalidOperation = Error(C.MA_INVALID_OPERATION)
	ErrOutOfMemory      = Error(C.MA_OUT_OF_MEMORY)
	ErrOutOfRange = Error(C.MA_OUT_OF_RANGE)
	ErrAccessDenied     = Error(C.MA_ACCESS_DENIED)
	ErrDoesNotExist = Error(C.MA_DOES_NOT_EXIST)
	ErrAlreadyExists = Error(C.MA_ALREADY_EXISTS)
	ErrTooManyOpenFiles = Error(C.MA_TOO_MANY_OPEN_FILES)
	ErrInvalidFile = Error(C.MA_INVALID_FILE)
	ErrTooBig         = Error(C.MA_TOO_BIG)
	ErrPathTooLong = Error(C.MA_PATH_TOO_LONG)
	ErrNameTooLong = Error(C.MA_NAME_TOO_LONG)
	ErrNotDirectory = Error(C.MA_NOT_DIRECTORY)
	ErrIsDirectory = Error(C.MA_IS_DIRECTORY)
	ErrDirectoryNotEmpty = Error(C.MA_DIRECTORY_NOT_EMPTY)
	ErrEndOfFile = Error(C.MA_END_OF_FILE)
	ErrNoSpace = Error(C.MA_NO_SPACE)
	ErrBusy = Error(C.MA_BUSY)
	ErrIO = Error(C.MA_IO_ERROR)
	ErrInterrupt = Error(C.MA_INTERRUPT)
	ErrUnavailable = Error(C.MA_UNAVAILABLE)
	ErrAlreadyInUse = Error(C.MA_ALREADY_IN_USE)
	ErrBadAddress = Error(C.MA_BAD_ADDRESS)
	ErrBadSeek = Error(C.MA_BAD_SEEK)
	ErrBadPipe = Error(C.MA_BAD_PIPE)
	ErrDeadlock = Error(C.MA_DEADLOCK)
	ErrTooManyLinks = Error(C.MA_TOO_MANY_LINKS)
	ErrNotImplemented = Error(C.MA_NOT_IMPLEMENTED)
	ErrNoMessage = Error(C.MA_NO_MESSAGE)
	ErrBadMessage = Error(C.MA_BAD_MESSAGE)
	ErrNoDataAvailable = Error(C.MA_NO_DATA_AVAILABLE)
	ErrInvalidData = Error(C.MA_INVALID_DATA)
	ErrTimeout = Error(C.MA_TIMEOUT)
	ErrNetwork = Error(C.MA_NO_NETWORK)
	ErrNotUnique = Error(C.MA_NOT_UNIQUE)
	ErrNotSocket = Error(C.MA_NOT_SOCKET)
	ErrNoAddress = Error(C.MA_NO_ADDRESS)
	ErrBadProtocol = Error(C.MA_BAD_PROTOCOL)
	ErrProtocolUnavailable = Error(C.MA_PROTOCOL_UNAVAILABLE)
	ErrProtocolNotSupported = Error(C.MA_PROTOCOL_NOT_SUPPORTED)
	ErrProtocolFamilyNotSupported = Error(C.MA_PROTOCOL_FAMILY_NOT_SUPPORTED)
	ErrAddressFamilyNotSupported = Error(C.MA_ADDRESS_FAMILY_NOT_SUPPORTED)
	ErrSocketNotSupported = Error(C.MA_SOCKET_NOT_SUPPORTED)
	ErrConnectionReset = Error(C.MA_CONNECTION_RESET)
	ErrAlreadyConnected = Error(C.MA_ALREADY_CONNECTED)
	ErrNotConnected = Error(C.MA_NOT_CONNECTED)
	ErrConnectionRefused = Error(C.MA_CONNECTION_REFUSED)
	ErrNoHost = Error(C.MA_NO_HOST)
	ErrInProgress = Error(C.MA_IN_PROGRESS)
	ErrCancelled = Error(C.MA_CANCELLED)
	ErrMemoryAlreadyMapped = Error(C.MA_MEMORY_ALREADY_MAPPED)
	ErrAtEnd = Error(C.MA_AT_END)






	// General miniaudio-specific errors.
	ErrFormatNotSupported     = Error(C.MA_FORMAT_NOT_SUPPORTED)
	ErrDeviceTypeNotSupported = Error(C.MA_DEVICE_TYPE_NOT_SUPPORTED)
	ErrShareModeNotSupported  = Error(C.MA_SHARE_MODE_NOT_SUPPORTED)
	ErrNoBackend              = Error(C.MA_NO_BACKEND)
	ErrNoDevice               = Error(C.MA_NO_DEVICE)
	ErrAPINotFound            = Error(C.MA_API_NOT_FOUND)
	ErrInvalidDeviceConfig    = Error(C.MA_INVALID_DEVICE_CONFIG)

	// State errors.

	ErrDeviceNotInitialized = Error(C.MA_DEVICE_NOT_INITIALIZED)
	ErrDeviceAlreadyInitialized = Error(C.MA_DEVICE_ALREADY_INITIALIZED)
	ErrDeviceNotStarted     = Error(C.MA_DEVICE_NOT_STARTED)
	ErrDeviceNotStopped = Error(C.MA_DEVICE_NOT_STOPPED)

	// Operation errors.

	ErrFailedToInitBackend            = Error(C.MA_FAILED_TO_INIT_BACKEND)
	ErrFailedToOpenBackendDevice      = Error(C.MA_FAILED_TO_OPEN_BACKEND_DEVICE)
	ErrFailedToStartBackendDevice     = Error(C.MA_FAILED_TO_START_BACKEND_DEVICE)
	ErrFailedToStopBackendDevice      = Error(C.MA_FAILED_TO_STOP_BACKEND_DEVICE)
)
// returns a brief description of what went wrong
func (self Error) Error() string {
	return errorTag + goString(C.ma_result_description(C.ma_result(self)))
}



func errorFromResult(r C.ma_result) error {
	switch r {
	case success:
		return nil
		default:
		return Error(r)
	}
}
