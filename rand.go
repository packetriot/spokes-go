package spokes

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	int64Max = big.NewInt(0x7FFFFFFFFFFFFFFF)
)

func randomInt() int64 {
	value, _ := rand.Int(rand.Reader, int64Max)
	return value.Int64()
}

func randomNum(max int64) (n int64) {
	n = (randomInt() % max)
	return n
}

func randomString(n int64) string {
	var i int64

	buffer := bytes.Buffer{}
	for i = 0; i < ((n / 8) + 1); i++ {
		num := randomInt()
		chunk := make([]byte, 8)

		for j := 0; j < 8; j++ {
			chunk[j] = byte(0x000000000000FFFF & num)
			num >>= 8
		}

		buffer.Write(chunk)
	}

	data := buffer.Bytes()
	s := fmt.Sprintf("%x", (data[0:]))
	return s[:n]
}

const (
	minNonceLength int64 = 32
	maxNonceLength int64 = 512
)

func nonce() (nonce string) {
	n := randomNum(maxNonceLength-minNonceLength) + minNonceLength
	return string(n)
}
