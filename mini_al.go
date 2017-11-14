// Package mal - Mini audio library (mini_al cgo bindings).
package mal

/*
#cgo CFLAGS: -std=gnu99 -Iexternal/include
#cgo linux LDFLAGS: -ldl
#cgo android LDFLAGS: -lOpenSLES

#include "mini_al.h"

extern void goRecvCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern mal_uint32 goSendCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern void goStopCallback(mal_device* pDevice);
extern void goLogCallback(mal_context* pContext, mal_device* pDevice, char* message);

void goSetRecvCallback(mal_device* pDevice);
void goSetSendCallback(mal_device* pDevice);
void goSetStopCallback(mal_device* pDevice);

mal_device* goGetDevice();
mal_context* goGetContext();

mal_device_config goConfigInit(mal_format format, mal_uint32 channels, mal_uint32 sampleRate);
mal_device_config goConfigInitCapture(mal_format format, mal_uint32 channels, mal_uint32 sampleRate);
mal_device_config goConfigInitPlayback(mal_format format, mal_uint32 channels, mal_uint32 sampleRate);
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Return codes
const (
	Success                              = 0
	Error                                = -1
	InvalidArgs                          = -2
	OutOfMemory                          = -3
	FormatNotSupported                   = -4
	NoBackend                            = -5
	NoDevice                             = -6
	APINotFound                          = -7
	DeviceBusy                           = -8
	DeviceNotInitialized                 = -9
	DeviceAlreadyStarted                 = -10
	DeviceAlreadyStarting                = -11
	DeviceAlreadyStopped                 = -12
	DeviceAlreadyStopping                = -13
	FailedToMapDeviceBuffer              = -14
	FailedToInitBackend                  = -15
	FailedToReadDataFromClient           = -16
	FailedToStartBackendDevice           = -17
	FailedToStopBackendDevice            = -18
	FailedToCreateMutex                  = -19
	FailedToCreateEvent                  = -20
	FailedToCreateUint                   = -21
	InvalidDeviceConfig                  = -22
	DsoundFailedToCreateDevice           = -1024
	DsoundFailedToSetCoopLevel           = -1025
	DsoundFailedToCreateBuffer           = -1026
	DsoundFailedToQueryInterface         = -1027
	DsoundFailedToSetNotifications       = -1028
	AlsaFailedToOpenDevice               = -2048
	AlsaFailedToSetHwParams              = -2049
	AlsaFailedToSetSwParams              = -2050
	AlsaFailedToPrepareDevice            = -2051
	AlsaFailedToRecoverDevice            = -2052
	WasapiFailedToCreateDeviceEnumerator = -3072
	WasapiFailedToCreateDevice           = -3073
	WasapiFailedToActivateDevice         = -3074
	WasapiFailedToInitializeDevice       = -3075
	WasapiFailedToFindBestFormat         = -3076
	WasapiFailedToGetInternalBuffer      = -3077
	WasapiFailedToReleaseInternalBuffer  = -3078
	WinmmFailedToGetDeviceCaps           = -4096
	WinmmFailedToGetSupportedFormats     = -4097
)

// Errors
var (
	errTag string = "mini_al"

	ErrError                                = fmt.Errorf("%s: generic error", errTag)
	ErrInvalidArgs                          = fmt.Errorf("%s: invalid args", errTag)
	ErrOutOfMemory                          = fmt.Errorf("%s: out of memory", errTag)
	ErrFormatNotSupported                   = fmt.Errorf("%s: format not supported", errTag)
	ErrNoBackend                            = fmt.Errorf("%s: no backend", errTag)
	ErrNoDevice                             = fmt.Errorf("%s: no device", errTag)
	ErrAPINotFound                          = fmt.Errorf("%s: api not found", errTag)
	ErrDeviceBusy                           = fmt.Errorf("%s: device busy", errTag)
	ErrDeviceNotInitialized                 = fmt.Errorf("%s: device busy", errTag)
	ErrDeviceAlreadyStarted                 = fmt.Errorf("%s: device already started", errTag)
	ErrDeviceAlreadyStarting                = fmt.Errorf("%s: device already starting", errTag)
	ErrDeviceAlreadyStopped                 = fmt.Errorf("%s: device already stopped", errTag)
	ErrDeviceAlreadyStopping                = fmt.Errorf("%s: device already stopping", errTag)
	ErrFailedToMapDeviceBuffer              = fmt.Errorf("%s: failed to map device buffer", errTag)
	ErrFailedToInitBackend                  = fmt.Errorf("%s: failed to init backend", errTag)
	ErrFailedToReadDataFromClient           = fmt.Errorf("%s: failed to read data from client", errTag)
	ErrFailedToStartBackendDevice           = fmt.Errorf("%s: failed to start backend device", errTag)
	ErrFailedToStopBackendDevice            = fmt.Errorf("%s: failed to stop backend device", errTag)
	ErrFailedToCreateMutex                  = fmt.Errorf("%s: failed to create mutex", errTag)
	ErrFailedToCreateEvent                  = fmt.Errorf("%s: failed to create event", errTag)
	ErrFailedToCreateUint                   = fmt.Errorf("%s: failed to create uint", errTag)
	ErrInvalidDeviceConfig                  = fmt.Errorf("%s: invalid device config", errTag)
	ErrDsoundFailedToCreateDevice           = fmt.Errorf("%s: dsound failed to create device", errTag)
	ErrDsoundFailedToSetCoopLevel           = fmt.Errorf("%s: dsound failed to set coop level", errTag)
	ErrDsoundFailedToCreateBuffer           = fmt.Errorf("%s: dsound failed to create buffer", errTag)
	ErrDsoundFailedToQueryInterface         = fmt.Errorf("%s: dsound failed to query interface", errTag)
	ErrDsoundFailedToSetNotifications       = fmt.Errorf("%s: dsound failed to set notifications", errTag)
	ErrAlsaFailedToOpenDevice               = fmt.Errorf("%s: alsa failed to open device", errTag)
	ErrAlsaFailedToSetHwParams              = fmt.Errorf("%s: alsa failed to set hw params", errTag)
	ErrAlsaFailedToSetSwParams              = fmt.Errorf("%s: alsa failed to set sw params", errTag)
	ErrAlsaFailedToPrepareDevice            = fmt.Errorf("%s: alsa failed to prepare device", errTag)
	ErrAlsaFailedToRecoverDevice            = fmt.Errorf("%s: alsa failed to recover device", errTag)
	ErrWasapiFailedToCreateDeviceEnumerator = fmt.Errorf("%s: wasapi failed to create device enumerator", errTag)
	ErrWasapiFailedToCreateDevice           = fmt.Errorf("%s: wasapi failed to create device", errTag)
	ErrWasapiFailedToActivateDevice         = fmt.Errorf("%s: wasapi failed to activate device", errTag)
	ErrWasapiFailedToInitializeDevice       = fmt.Errorf("%s: wasapi failed to initialize device", errTag)
	ErrWasapiFailedToFindBestFormat         = fmt.Errorf("%s: wasapi failed to find best format", errTag)
	ErrWasapiFailedToGetInternalBuffer      = fmt.Errorf("%s: wasapi failed to get internal buffer", errTag)
	ErrWasapiFailedToReleaseInternalBuffer  = fmt.Errorf("%s: wasapi failed to release internal buffer", errTag)
	ErrWinmmFailedToGetDeviceCaps           = fmt.Errorf("%s: winmm failed to get device caps", errTag)
	ErrWinmmFailedToGetSupportedFormats     = fmt.Errorf("%s: winmm failed to get supported formats", errTag)
)

// errorFromResult returns error for result code
func errorFromResult(r Result) error {
	switch r {
	case Success:
		return nil
	case Error:
		return ErrError
	case InvalidArgs:
		return ErrInvalidArgs
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
	case FailedToInitBackend:
		return ErrFailedToInitBackend
	case FailedToReadDataFromClient:
		return ErrFailedToReadDataFromClient
	case FailedToStartBackendDevice:
		return ErrFailedToStartBackendDevice
	case FailedToStopBackendDevice:
		return ErrFailedToStopBackendDevice
	case FailedToCreateMutex:
		return ErrFailedToCreateMutex
	case FailedToCreateEvent:
		return ErrFailedToCreateEvent
	case FailedToCreateUint:
		return ErrFailedToCreateUint
	case InvalidDeviceConfig:
		return ErrInvalidDeviceConfig
	case DsoundFailedToCreateDevice:
		return ErrDsoundFailedToCreateDevice
	case DsoundFailedToSetCoopLevel:
		return ErrDsoundFailedToSetCoopLevel
	case DsoundFailedToCreateBuffer:
		return ErrDsoundFailedToCreateBuffer
	case DsoundFailedToQueryInterface:
		return ErrDsoundFailedToQueryInterface
	case DsoundFailedToSetNotifications:
		return ErrDsoundFailedToSetNotifications
	case AlsaFailedToOpenDevice:
		return ErrAlsaFailedToOpenDevice
	case AlsaFailedToSetHwParams:
		return ErrAlsaFailedToSetHwParams
	case AlsaFailedToSetSwParams:
		return ErrAlsaFailedToSetSwParams
	case AlsaFailedToPrepareDevice:
		return ErrAlsaFailedToPrepareDevice
	case AlsaFailedToRecoverDevice:
		return ErrAlsaFailedToRecoverDevice
	case WasapiFailedToCreateDeviceEnumerator:
		return ErrWasapiFailedToCreateDeviceEnumerator
	case WasapiFailedToCreateDevice:
		return ErrWasapiFailedToCreateDevice
	case WasapiFailedToActivateDevice:
		return ErrWasapiFailedToActivateDevice
	case WasapiFailedToInitializeDevice:
		return ErrWasapiFailedToInitializeDevice
	case WasapiFailedToFindBestFormat:
		return ErrWasapiFailedToFindBestFormat
	case WasapiFailedToGetInternalBuffer:
		return ErrWasapiFailedToGetInternalBuffer
	case WasapiFailedToReleaseInternalBuffer:
		return ErrWasapiFailedToReleaseInternalBuffer
	case WinmmFailedToGetDeviceCaps:
		return ErrWinmmFailedToGetDeviceCaps
	case WinmmFailedToGetSupportedFormats:
		return ErrWinmmFailedToGetSupportedFormats
	default:
		return ErrError
	}
}

// Backend type
type Backend uint32

// Backend enumeration
const (
	BackendNull Backend = iota
	BackendWasapi
	BackendDsound
	BackendWinmm
	BackendAlsa
	BackendOss
	BackendOpensl
	BackendOpenal
)

// DeviceType type
type DeviceType uint32

// DeviceType enumeration
const (
	Playback DeviceType = iota
	Capture
)

// FormatType type
type FormatType uint32

// Format enumeration
const (
	FormatUnknown FormatType = iota
	FormatU8
	FormatS16
	FormatS24
	FormatS32
	FormatF32
)

// Result type
type Result int32

// Device type
type Device struct {
	context *C.mal_context
	device  *C.mal_device
}

// NewDevice returns new Device
func NewDevice() *Device {
	d := &Device{}
	d.context = C.goGetContext()
	d.device = C.goGetDevice()
	return d
}

// DeviceID type
type DeviceID [unsafe.Sizeof(C.mal_device_id{})]byte

// cptr return C pointer
func (d *DeviceID) cptr() *C.mal_device_id {
	return (*C.mal_device_id)(unsafe.Pointer(d))
}

// DeviceInfo type
type DeviceInfo struct {
	ID   DeviceID
	Name [256]byte
}

// String returns string
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("{ID: %s, Name: %s}", string(d.ID[:]), string(d.Name[:]))
}

// NewDeviceInfoFromPointer returns new DeviceInfo from pointer
func NewDeviceInfoFromPointer(ptr unsafe.Pointer) DeviceInfo {
	return *(*DeviceInfo)(ptr)
}

// AlsaDeviceConfig type
type AlsaDeviceConfig struct {
	NoMMap uint32
}

// DeviceConfig type
type DeviceConfig struct {
	Format             FormatType
	Channels           uint32
	SampleRate         uint32
	ChannelMap         [18]byte
	_                  [2]byte
	BufferSizeInFrames uint32
	Periods            uint32
	OnRecvCallback     *[0]byte
	OnSendCallback     *[0]byte
	OnStopCallback     *[0]byte
	OnLogCallback      *[0]byte
	Alsa               AlsaDeviceConfig
}

// cptr return C pointer
func (d *DeviceConfig) cptr() *C.mal_device_config {
	return (*C.mal_device_config)(unsafe.Pointer(d))
}

// NewDeviceConfigFromPointer returns new DeviceConfig from pointer
func NewDeviceConfigFromPointer(ptr unsafe.Pointer) DeviceConfig {
	return *(*DeviceConfig)(ptr)
}

// AlsaContextConfig type
type AlsaContextConfig struct {
	UseVerboseDeviceEnumeration uint32
	ExcludeNullDevice           uint32
}

// ContextConfig type
type ContextConfig struct {
	OnLogCallback *[0]byte
	Alsa          AlsaContextConfig
}

// RecvProc type
type RecvProc func(framecount uint32, psamples []byte)

// SendProc type
type SendProc func(framecount uint32, psamples []byte) uint32

// StopProc type
type StopProc func()

// Handlers
var (
	recvHandler RecvProc
	sendHandler SendProc
	stopHandler StopProc
)

//export goRecvCallback
func goRecvCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) {
	if recvHandler != nil {
		sampleCount := uint32(frameCount) * uint32(pDevice.channels)
		sizeInBytes := uint32(C.mal_get_sample_size_in_bytes(pDevice.format))
		psamples := (*[1 << 20]byte)(pSamples)[0 : sampleCount*sizeInBytes]
		recvHandler(uint32(frameCount), psamples)
	}
}

//export goSendCallback
func goSendCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) (r C.mal_uint32) {
	if sendHandler != nil {
		sampleCount := uint32(frameCount) * uint32(pDevice.channels)
		sizeInBytes := uint32(C.mal_get_sample_size_in_bytes(pDevice.format))
		psamples := (*[1 << 20]byte)(pSamples)[0 : sampleCount*sizeInBytes]
		r = C.mal_uint32(sendHandler(uint32(frameCount), psamples))
	}
	return r
}

//export goStopCallback
func goStopCallback(pDevice *C.mal_device) {
	if stopHandler != nil {
		stopHandler()
	}
}

//export goLogCallback
func goLogCallback(pContext *C.mal_context, pDevice *C.mal_device, message *C.char) {
	fmt.Printf("%s:%s\n", errTag, C.GoString(message))
}

// ContextInit initializes a context.
//
// The context is used for selecting and initializing the relevant backends.
//
// <backends> is used to allow the application to prioritize backends depending on it's specific
// requirements. This can be nil in which case it uses the default priority, which is as follows:
//   - WASAPI
//   - DirectSound
//   - WinMM
//   - ALSA
//   - OSS
//   - OpenSL|ES
//   - OpenAL
//   - Null
//
// This will dynamically load backends DLLs/SOs (such as dsound.dll).
func (d *Device) ContextInit(backends []Backend, logging bool) error {
	cbackends := (*C.mal_backend)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&backends)).Data))
	cbackendcount := (C.mal_uint32)(len(backends))

	var ret C.mal_result
	if logging {
		config := C.mal_context_config_init((*[0]byte)(C.goRecvCallback))
		ret = C.mal_context_init(cbackends, cbackendcount, &config, d.context)
	} else {
		config := C.mal_context_config_init(nil)
		ret = C.mal_context_init(cbackends, cbackendcount, &config, d.context)
	}

	v := (Result)(ret)
	return errorFromResult(v)
}

// ContextUninit uninitializes a context.
//
// This will unload the backend DLLs/SOs.
func (d *Device) ContextUninit() error {
	ret := C.mal_context_uninit(d.context)
	v := (Result)(ret)
	return errorFromResult(v)
}

// EnumerateDevices enumerates over each device of the given type (Playback or Capture).
//
// It is _not_ safe to assume the first enumerated device is the default device.
//
// This API dynamically links to backend DLLs/SOs (such as dsound.dll).
func (d *Device) EnumerateDevices(kind DeviceType) ([]DeviceInfo, error) {
	var pcount uint32 = 32
	pinfo := make([]C.mal_device_info, pcount)

	ckind := (C.mal_device_type)(kind)
	cpcount := (*C.mal_uint32)(unsafe.Pointer(&pcount))
	cpinfo := (*C.mal_device_info)(unsafe.Pointer(&pinfo[0]))

	ret := C.mal_enumerate_devices(d.context, ckind, cpcount, cpinfo)
	v := (Result)(ret)

	if v == Success {
		tmp := (*[1 << 20]C.mal_device_info)(unsafe.Pointer(cpinfo))[:pcount]
		res := make([]DeviceInfo, pcount)
		for i, d := range tmp {
			res[i] = NewDeviceInfoFromPointer(unsafe.Pointer(&d))
		}

		return res, nil
	}

	return nil, errorFromResult(v)
}

// Init initializes a device.
//
// The device ID (pdeviceid) can be nil, in which case the default device is used. Otherwise, you
// can retrieve the ID by calling EnumerateDevices() and use the ID from the returned data.
//
// Set pdeviceid to nil to use the default device. Do _not_ rely on the first device ID returned
// by EnumerateDevices() to be the default device.
//
// Consider using ConfigInit(), ConfigInitPlayback(), etc. to make it easier
// to initialize a DeviceConfig object.
func (d *Device) Init(kind DeviceType, pdeviceid *DeviceID, pconfig *DeviceConfig) error {
	ckind := (C.mal_device_type)(kind)
	cpdeviceid := pdeviceid.cptr()
	cpconfig := pconfig.cptr()

	ret := C.mal_device_init(d.context, ckind, cpdeviceid, cpconfig, nil, d.device)
	v := (Result)(ret)
	return errorFromResult(v)
}

// Uninit uninitializes a device.
//
// This will explicitly stop the device. You do not need to call Stop() beforehand, but it's harmless if you do.
func (d *Device) Uninit() {
	C.mal_device_uninit(d.device)
}

// SetRecvCallback sets the callback to use when the application has received data from the device.
func (d *Device) SetRecvCallback(proc RecvProc) {
	recvHandler = proc
	C.goSetRecvCallback(d.device)
}

// SetSendCallback sets the callback to use when the application needs to send data to the device for playback.
func (d *Device) SetSendCallback(proc SendProc) {
	sendHandler = proc
	C.goSetSendCallback(d.device)
}

// SetStopCallback sets the callback to use when the device has stopped, either explicitly or as a result of an error.
func (d *Device) SetStopCallback(proc StopProc) {
	stopHandler = proc
	C.goSetStopCallback(d.device)
}

// Start activates the device. For playback devices this begins playback. For capture devices it begins recording.
//
// For a playback device, this will retrieve an initial chunk of audio data from the client before
// returning. The reason for this is to ensure there is valid audio data in the buffer, which needs
// to be done _before_ the device begins playback.
func (d *Device) Start() error {
	ret := C.mal_device_start(d.device)
	v := (Result)(ret)
	return errorFromResult(v)
}

// Stop puts the device to sleep, but does not uninitialize it. Use Start() to start it up again.
func (d *Device) Stop() error {
	ret := C.mal_device_stop(d.device)
	v := (Result)(ret)
	return errorFromResult(v)
}

// IsStarted determines whether or not the device is started.
func (d *Device) IsStarted() (r bool) {
	ret := C.mal_device_is_started(d.device)
	v := (uint32)(ret)
	if v == 1 {
		r = true
	}
	return r
}

// BufferSizeInBytes retrieves the size of the buffer in bytes.
func (d *Device) BufferSizeInBytes() uint32 {
	ret := C.mal_device_get_buffer_size_in_bytes(d.device)
	v := (uint32)(ret)
	return v
}

// SampleSizeInBytes retrieves the size of a sample in bytes for the given format.
func (d *Device) SampleSizeInBytes(format FormatType) uint32 {
	cformat := (C.mal_format)(format)
	ret := C.mal_get_sample_size_in_bytes(cformat)
	v := (uint32)(ret)
	return v
}

// ConfigInit is a helper function for initializing a DeviceConfig object.
func (d *Device) ConfigInit(format FormatType, channels uint32, samplerate uint32, onrecvcallback RecvProc, onsendcallback SendProc) DeviceConfig {
	cformat := (C.mal_format)(format)
	cchannels := (C.mal_uint32)(channels)
	csamplerate := (C.mal_uint32)(samplerate)

	recvHandler = onrecvcallback
	sendHandler = onsendcallback

	ret := C.goConfigInit(cformat, cchannels, csamplerate)
	v := NewDeviceConfigFromPointer(unsafe.Pointer(&ret))
	return v
}

// ConfigInitCapture is a simplified version of DeviceConfigInit() for capture devices.
func (d *Device) ConfigInitCapture(format FormatType, channels uint32, samplerate uint32, onrecvcallback RecvProc) DeviceConfig {
	cformat := (C.mal_format)(format)
	cchannels := (C.mal_uint32)(channels)
	csamplerate := (C.mal_uint32)(samplerate)

	recvHandler = onrecvcallback

	ret := C.goConfigInitCapture(cformat, cchannels, csamplerate)
	v := NewDeviceConfigFromPointer(unsafe.Pointer(&ret))
	return v
}

// ConfigInitPlayback is a simplified version of DeviceConfigInit() for playback devices.
func (d *Device) ConfigInitPlayback(format FormatType, channels uint32, samplerate uint32, onsendcallback SendProc) DeviceConfig {
	cformat := (C.mal_format)(format)
	cchannels := (C.mal_uint32)(channels)
	csamplerate := (C.mal_uint32)(samplerate)

	sendHandler = onsendcallback

	ret := C.goConfigInitPlayback(cformat, cchannels, csamplerate)
	v := NewDeviceConfigFromPointer(unsafe.Pointer(&ret))
	return v
}

// BufferSizeInFrames retrieves the size of the buffer in frames.
func (d *Device) BufferSizeInFrames() uint32 {
	return uint32(d.device.bufferSizeInFrames)
}

// Type returns device type.
func (d *Device) Type() DeviceType {
	return DeviceType(d.device._type)
}

// Format returns device format.
func (d *Device) Format() FormatType {
	return FormatType(d.device.format)
}

// Channels returns number of channels.
func (d *Device) Channels() uint32 {
	return uint32(d.device.channels)
}

// SampleRate returns sample rate.
func (d *Device) SampleRate() uint32 {
	return uint32(d.device.sampleRate)
}
