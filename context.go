package malgo

/*
 #include "malgo.h"
#include "callbacks.h"
static void setGoContextConfigCallbacks(ma_context_config* pConfig, ma_bool8 log) {
	if(log) {
		pConfig->logCallback = goLogCallbackWrapper;
	}
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

// LogProc type.
type LogProc func(ctx Context, device Device, logLevel int, message string)
// ContextConfig type.
type ContextConfig struct {
	LogCallback LogProc
	ThreadPriority  int
	ThreadStackSize int
	Alsa struct {
		UseVerboseDeviceEnumeration bool
	}
	Pulse struct {
		ApplicationName string
		ServerName      string
		// Enables autospawning of the PulseAudio daemon if necessary.
		TryAutoSpawn bool
	}
	CoreAudio struct {
		SessionCategory                                  int
		SessionCategoryOptions                           int
		NoAudioSessionActivate, NoAudioSessionDeactivate bool
	}
	Jack struct {
		ClientName     string
		TryStartServer bool
	}
}
type internalContextInfo struct {
	LogCallback     LogProc
	memory pointerList
}


// Initializes a new ContextConfig with defaults. You must call this instead of creating a ContextConfig object directly.
func NewContextConfig() *ContextConfig {
	cConfig := C.ma_context_config_init()
	config := &ContextConfig{}
	config.Alsa.UseVerboseDeviceEnumeration = intToBool(cConfig.alsa.useVerboseDeviceEnumeration)

	config.Pulse.ApplicationName = goString(cConfig.pulse.pApplicationName)
	config.Pulse.ServerName = goString(cConfig.pulse.pServerName)
	config.Pulse.TryAutoSpawn = intToBool(cConfig.pulse.tryAutoSpawn)

	config.CoreAudio.SessionCategory = int(cConfig.coreaudio.sessionCategory)
	config.CoreAudio.SessionCategoryOptions = int(cConfig.coreaudio.sessionCategoryOptions)
	config.CoreAudio.NoAudioSessionActivate = intToBool(cConfig.coreaudio.noAudioSessionActivate)
	config.CoreAudio.NoAudioSessionDeactivate = intToBool(cConfig.coreaudio.noAudioSessionDeactivate)

	config.Jack.ClientName = goString(cConfig.jack.pClientName)
	config.Jack.TryStartServer = intToBool(cConfig.jack.tryStartServer)
	config.ThreadPriority = int(cConfig.threadPriority)
	config.ThreadStackSize = int(cConfig.threadStackSize)
	return config
}

// converts this config to a C.ma_context_config for use in ma_context_init
func (self *ContextConfig) toCRepr() (C.ma_context_config, pointerList) {
	// even if we forget to initialize some fields, telling Miniaudio to reinitialize the config ensures it doesn't break anything
	config := C.ma_context_config_init()
	var memory pointerList
	// if the user set callbacks to a value other than nil, we give miniaudio our Go wrapper callbacks. More efficient / correct than our callback checking if the user-provided callback exists every time it's called.
	config.threadPriority = C.ma_thread_priority(self.ThreadPriority)
	config.threadStackSize = C.size_t(self.ThreadStackSize)
	// use a C function to set the callbacks so the C compiler has a chance to check if the callback types match
	C.setGoContextConfigCallbacks(&config, boolToInt8(self.LogCallback != nil))
	config.alsa.useVerboseDeviceEnumeration = boolToInt(self.Alsa.UseVerboseDeviceEnumeration)

	config.pulse.pApplicationName = memory.cString(self.Pulse.ApplicationName)
	config.pulse.pServerName = memory.cString(self.Pulse.ServerName)
	config.pulse.tryAutoSpawn = boolToInt(self.Pulse.TryAutoSpawn)

	config.coreaudio.sessionCategory = C.ma_ios_session_category(self.CoreAudio.SessionCategory)
	config.coreaudio.sessionCategoryOptions = C.ma_uint32(self.CoreAudio.SessionCategoryOptions)
	config.coreaudio.noAudioSessionActivate = boolToInt(self.CoreAudio.NoAudioSessionActivate)
	config.coreaudio.noAudioSessionDeactivate = boolToInt(self.CoreAudio.NoAudioSessionDeactivate)

	config.jack.pClientName = memory.cString(self.Jack.ClientName)
	config.jack.tryStartServer = boolToInt(self.Jack.TryStartServer)
	return config, memory
}

// for the Context, we handle callbacks with a map[*C.ma_context]*internalContextInfo and a mutex. Because the Context is a C object. it's safe to pass it between C and go. Our Go callback wrapper functions will then get the context info via this map and call the right callback
var (
	contextConfigs      = make(map[*C.ma_context]*internalContextInfo, 3)
	contextConfigsMutex sync.Mutex
)

// The Context is a uintptr because CGO rules disallow letting C code hang on to pointers to Go memory. So we allocate Context in the C heap, allowing both C and Go to do stuff with it.

// Context is used for selecting and initializing the relevant backends.
type Context uintptr

// DefaultContext is an unspecified context. It can be used to initialize a streaming
// function with implicit context defaults.
const DefaultContext Context = 0

func (ctx Context) cptr() *C.ma_context {
	// 0 is sometimes a valid value for contexts, so those cases where it is will just have to check
	if ctx == 0 {
		panic("Malgo: context must not be nil")
	}
	return (*C.ma_context)(unsafe.Pointer(ctx))
}

// Uninit uninitializes a context.
// Errors are undefined if you call this while any device created by this context is still active.
func (ctx Context) Uninit() error {
	result := C.ma_context_uninit(ctx.cptr())
	// free the Context underlying C memory, remove it from contextConfigs, and free the C versions of the Go strings that were set in ContextConfig
	C.ma_free(unsafe.Pointer(ctx), nil)
	contextConfigsMutex.Lock()
	config, exists := contextConfigs[ctx.cptr()]
	delete(contextConfigs, ctx.cptr())
	contextConfigsMutex.Unlock()
	if exists {
		config.memory.free()
	}
	return errorFromResult(result)
}
func deviceInfosFromPointer(ptr *C.ma_device_info, length C.ma_uint32) []DeviceInfo {
	infos := make([]DeviceInfo, length)
	deviceInfoAddr := uintptr(unsafe.Pointer(ptr))
	for i := 0; i < int(length); i++ {
		deviceInfoAddr := deviceInfoAddr + (uintptr(i) * rawDeviceInfoSize)
		info := *(*C.ma_device_info)(unsafe.Pointer(deviceInfoAddr))
		infos[i] = deviceInfoFromCRepr(info)
	}
	return infos
}

// PlaybackDevices retrieves basic information about every active playback device.
func (ctx Context) PlaybackDevices() ([]DeviceInfo, error) {
	contextMutex.Lock()
	defer contextMutex.Unlock()
	var devices *C.ma_device_info
	var deviceCount C.ma_uint32
	result := C.ma_context_get_devices(ctx.cptr(), &devices, &deviceCount, nil, nil)
	err := errorFromResult(result)
	if err != nil {
		return nil, err
	}
	return deviceInfosFromPointer(devices, deviceCount), nil
}

// CaptureDevices retrieves basic information about every active Capture device.
func (ctx Context) CaptureDevices() ([]DeviceInfo, error) {
	contextMutex.Lock()
	defer contextMutex.Unlock()
	var devices *C.ma_device_info
	var deviceCount C.ma_uint32
	result := C.ma_context_get_devices(ctx.cptr(), nil, nil, &devices, &deviceCount)
	err := errorFromResult(result)
	if err != nil {
		return nil, err
	}
	return deviceInfosFromPointer(devices, deviceCount), nil
}

// AllDevices returns basic info about every playback or capture device
func (ctx Context) AllDevices() ([]DeviceInfo, []DeviceInfo, error) {
	contextMutex.Lock()
	defer contextMutex.Unlock()
	var playbackDevices, captureDevices *C.ma_device_info
	var playbackDeviceCount, captureDeviceCount C.ma_uint32
	result := C.ma_context_get_devices(ctx.cptr(), &playbackDevices, &playbackDeviceCount, &captureDevices, &captureDeviceCount)
	err := errorFromResult(result)
	if err != nil {
		return nil, nil, err
	}
	return deviceInfosFromPointer(playbackDevices, playbackDeviceCount), deviceInfosFromPointer(captureDevices, captureDeviceCount), nil
}

// DeviceInfo retrieves information about a device of the given type, with the specified ID and share mode.
func (ctx Context) DeviceInfo(kind DeviceType, id DeviceID, mode ShareMode) (DeviceInfo, error) {
	var info C.ma_device_info

	result := C.ma_context_get_device_info(ctx.cptr(), C.ma_device_type(kind), id.cptr(), C.ma_share_mode(mode), &info)
	err := errorFromResult(result)
	if err != nil {
		return DeviceInfo{}, err
	}
	return deviceInfoFromCRepr(info), nil
}

// returns weather this context supports loopback devices
func (ctx Context) IsLoopbackSupported() bool {
	return intToBool(C.ma_context_is_loopback_supported(ctx.cptr()))
}

// returns the backend that this context is using
func (ctx Context) Backend() Backend {
	return Backend(ctx.cptr().backend)
}

var contextMutex sync.Mutex

//export goLogCallback
func goLogCallback(pContext *C.ma_context, pDevice *C.ma_device, logLevel C.ma_uint32, message *C.char) {
	contextConfigsMutex.Lock()
	callback := contextConfigs[pContext].LogCallback
	contextConfigsMutex.Unlock()
	callback(Context(unsafe.Pointer(pContext)), Device(unsafe.Pointer(pDevice)), int(logLevel), C.GoString(message))
}

// InitContext creates and initializes a context.
// When the application no longer needs the context instance, it needs to call Uninit() .
func InitContext(backends []Backend, config *ContextConfig) (Context, error) {
	ctx := Context(C.ma_malloc(C.sizeof_ma_context, nil))
	if ctx == 0 {
		return 0, ErrOutOfMemory
	}
	var cConfigPtr *C.ma_context_config
	if config != nil {
		cConfig, memory := config.toCRepr()
		cConfigPtr = &cConfig
		contextConfigsMutex.Lock()
		contextConfigs[ctx.cptr()] = &internalContextInfo{LogCallback: config.LogCallback, memory: memory}
		contextConfigsMutex.Unlock()
	}
	var backendsArg *C.ma_backend

	if len(backends) > 0 {
		cBackends := make([]C.ma_backend, len(backends))
		for pos, i := range backends {
			cBackends[pos] = C.ma_backend(backends[i])
		}
		backendsArg = &cBackends[0]
	}
	result := C.ma_context_init(backendsArg, C.ma_uint32(len(backends)), cConfigPtr, ctx.cptr())
	err := errorFromResult(result)
	if err != nil {
		C.ma_free(unsafe.Pointer(ctx), nil)
		return 0, err
	}
	return ctx, nil
}
