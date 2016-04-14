package util

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

const tableSeperator rune = '/'

// SanitizeKey ensures that the key does not have any trailing '/'s and starts with a '/'
func SanitizeKey(key string) string {
	sanitized := []rune(key)
	if sanitized[0] != tableSeperator {
		sanitized = append([]rune{tableSeperator}, sanitized...)
	}
	if sanitized[len(sanitized)-1] == tableSeperator {
		sanitized = sanitized[:len(sanitized)-1]
	}
	return string(sanitized)
}

// EncodeULeb128 encode's an unsigned int32 value to an unsigned LEB128 value. Returns the result in a byte slice
func EncodeULeb128(value uint32) []byte {
	remaining := value >> 7
	var buf = new(bytes.Buffer)
	for remaining != 0 {
		buf.WriteByte(byte(value&0x7f | 0x80))
		value = remaining
		remaining >>= 7
	}
	buf.WriteByte(byte(value & 0x7f))
	return buf.Bytes()
}

// ReadULeb128 reads and decodes an unsigned LEB128 value from a ByteReader to an unsigned int32 value. Returns the result as a uint32
func ReadULeb128(reader io.ByteReader) (uint32, uint) {
	var result uint32
	var ctr uint
	var cur byte = 0x80
	var err error
	for (cur&0x80 == 0x80) && ctr < 5 {
		cur, err = reader.ReadByte()
		if err != nil {
			panic(err)
		}
		result += uint32((cur & 0x7f)) << (ctr * 7)
		ctr++
	}
	return result, ctr
}

// BytesToFloat64 converts bytes to Float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
