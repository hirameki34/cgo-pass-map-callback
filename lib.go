package main

/*
#include "stdlib.h"

#define _CALLBACK_PARAMS int size, char **keys, char **values, void *fp

typedef void (*_callback_t) (_CALLBACK_PARAMS);
void __makeCallback(_callback_t cb, _CALLBACK_PARAMS);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export __registerCallback
func __registerCallback(callback C._callback_t, fp *C.void) {
	fn := func(fields map[string]interface{}) {
		size := len(fields)
		ks, vs := make([]*C.char, 0, size), make([]*C.char, 0, size)
		for k, v := range fields {
			ks = append(ks, C.CString(fmt.Sprint(k)))
			vs = append(vs, C.CString(fmt.Sprint(v)))
		}
		defer func() {
			for _, k := range ks {
				C.free(unsafe.Pointer(k))
			}
			for _, v := range vs {
				C.free(unsafe.Pointer(v))
			}
		}()

		pk, pv := (**C.char)(unsafe.Pointer(&ks[0])), (**C.char)(unsafe.Pointer(&vs[0]))
		C.__makeCallback(callback, C.int(size), pk, pv, unsafe.Pointer(fp))
	}
	register(fn)
}

//export testCallback
func testCallback() {
	makeCallback(map[string]interface{}{
		"value": 100,
		"foo":   "bar",
	})
}

func main() {}
