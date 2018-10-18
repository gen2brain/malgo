#include "_cgo_export.h"

#define MINI_AL_IMPLEMENTATION
#include "mini_al.h"

void goSetContextConfigCallbacks(mal_context_config* pConfig) {
    pConfig->onLog = goLogCallback;
}

void goSetDeviceConfigCallbacks(mal_device_config* pConfig) {
    pConfig->onRecvCallback = goRecvCallback;
    pConfig->onSendCallback = goSendCallback;
    pConfig->onStopCallback = goStopCallback;
}
