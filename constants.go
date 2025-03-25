package malgo

/*
#include "miniaudio.h"
*/
import "C"

const (
	simdAlignment = 64
)

const (
	True  = C.ma_bool32(1) // MA_TRUE
	False = C.ma_bool32(0) // MA_FALSE
)
