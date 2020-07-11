package easycrypt_test

import (
	"testing"

	"github.com/hsmtkk/easy_crypt_go/pkg/easycrypt"
	"github.com/hsmtkk/easy_crypt_go/test"
)

func TestCrypt(t *testing.T) {
	orig := []byte("0123456789")
	want := []byte{170, 166, 196, 104, 175, 121, 68, 44, 174, 51}
	encrypted := easycrypt.Crypt(orig)
	test.AssertEqualBytes(t, want, encrypted)
	decrypted := easycrypt.Crypt(encrypted)
	test.AssertEqualBytes(t, orig, decrypted)
}
