#ifndef H_MALGO
#define H_MALGO

#include "mini_al.h"

#ifdef __cplusplus
extern "C" {
#endif

size_t goContextSize(void);

extern void goRecvCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern mal_uint32 goSendCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern void goStopCallback(mal_device* pDevice);

extern void goLogCallback(mal_context* pContext, mal_device* pDevice, char* message);
void *goLogCallbackPointer(void);

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

#ifdef __cplusplus
}
#endif

#endif
