package gost

import (
	"bytes"
	"testing"
)

/*
func TestECB(t *testing.T) {
	key := []uint8{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}

	// "4e6f77206973207468652074696d6520666f7220616c6c20",
	// "281630d0d5770030068c252d841e84149ccc1912052dbc02",

	plaintext := []uint8{0x4e, 0x6f, 0x77, 0x20, 0x69, 0x73, 0x20, 0x74}
	ciphertext := []uint8{0x28, 0x16, 0x30, 0xd0, 0xd5, 0x77, 0x00, 0x30}

	cipher, err := NewBlockCipher(key, TestSbox)
	if err != nil {
		t.Fatalf("NewBlockCipher failed: %+v", err)
	}

	result := make([]uint8, BlockSize)
	cipher.Encrypt(result, plaintext)

	if !bytes.Equal(result, ciphertext) {
		t.Fail()
	}

	t.Log("EXPECTED:")
	t.Logf("%+v", ciphertext)

	t.Log("ACTUAL:")
	t.Logf("%+v", result)

	t.Log("PLAINTEXT:")
	t.Logf("%+v", plaintext)

	t.Log("DECODED:")
	cipher.Decrypt(result, result)
	t.Logf("%+v", result)

	if !bytes.Equal(result, plaintext) {
		t.Fail()
	}
}

func TestECB2(t *testing.T) {
	key := []uint8{0xbe, 0x5e, 0xc2, 0x00, 0x6c, 0xff, 0x9d, 0xcf,
		0x52, 0x35, 0x49, 0x59, 0xf1, 0xff, 0x0c, 0xbf,
		0xe9, 0x50, 0x61, 0xb5, 0xa6, 0x48, 0xc1, 0x03,
		0x87, 0x06, 0x9c, 0x25, 0x99, 0x7c, 0x06, 0x72}

	// BE5EC2006CFF9DCF52354959F1FF0CBFE95061B5A648C10387069C25997C0672
	// 0DF82802B741A292    07F9027DF7F7DF89

	plaintext := []uint8{0x0d, 0xf8, 0x28, 0x02, 0xb7, 0x41, 0xa2, 0x92}
	ciphertext := []uint8{0x07, 0xf9, 0x02, 0x7d, 0xf7, 0xf7, 0xdf, 0x89}

	cipher, err := NewBlockCipher(key, TestSbox)
	if err != nil {
		t.Fatalf("NewBlockCipher failed: %+v", err)
	}

	result := make([]uint8, BlockSize)
	cipher.Encrypt(result, plaintext)

	if !bytes.Equal(result, ciphertext) {
		t.Fail()
	}

	t.Log("EXPECTED:")
	t.Logf("%+v", ciphertext)

	t.Log("ACTUAL:")
	t.Logf("%+v", result)

	t.Log("PLAINTEXT:")
	t.Logf("%+v", plaintext)

	t.Log("DECODED:")
	cipher.Decrypt(result, result)
	t.Logf("%+v", result)

	if !bytes.Equal(result, plaintext) {
		t.Fail()
	}
}
*/

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
		cipher, err := NewBlockCipher(v.key, TestSbox)
		if err != nil {
			t.Fatalf("NewBlockCipher failed: %+v", err)
		}

		encryptResult := make([]uint8, BlockSize)
		cipher.Encrypt(encryptResult, v.plaintext)

		if !bytes.Equal(v.ciphertext, encryptResult) {
			t.Fatalf("testVectors[%d]: output of Encrypt() did not match expected ciphertext", i)
		}

		decryptResult := make([]uint8, BlockSize)
		cipher.Decrypt(decryptResult, encryptResult)

		if !bytes.Equal(v.plaintext, decryptResult) {
			t.Fatalf("testVectors[%d]: output of Decrypt() did not match expected plaintext", i)
		}
	}
}