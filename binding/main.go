package main

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
typedef struct  { void* message; int size; char* error; } GoResponse;

GoResponse* Encode(char*);

__attribute__((weak))
int myFunction() {

	char arr[4] = "mama";
	void* z = (void *)(intptr_t)Encode(arr);
	uint8_t *boss = z;
	printf("%d", boss[1]);
	printf("\n");
	return 1;
}

__attribute__((weak))
void* getArray() {
	uint8_t* boss = (uint8_t*)malloc(4 * sizeof(uint8_t));
    boss[0] = 1;
    boss[1] = 2;
    boss[2] = 3;
    boss[3] = 4;

	return boss;
}
*/
import "C"
import (
	"encoder"
)

//export Encode
func Encode(message *C.char) *C.GoResponse {
	encoded := encoder.Encode(C.GoString(message))
	result := (*C.GoResponse)(C.malloc(C.size_t(C.sizeof_GoResponse)))
	result.error = nil
	result.message = C.CBytes(encoded)
	result.size = C.int(len(encoded))
	return result
}
//export Decode
func Decode(encoded []byte) string {
	return encoder.Decode(encoded)
}

func main() {
	C.myFunction();
	// gg := C.getArray()

	// // Convert the C array to a Go slice for easier manipulation
	// arraySize := 4
	// array := (*[1 << 30]C.uint8_t)(gg)[:arraySize:arraySize]

	// // Access the elements of the array
	// for i, value := range array {
	// 	fmt.Printf("Element %d: %v\n", i, value)
	// }
}
