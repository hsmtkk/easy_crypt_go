package easycrypt

import (
	"golang.org/x/crypto/salsa20"
)

func Crypt(in []byte) []byte {
	out := make([]byte, len(in))
	nonce := make([]byte, 8)
	var key [32]byte
	salsa20.XORKeyStream(out, in, nonce, &key)
	return out
}
