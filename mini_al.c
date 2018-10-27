#include "_cgo_export.h"

#define MINI_AL_IMPLEMENTATION
#include "mini_al.h"

mal_context context;
mal_device device;

static void goRecvCallbackWrapper(mal_device *pDevice, mal_uint32 frames, const void *pSamples) {
    goRecvCallback(pDevice, frames, (void *)pSamples);
}

static void goLogCallbackWrapper(mal_context *pContext, mal_device *pDevice, const char *message) {
    goLogCallback(pContext, pDevice, (char *)message);
}

void goSetRecvCallback(mal_device* pDevice) {
    mal_device_set_recv_callback(pDevice, goRecvCallbackWrapper);
}

void goSetSendCallback(mal_device* pDevice) {
    mal_device_set_send_callback(pDevice, goSendCallback);
}

void goSetStopCallback(mal_device* pDevice) {
    mal_device_set_stop_callback(pDevice, goStopCallback);
}

mal_device* goGetDevice() {
    return &device;
}

mal_context* goGetContext() {
    return &context;
}

mal_device_config goConfigInit(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, goRecvCallbackWrapper, goSendCallback);
}

mal_device_config goConfigInitCapture(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, goRecvCallbackWrapper, NULL);
}

mal_device_config goConfigInitPlayback(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, NULL, goSendCallback);
}

mal_device_config goConfigInitDefaultCapture() {
    return mal_device_config_init_default_capture(goRecvCallbackWrapper);
}

mal_device_config goConfigInitDefaultPlayback() {
    return mal_device_config_init_default_playback(goSendCallback);
}

mal_context_config goContextConfigInit() {
    return mal_context_config_init(goLogCallbackWrapper);
}
