#include "_cgo_export.h"

#define MINI_AL_IMPLEMENTATION
#include "mini_al.h"

static void goLogCallbackWrapper(mal_context *pContext, mal_device *pDevice, const char *message) {
    goLogCallback(pContext, pDevice, (char *)message);
}

void goSetContextConfigCallbacks(mal_context_config* pConfig) {
    pConfig->onLog = goLogCallbackWrapper;
}

static void goRecvCallbackWrapper(mal_device *pDevice, mal_uint32 frames, const void *pSamples) {
    goRecvCallback(pDevice, frames, (void *)pSamples);
}

void goSetDeviceConfigCallbacks(mal_device_config* pConfig) {
    pConfig->onRecvCallback = goRecvCallbackWrapper;
    pConfig->onSendCallback = goSendCallback;
    pConfig->onStopCallback = goStopCallback;
}
