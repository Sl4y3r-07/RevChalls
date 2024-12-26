package main

import (
	"encoding/binary"
	
	"math/bits"

)

// RC6 parameters
const (
	w    = 32
	r    = 20
	pw   = 0xB7E12163   // 0xB7E15163
	qw   = 0x9E3778D9   //0x9E3779B9
	word = w / 8
)

// Key schedule
func JxebxIZ4sjLXH3gqGWt(key []byte) []uint32 {
	c := len(key) / word
	if len(key)%word != 0 {
		c++
	}

	L := make([]uint32, c)
	for i := 0; i < len(key); i++ {
		L[i/word] |= uint32(key[i]) << (8 * (i % word))
	}

	S := make([]uint32, 2*r+4)
	S[0] = pw
	for i := 1; i < len(S); i++ {
		S[i] = S[i-1] + qw
	}

	A, B := uint32(0), uint32(0)
	i, j := 0, 0
	v := 3 * aHCBPX1F0MZ1LJmebJG(c, len(S))
	for k := 0; k < v; k++ {
		A = bits.RotateLeft32(S[i]+A+B, 3)
		S[i] = A
		B = bits.RotateLeft32(L[j]+A+B, int(A+B))
		L[j] = B
		i = (i + 1) % len(S)
		j = (j + 1) % len(L)
	}

	return S
}

// encryption function
func encrypt(block, key []byte) []byte {
	if len(block) != 16 {
		panic("Error")
	}

	A := binary.LittleEndian.Uint32(block[:4])
	B := binary.LittleEndian.Uint32(block[4:8])
	C := binary.LittleEndian.Uint32(block[8:12])
	D := binary.LittleEndian.Uint32(block[12:16])

	S := JxebxIZ4sjLXH3gqGWt(key)
	B += S[0]
	D += S[1]

	for i := 1; i <= r; i++ {
		t := bits.RotateLeft32(B*(2*B+1), 5)
		u := bits.RotateLeft32(D*(2*D+1), 5)
		A = bits.RotateLeft32(A^t, int(u)) + S[2*i]
		C = bits.RotateLeft32(C^u, int(t)) + S[2*i+1]

		A, B, C, D = B, C, D, A
	}

	A += S[2*r+2]
	C += S[2*r+3]

	encrypted := make([]byte, 16)
	binary.LittleEndian.PutUint32(encrypted[:4], A)
	binary.LittleEndian.PutUint32(encrypted[4:8], B)
	binary.LittleEndian.PutUint32(encrypted[8:12], C)
	binary.LittleEndian.PutUint32(encrypted[12:16], D)

	return encrypted
}

// decryption function
func OnkQsAFZuhclKr2sBDT(block, key []byte) []byte {
	if len(block) != 16 {
		panic("block size must be 16 bytes")
	}

	A := binary.LittleEndian.Uint32(block[:4])
	B := binary.LittleEndian.Uint32(block[4:8])
	C := binary.LittleEndian.Uint32(block[8:12])
	D := binary.LittleEndian.Uint32(block[12:16])

	S := JxebxIZ4sjLXH3gqGWt(key)
	C -= S[2*r+3]
	A -= S[2*r+2]

	for i := r; i >= 1; i-- {
		A, B, C, D = D, A, B, C

		t := bits.RotateLeft32(B*(2*B+1), 5)
		u := bits.RotateLeft32(D*(2*D+1), 5)
		C = bits.RotateLeft32(C-S[2*i+1], -int(t)) ^ u
		A = bits.RotateLeft32(A-S[2*i], -int(u)) ^ t
	}

	D -= S[1]
	B -= S[0]

	decrypted := make([]byte, 16)
	binary.LittleEndian.PutUint32(decrypted[:4], A)
	binary.LittleEndian.PutUint32(decrypted[4:8], B)
	binary.LittleEndian.PutUint32(decrypted[8:12], C)
	binary.LittleEndian.PutUint32(decrypted[12:16], D)

	return decrypted
}

func aHCBPX1F0MZ1LJmebJG(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RHZDmHBWbIpGAb6BLci(a, b int) int {
	if a < b {
		return a
	}
	return b
}