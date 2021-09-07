#include "miniaudio.h"
extern void goLogCallbackWrapper(ma_context* pContext, ma_device* pDevice, ma_uint32 logLevel, const char* message);
extern void goDataCallbackWrapper(ma_device* pDevice, void* pOutput, const void* pInput, ma_uint32 frameCount);
extern void goStopCallbackWrapper(ma_device* pDevice);
