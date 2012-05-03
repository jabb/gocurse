package menus

// #define _Bool int
// #include <curses.h>
import "C"

import "unsafe"

type void unsafe.Pointer
type Chtype C.chtype

func boolToInt(b bool) C.int {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

func intToBool(b C.int) bool {
	if b == C.TRUE {
		return true
	}
	return false
}

func isOk(ok C.int) bool {
	if ok == C.OK {
		return true
	}
	return false
}
