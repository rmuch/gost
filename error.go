package gost

import "strconv"

type KeySizeError int

func (k KeySizeError) Error() string {
	return "gost: invalid key size: " + strconv.Itoa(int(k))
}

type SboxSizeError int

func (k SboxSizeError) Error() string {
	return "gost: invalid sbox size: " + strconv.Itoa(int(k))
}
