#include "_cgo_export.h"

#define MINIAUDIO_IMPLEMENTATION
#include "miniaudio.h"

static void goLogCallbackWrapper(ma_context *pContext, ma_device *pDevice, 
                                 ma_uint32 logLevel, const char *message) {
    goLogCallback(pContext, pDevice, (char *)message);
}

void goSetContextConfigCallbacks(ma_context_config* pConfig) {
    pConfig->logCallback = goLogCallbackWrapper;
}

static void goDataCallbackWrapper(ma_device *pDevice,
                                  void *pOutput, const void *pInput,
                                  ma_uint32 frames)
{
    goDataCallback(pDevice, pOutput, (void *)pInput, frames);
}

void goSetDeviceConfigCallbacks(ma_device_config* pConfig) {
    pConfig->dataCallback = goDataCallbackWrapper;
    pConfig->stopCallback = goStopCallback;
}
