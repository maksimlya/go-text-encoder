package main

//#include <stdint.h>
//#include <stdlib.h>
//typedef struct  { void* message; int size; char* error; } GoResponse;
import "C"
import (
	"encoder"
	"unsafe"
)

//export Encode
func Encode(message *C.char) *C.GoResponse {
	encoded := encoder.Encode(C.GoString(message))
    result := (*C.GoResponse)(C.malloc(C.size_t(C.sizeof_GoResponse)))
	result.error = nil
	result.message = C.CBytes(encoded)
	result.size = C.int(len(encoded))

	// Memory is freed on caller side
	return result
}

//export Decode
func Decode(ptr unsafe.Pointer, length C.int, enc *C.char, index C.int, stream C.int) *C.char {
    data := C.GoBytes(ptr, length)
    result := encoder.Decode(data, C.GoString(enc), int(index), int(stream))
	decodedCString := C.CString(result)
	// Memory is freed on caller side
	return decodedCString
}

func main() {}
