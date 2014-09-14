package gost

// Endianness option
const littleEndian bool = false

// Convert a byte slice to a uint32 slice
func bytesToUint32s(b []byte) []uint32 {
	size := len(b) / 4
	dst := make([]uint32, size)

	for i := 0; i < size; i++ {
		j := i * 4

		if littleEndian {
			dst[i] = uint32(b[j+0])<<24 | uint32(b[j+1])<<16 | uint32(b[j+2])<<8 | uint32(b[j+3])
		} else {
			dst[i] = uint32(b[j+0]) | uint32(b[j+1])<<8 | uint32(b[j+2])<<16 | uint32(b[j+3])<<24
		}
	}

	return dst
}

// Convert a uint32 slice to a byte slice
func uint32sToBytes(w []uint32) []byte {
	size := len(w) * 4
	dst := make([]byte, size)

	for i := 0; i < len(w); i++ {
		j := i * 4

		if littleEndian {
			dst[j+0] = byte(w[i] >> 24)
			dst[j+1] = byte(w[i] >> 16)
			dst[j+2] = byte(w[i] >> 8)
			dst[j+3] = byte(w[i])
		} else {
			dst[j+0] = byte(w[i])
			dst[j+1] = byte(w[i] >> 8)
			dst[j+2] = byte(w[i] >> 16)
			dst[j+3] = byte(w[i] >> 24)
		}
	}

	return dst
}
