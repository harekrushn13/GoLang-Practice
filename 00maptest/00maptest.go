package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func main() {
	m := make(map[string]int)
	h := make(map[uint64]int)

	size := 10000000

	ms := make([]string, size)
	hs := make([]uint64, size)

	sequenceCounter := 1

	for i := 0; i < size; i++ {
		key, hash := generateKey(&sequenceCounter)

		if len(key) != 10 {
			panic("key is not 10 bytes long")
		}

		m[key] = len(m)
		h[hash] = len(m)

		ms[i] = key
		hs[i] = hash
	}

	start := time.Now()
	for i := 0; i < size; i++ {
		_ = m[ms[i]]
	}
	end1 := time.Since(start)

	start2 := time.Now()
	for i := 0; i < size; i++ {
		_ = h[hs[i]]
	}
	end2 := time.Since(start2)

	fmt.Println("string : ", end1)
	fmt.Println("uint64 : ", end2)

}

func sha256Hash10Bytes(s []byte) uint64 {

	hash := sha256.Sum256(s)

	return binary.BigEndian.Uint64(hash[:8])
}

func generateKey(sequenceCounter *int) (string, uint64) {
	var keyBuilder strings.Builder

	for keyBuilder.Len() < 10 {
		keyBuilder.WriteString(fmt.Sprintf("%d#", *sequenceCounter))
		*sequenceCounter++
	}

	key := keyBuilder.String()
	if len(key) > 10 {
		key = key[:10]
	}

	keyBytes := []byte(key)
	hash := sha256Hash10Bytes(keyBytes)

	return key, hash
}
