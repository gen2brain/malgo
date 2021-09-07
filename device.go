package malgo

/*
 #include "malgo.h"
#include "callbacks.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// DataProc type.
type DataProc func(device Device, output, input *DataBuffer, framecount int)

// StopProc type.
type StopProc func(device Device)
// embedded into DeviceConfig and stores info that we can't just convert to a C type and forgetabout
type internalDeviceInfo struct {
	DataCallback            DataProc
	StopCallback            StopProc
	// stores pointers to any C memory allocated by Malgo, including DeviceID
	memory        pointerList
	// allocation space for the input and output DataBuffer structs so Go doesn't have a chance to allocate them on the heap every time the data callback get called
	input, output DataBuffer
}

// Device represents a streaming instance.
type Device uintptr

// InitDevice initializes a device.
//
// The device ID can be nil, in which case the default device is used. Otherwise, you
// can retrieve the ID by calling Context.Devices() and use the ID from the returned data.
//
// Set device ID to nil to use the default device. Do _not_ rely on the first device ID returned
// by Context.Devices() to be the default device.
//
// The returned instance has to be cleaned up using Uninit().
func InitDevice(context Context, config DeviceConfig) (Device, error) {
	dev := Device(C.ma_aligned_malloc(C.sizeof_ma_device, simdAlignment, nil))
	if dev == 0 {
		return 0, ErrOutOfMemory
	}

	rawDevice := dev.cptr()
	cConfig, memory := config.toCRepr()
	result := C.ma_device_init((*C.ma_context)(unsafe.Pointer(context)), &cConfig, rawDevice)
	if result != C.MA_SUCCESS {
		C.ma_aligned_free(unsafe.Pointer(dev), nil)
		return 0, errorFromResult(result)
	}
	deviceConfigsMutex.Lock()
	deviceConfigs[rawDevice] = &internalDeviceInfo{DataCallback: config.DataCallback, StopCallback: config.StopCallback, memory: memory}
	deviceConfigsMutex.Unlock()

	return dev, nil
}

func (dev Device) cptr() *C.ma_device {
	if dev == 0 {
		panic("Malgo: device must not be nil")
	}
	return (*C.ma_device)(unsafe.Pointer(dev))
}

// Type returns device type.
func (dev Device) Type() DeviceType {
	return DeviceType(dev.cptr()._type)
}

// PlaybackFormat returns device playback format.
func (dev Device) PlaybackFormat() Format {
	return Format(dev.cptr().playback.format)
}

// CaptureFormat returns device capture format.
func (dev Device) CaptureFormat() Format {
	return Format(dev.cptr().capture.format)
}

// PlaybackChannels returns number of playback channels.
func (dev Device) PlaybackChannels() int {
	return int(dev.cptr().playback.channels)
}

// CaptureChannels returns number of capture channels.
func (dev Device) CaptureChannels() int {
	return int(dev.cptr().capture.channels)
}

// SampleRate returns sample rate.
func (dev Device) SampleRate() int {
	return int(dev.cptr().sampleRate)
}

// Start activates the device.
// For playback devices this begins playback. For capture devices it begins recording.
//
// For a playback device, this will retrieve an initial chunk of audio data from the client before
// returning. The reason for this is to ensure there is valid audio data in the buffer, which needs
// to be done _before_ the device begins playback.
//
// This API waits until the backend device has been started for real by the worker thread. It also
// waits on a mutex for thread-safety.
func (dev Device) Start() error {
	result := C.ma_device_start(dev.cptr())
	return errorFromResult(result)
}

// IsStarted determines whether or not the device is started.
func (dev Device) IsStarted() bool {
	result := C.ma_device_is_started(dev.cptr())
	return result == C.MA_TRUE
}

// Stop puts the device to sleep, but does not uninitialize it. Use Start() to start it up again.
//
// This API needs to wait on the worker thread to stop the backend device properly before returning. It
// also waits on a mutex for thread-safety. In addition, some backends need to wait for the device to
// finish playback/recording of the current fragment which can take some time (usually proportionate to
// the buffer size that was specified at initialization time).
func (dev Device) Stop() error {
	result := C.ma_device_stop(dev.cptr())
	return errorFromResult(result)
}

// Uninit uninitializes a device.
//
// This will explicitly stop the device. You do not need to call Stop() beforehand, but it's
// harmless if you do.
func (dev Device) Uninit() {
	rawDevice := dev.cptr()
	C.ma_device_uninit(rawDevice)
	C.ma_aligned_free(unsafe.Pointer(dev), nil)
	deviceConfigsMutex.Lock()
	config, exists := deviceConfigs[rawDevice]
	delete(deviceConfigs, rawDevice)
	deviceConfigsMutex.Unlock()
	if exists {
		// free any malgo allocated C memory associated with this device
		config.memory.free()
	}
}

// changes this devices data callback
// it's safe to call this while the device is running, as we use a mutex
// do not pass a nil data callback
func (dev Device) SetDataCallback(dataCallback DataProc) {
	if dataCallback == nil {
		panic("Malgo: dataCallback cannot be nil")
	}
	deviceConfigsMutex.Lock()
	deviceConfigs[dev.cptr()].DataCallback = dataCallback
	deviceConfigsMutex.Unlock()
}

// changes this devices stop callback
// it's safe to call this while the device is running, as we use a mutex
// do not pass a nil stop callback
func (dev Device) SetStopCallback(stopCallback StopProc) {
	if stopCallback == nil {
		panic("Malgo: dataCallback cannot be nil")
	}
	deviceConfigsMutex.Lock()
	deviceConfigs[dev.cptr()].StopCallback = stopCallback
	deviceConfigsMutex.Unlock()
}

var deviceConfigsMutex sync.Mutex
var deviceConfigs = make(map[*C.ma_device]*internalDeviceInfo)

//export goDataCallback
func goDataCallback(pDevice *C.ma_device, pOutput, pInput unsafe.Pointer, frameCount C.ma_uint32, outputSizeInBytes, inputSizeInBytes C.ma_uint32) {
	deviceConfigsMutex.Lock()
	config := deviceConfigs[pDevice]
	deviceConfigsMutex.Unlock()
	var inputSamples, outputSamples *DataBuffer
	// just in-case Go allocates our DataBuffer structs on the heap, we store them inside the Context
	if pOutput != nil {
		config.output = DataBuffer{data: pOutput, length: int(outputSizeInBytes), format: Format(pDevice.playback.format)}
		outputSamples = &config.output
	}
	if pInput != nil {
		config.input = DataBuffer{data: pInput, length: int(inputSizeInBytes), format: Format(pDevice.capture.format)}
		inputSamples = &config.input
	}
	config.DataCallback(Device(unsafe.Pointer(pDevice)), outputSamples, inputSamples, int(frameCount))
}

//export goStopCallback
func goStopCallback(pDevice *C.ma_device) {
	deviceConfigsMutex.Lock()
	callback := deviceConfigs[pDevice].StopCallback
	deviceConfigsMutex.Unlock()

	callback(Device(unsafe.Pointer(pDevice)))
}
