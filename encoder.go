package encoder

import (
	"sync"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding"
	"unicode/utf8"
)

type Decoder struct {
	leftover []byte
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

var (
	decoderArray     []*Decoder
	decoderArrayLock sync.Mutex
)

const initialDecoderArraySize = 10000

func init() {
	decoderArray = make([]*Decoder, initialDecoderArraySize)
}

func getDecoder(index int) *Decoder {
	decoderArrayLock.Lock()
	defer decoderArrayLock.Unlock()
	if index >= len(decoderArray) {
		newSize := len(decoderArray) * 2 // Double the size
		if index >= newSize {
			newSize = index + 1 // Ensure newSize accommodates index
		}
		newDecoderArray := make([]*Decoder, newSize)
		copy(newDecoderArray, decoderArray)
		decoderArray = newDecoderArray
	}

	if decoderArray[index] == nil {
		decoderArray[index] = NewDecoder()
	}
	return decoderArray[index]
}

func Encode(message string) []byte {
	return []byte(message)
}
func Decode(input[]byte, enc string, index int, stream int) string {

	decoderObj := getDecoder(index)

	combinedLen := len(decoderObj.leftover) + len(input)
	combinedBytes := make([]byte, combinedLen)

	copy(combinedBytes, decoderObj.leftover)
    copy(combinedBytes[len(decoderObj.leftover):], input)

	// Handle non-standard encodings.
	var customDecoder *encoding.Decoder
	switch(enc) {
		case "windows-1250":
			customDecoder = charmap.Windows1250.NewDecoder()
		case "windows-1251":
			customDecoder = charmap.Windows1251.NewDecoder()
		case "windows-1252":
			customDecoder = charmap.Windows1252.NewDecoder()
		case "windows-1253":
			customDecoder = charmap.Windows1253.NewDecoder()
		case "windows-1254":
			customDecoder = charmap.Windows1254.NewDecoder()
		case "windows-1255":
			customDecoder = charmap.Windows1255.NewDecoder()
		case "windows-1256":
			customDecoder = charmap.Windows1256.NewDecoder()
		case "windows-1257":
			customDecoder = charmap.Windows1257.NewDecoder()
		case "windows-1258":
			customDecoder = charmap.Windows1258.NewDecoder()
		case "iso-8859-1": 
			customDecoder = charmap.ISO8859_1.NewDecoder()
		case "iso-8859-2":
			customDecoder = charmap.ISO8859_2.NewDecoder()
		case "iso-8859-3":
			customDecoder = charmap.ISO8859_3.NewDecoder()
		case "iso-8859-4":
			customDecoder = charmap.ISO8859_4.NewDecoder()
		case "iso-8859-5":
			customDecoder = charmap.ISO8859_5.NewDecoder()
		case "iso-8859-6":
			customDecoder = charmap.ISO8859_6.NewDecoder()
		case "iso-8859-7":
			customDecoder = charmap.ISO8859_7.NewDecoder()
		case "iso-8859-8":
			customDecoder = charmap.ISO8859_8.NewDecoder()
		case "iso-8859-9":
			customDecoder = charmap.ISO8859_9.NewDecoder()
		case "iso-8859-10":
			customDecoder = charmap.ISO8859_10.NewDecoder()
		case "iso-8859-13":
			customDecoder = charmap.ISO8859_13.NewDecoder()
		case "iso-8859-14":
			customDecoder = charmap.ISO8859_14.NewDecoder()
		case "iso-8859-15":
			customDecoder = charmap.ISO8859_15.NewDecoder()
		case "iso-8859-16":
			customDecoder = charmap.ISO8859_16.NewDecoder()
	}

	if customDecoder != nil {
		combinedBytes, _ = customDecoder.Bytes([]byte(combinedBytes))
	}

	
	decoded := string(combinedBytes)
	if(stream == 1) {
		// Find the last complete rune
		lastRuneIndex := len(decoded)
		for lastRuneIndex > 0 && !utf8.ValidString(decoded[:lastRuneIndex]) {
			lastRuneIndex--
		}
		 // Update prevLeftover with the remaining incomplete bytes
		 decoderObj.leftover = combinedBytes[lastRuneIndex:]
		 return decoded[:lastRuneIndex]
	} else {
		return decoded
	}
}
