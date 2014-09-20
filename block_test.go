package gost

import (
	"bytes"
	"testing"
)

type testVector struct {
	plaintext  []uint8
	ciphertext []uint8
	key        []uint8
}

var testVectors = []testVector{
	{
		plaintext:  []uint8{0x4e, 0x6f, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74},
		ciphertext: []uint8{0x28, 0x16, 0x30, 0xd0, 0xd5, 0x77, 0x00, 0x30},
		key: []uint8{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	},
	{
		plaintext:  []uint8{0x0d, 0xf8, 0x28, 0x02, 0xb7, 0x41, 0xa2, 0x92},
		ciphertext: []uint8{0x07, 0xf9, 0x02, 0x7d, 0xf7, 0xf7, 0xdf, 0x89},
		key: []uint8{0xbe, 0x5e, 0xc2, 0x00, 0x6c, 0xff, 0x9d, 0xcf,
			0x52, 0x35, 0x49, 0x59, 0xf1, 0xff, 0x0c, 0xbf,
			0xe9, 0x50, 0x61, 0xb5, 0xa6, 0x48, 0xc1, 0x03,
			0x87, 0x06, 0x9c, 0x25, 0x99, 0x7c, 0x06, 0x72},
	},
}

func TestVectors(t *testing.T) {
	for i, v := range testVectors {
		t.Logf("testVectors[%d] with values:", i)
		t.Logf("\t plaintext: %+v", v.plaintext)
		t.Logf("\tciphertext: %+v", v.ciphertext)
		t.Logf("\t       key: %+v", v.key)

		cipher, err := NewBlockCipher(v.key, TestSbox)
		if err != nil {
			t.Errorf("testVectors[%d]: NewBlockCipher() failed: %+v", i, err)
		} else {
			t.Logf("testVectors[%d]: NewBlockCipher() returned successfully", i)
		}

		encryptResult := make([]uint8, BlockSize)
		cipher.Encrypt(encryptResult, v.plaintext)

		if !bytes.Equal(v.ciphertext, encryptResult) {
			t.Errorf("testVectors[%d]: output of Encrypt() did not match expected ciphertext", i)
			t.Errorf("\tEncrypt() output: %+v", encryptResult)
		} else {
			t.Logf("testVectors[%d]: output of Encrypt() matched expected ciphertext", i)
		}

		decryptResult := make([]uint8, BlockSize)
		cipher.Decrypt(decryptResult, encryptResult)

		if !bytes.Equal(v.plaintext, decryptResult) {
			t.Errorf("testVectors[%d]: output of Decrypt() did not match expected plaintext", i)
			t.Errorf("\tDecrypt() output: %+v", decryptResult)
		} else {
			t.Logf("testVectors[%d]: output of Decrypt() matched expected plaintext", i)
		}
	}
}
