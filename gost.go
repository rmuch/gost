package gost

// Common operations shared between the GOST block cipher and hash function

// Expand s-box
func sboxExpansion(s [][]byte) [][]byte {
	// allocate buffer
	k := make([][]byte, 4)
	for i := 0; i < len(k); i++ {
		k[i] = make([]byte, 256)
	}

	// compute expansion
	for i := 0; i < 256; i++ {
		k[0][i] = (s[7][i>>4] << 4) | s[6][i&15]
		k[1][i] = (s[5][i>>4] << 4) | s[4][i&15]
		k[2][i] = (s[3][i>>4] << 4) | s[2][i&15]
		k[3][i] = (s[1][i>>4] << 4) | s[0][i&15]
	}

	return k
}

// Compute GOST cycle
func cycle(x uint32, kbox [][]byte) uint32 {
	x = uint32(kbox[0][(x>>24)&255])<<24 | uint32(kbox[1][(x>>16)&255])<<16 |
		uint32(kbox[2][(x>>8)&255])<<8 | uint32(kbox[3][x&255])

	return (x << 11) | (x >> (32 - 11))
}
