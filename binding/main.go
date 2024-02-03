package main
//#include <stdint.h>
//#include <stdlib.h>
import "C"
import(
	"fmt"
	"encoder"
)

//export Encode
func Encode(message *C.char) string {
	fmt.Printf("hello = %v\n", C.GoString(message))
	return "Hey there"
	// return encoder.Encode(C.GoString(message))
}

//export Decode
func Decode(encoded []byte) string {
	return encoder.Decode(encoded)
}

func main() {}