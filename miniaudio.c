#include "_cgo_export.h"

#define MINIAUDIO_IMPLEMENTATION
#include "miniaudio.h"

static void goLogCallbackWrapper(void* pUserData, ma_uint32 logLevel, const char *message) {
    goLogCallback((ma_context*) pUserData, (char *)message);
}

// Note that the context in the argument has not been initialized here. We will only use the pointer as pUserData
// so that we could easily identify which context a log callback belongs to.
void goSetContextConfigCallbacks(ma_context_config* pConfig, ma_context* pContext) {
    ma_log* log = malloc(sizeof(ma_log));
    ma_log_init(NULL, log); // TODO: Set allocation callback?
    ma_log_register_callback(log, ma_log_callback_init(goLogCallbackWrapper, (void*)pContext));
    pConfig->pLog = log;
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
