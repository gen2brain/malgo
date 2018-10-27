// Package mal - Mini audio library (mini_al cgo bindings).
package mal

/*
#cgo CFLAGS: -std=gnu99
#cgo mal_debug CFLAGS: -DMAL_DEBUG_OUTPUT=1

#cgo linux LDFLAGS: -ldl -lpthread -lm
#cgo openbsd LDFLAGS: -lpthread -lm -lossaudio
#cgo netbsd LDFLAGS: -lpthread -lm -lossaudio
#cgo freebsd LDFLAGS: -lpthread -lm
#cgo android LDFLAGS: -lOpenSLES

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMAL_NO_SSE2 -DMAL_NO_AVX -DMAL_NO_AVX512 -DMAL_NO_NEON

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
mal_device_config goConfigInitDefaultCapture();
mal_device_config goConfigInitDefaultPlayback();

mal_context_config goContextConfigInit();
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

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
)

// Errors.
var (
	errTag string = "mini_al"

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
	Idle     ThreadPriority = -5
	Lowest   ThreadPriority = -4
	Low      ThreadPriority = -3
	Normal   ThreadPriority = -2
	High     ThreadPriority = -1
	Highest  ThreadPriority = -0
	Realtime ThreadPriority = -1
	Default  ThreadPriority = -2
)

// Result type.
type Result int32

// Device type.
type Device struct {
	context *C.mal_context
	device  *C.mal_device
}

// NewDevice returns new Device.
func NewDevice() *Device {
	d := &Device{}
	d.context = C.goGetContext()
	d.device = C.goGetDevice()
	return d
}

// DeviceID type.
type DeviceID [unsafe.Sizeof(C.mal_device_id{})]byte

// cptr return C pointer.
func (d *DeviceID) cptr() *C.mal_device_id {
	return (*C.mal_device_id)(unsafe.Pointer(d))
}

// DeviceInfo type.
type DeviceInfo struct {
	ID            DeviceID
	Name          [256]byte
	FormatCount   uint32
	Formats       [6]uint32
	MinChannels   uint32
	MaxChannels   uint32
	MinSampleRate uint32
	MaxSampleRate uint32
}

// String returns string.
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("{ID: %s, Name: %s}", string(d.ID[:]), string(d.Name[:]))
}

// NewDeviceInfoFromPointer returns new DeviceInfo from pointer.
func NewDeviceInfoFromPointer(ptr unsafe.Pointer) DeviceInfo {
	return *(*DeviceInfo)(ptr)
}

// AlsaDeviceConfig type.
type AlsaDeviceConfig struct {
	NoMMap uint32
}

// PulseDeviceConfig type.
type PulseDeviceConfig struct {
	StreamName *byte
}

// DeviceConfig type.
type DeviceConfig struct {
	Format             FormatType
	Channels           uint32
	SampleRate         uint32
	ChannelMap         [32]byte
	BufferSizeInFrames uint32
	Periods            uint32
	ShareMode          ShareMode
	PerformanceProfile PerformanceProfile
	_                  [4]byte
	OnRecvCallback     *[0]byte
	OnSendCallback     *[0]byte
	OnStopCallback     *[0]byte
	Alsa               AlsaDeviceConfig
	_                  [4]byte
	Pulse              PulseDeviceConfig
}

// cptr return C pointer.
func (d *DeviceConfig) cptr() *C.mal_device_config {
	return (*C.mal_device_config)(unsafe.Pointer(d))
}

// NewDeviceConfigFromPointer returns new DeviceConfig from pointer.
func NewDeviceConfigFromPointer(ptr unsafe.Pointer) DeviceConfig {
	return *(*DeviceConfig)(ptr)
}

// AlsaContextConfig type.
type AlsaContextConfig struct {
	UseVerboseDeviceEnumeration uint32
}

// PulseContextConfig type.
type PulseContextConfig struct {
	PApplicationName *byte
	PServerName      *byte
	// Enables autospawning of the PulseAudio daemon if necessary.
	TryAutoSpawn uint32
	Pad_cgo_0    [4]byte
}

// JackContextConfig type.
type JackContextConfig struct {
	PClientName    *byte
	TryStartServer uint32
	Pad_cgo_0      [4]byte
}

// ContextConfig type.
type ContextConfig struct {
	OnLog          *[0]byte
	ThreadPriority ThreadPriority
	Alsa           AlsaContextConfig
	Pulse          PulseContextConfig
	Jack           JackContextConfig
}

// cptr return C pointer.
func (d *ContextConfig) cptr() *C.mal_context_config {
	return (*C.mal_context_config)(unsafe.Pointer(d))
}

// NewContextConfigFromPointer returns new ContextConfig from pointer.
func NewContextConfigFromPointer(ptr unsafe.Pointer) ContextConfig {
	return *(*ContextConfig)(ptr)
}

// RecvProc type.
type RecvProc func(framecount uint32, psamples []byte)

// SendProc type.
type SendProc func(framecount uint32, psamples []byte) uint32

// StopProc type.
type StopProc func()

// LogProc type.
type LogProc func(message string)

// Handlers.
var (
	recvHandler RecvProc
	sendHandler SendProc
	stopHandler StopProc
	logHandler  LogProc
)

//export goRecvCallback
func goRecvCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) {
	if recvHandler != nil {
		sampleCount := uint32(frameCount) * uint32(pDevice.channels)
		sizeInBytes := uint32(C.mal_get_bytes_per_sample(pDevice.format))
		psamples := (*[1 << 20]byte)(pSamples)[0 : sampleCount*sizeInBytes]
		recvHandler(uint32(frameCount), psamples)
	}
}

//export goSendCallback
func goSendCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) (r C.mal_uint32) {
	if sendHandler != nil {
		sampleCount := uint32(frameCount) * uint32(pDevice.channels)
		sizeInBytes := uint32(C.mal_get_bytes_per_sample(pDevice.format))
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
	if logHandler != nil {
		logHandler(C.GoString(message))
	}
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
func (d *Device) ContextInit(backends []Backend, config ContextConfig) error {
	cbackends := (*C.mal_backend)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&backends)).Data))
	cbackendcount := (C.mal_uint32)(len(backends))
	cconfig := config.cptr()

	ret := C.mal_context_init(cbackends, cbackendcount, cconfig, d.context)
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

// Devices retrieves basic information about every active playback or capture device.
func (d *Device) Devices(kind DeviceType) ([]DeviceInfo, error) {
	var pcount uint32
	var ccount uint32

	pinfo := make([]*C.mal_device_info, 32)
	cinfo := make([]*C.mal_device_info, 32)

	cpcount := (*C.mal_uint32)(unsafe.Pointer(&pcount))
	cccount := (*C.mal_uint32)(unsafe.Pointer(&ccount))

	cpinfo := (**C.mal_device_info)(unsafe.Pointer(&pinfo[0]))
	ccinfo := (**C.mal_device_info)(unsafe.Pointer(&cinfo[0]))

	ret := C.mal_context_get_devices(d.context, cpinfo, cpcount, ccinfo, cccount)
	v := (Result)(ret)

	if v == Success {
		res := make([]DeviceInfo, 0)

		if kind == Playback {
			tmp := (*[1 << 20]*C.mal_device_info)(unsafe.Pointer(cpinfo))[:pcount]
			for _, d := range tmp {
				if d != nil {
					res = append(res, NewDeviceInfoFromPointer(unsafe.Pointer(d)))
				}
			}
		} else if kind == Capture {
			tmp := (*[1 << 20]*C.mal_device_info)(unsafe.Pointer(ccinfo))[:ccount]
			for _, d := range tmp {
				if d != nil {
					res = append(res, NewDeviceInfoFromPointer(unsafe.Pointer(d)))
				}
			}
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

// SetLogCallback sets the log callback.
func (d *Device) SetLogCallback(proc LogProc) {
	logHandler = proc
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

// ConfigInitDefaultCapture initializes a default capture device config.
func (d *Device) ConfigInitDefaultCapture(onrecvcallback RecvProc) DeviceConfig {
	recvHandler = onrecvcallback

	ret := C.goConfigInitDefaultCapture()
	v := NewDeviceConfigFromPointer(unsafe.Pointer(&ret))
	return v
}

// ConfigInitDefaultPlayback initializes a default playback device config.
func (d *Device) ConfigInitDefaultPlayback(onsendcallback SendProc) DeviceConfig {
	sendHandler = onsendcallback

	ret := C.goConfigInitDefaultPlayback()
	v := NewDeviceConfigFromPointer(unsafe.Pointer(&ret))
	return v
}

// ContextConfigInit is a helper function for initializing a ContextConfig object.
func (d *Device) ContextConfigInit(onlogcallback LogProc) ContextConfig {
	logHandler = onlogcallback

	ret := C.goContextConfigInit()
	v := NewContextConfigFromPointer(unsafe.Pointer(&ret))
	return v
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
	ret := C.mal_get_bytes_per_sample(cformat)
	v := (uint32)(ret)
	return v
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
