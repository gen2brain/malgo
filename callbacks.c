#include "_cgo_export.h"
#include "miniaudio.h"
#include "callbacks.h"
extern void goLogCallbackWrapper(ma_context* pContext, ma_device* pDevice, ma_uint32 logLevel, const char* message) {
	goLogCallback(pContext, pDevice, (char*)(message)); // cast to remove const qualifier that  Go doesn't know how to handle
}
extern void goDataCallbackWrapper(ma_device* pDevice, void* pOutput, const void* pInput, ma_uint32 frameCount) {
	// calculate the buffers size here to avoid C > Go > C calls
	ma_uint64 outputSizeInBytes, inputSizeInBytes;
	ma_uint32 outputSampleSizeInBytes, inputSampleSizeInBytes;
	outputSampleSizeInBytes = 0;
	inputSampleSizeInBytes = 0;
	outputSizeInBytes = 0;
	inputSizeInBytes = 0;
	if (pOutput != NULL) {
		ma_uint32 sampleCount = frameCount * pDevice->playback.channels;
		outputSampleSizeInBytes = ma_get_bytes_per_sample(pDevice->playback.format);
		outputSizeInBytes = (ma_uint64)(outputSampleSizeInBytes * sampleCount);
	}
	if (pInput != NULL) {
		ma_uint32 sampleCount = frameCount * pDevice->capture.channels;
inputSampleSizeInBytes = ma_get_bytes_per_sample(pDevice->capture.format);
		inputSizeInBytes = (ma_uint64)(inputSampleSizeInBytes * sampleCount);
	}
	goDataCallback(pDevice, pOutput, (void*)(pInput), frameCount, outputSizeInBytes, inputSizeInBytes);
}
extern void goStopCallbackWrapper(ma_device* pDevice) {
	goStopCallback(pDevice);
}