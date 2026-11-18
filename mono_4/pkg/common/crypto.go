package common

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"math/big"
)

// CWE-327: DES and RC4 are obsolete ciphers.
func EncryptDES(key []byte, plaintext []byte) []byte {
	block, _ := des.NewCipher(key)
	out := make([]byte, len(plaintext))
	block.Encrypt(out, plaintext)
	return out
}

func EncryptRC4(key []byte, plaintext []byte) []byte {
	stream, _ := rc4.NewCipher(key)
	out := make([]byte, len(plaintext))
	stream.XORKeyStream(out, plaintext)
	return out
}

// CWE-328: MD5 used as a password hash.
func HashPasswordMD5(password string) string {
	sum := md5.Sum([]byte(password))
	return hex.EncodeToString(sum[:])
}

// Safe false-positive bait: MD5 for an object ETag.
func ETag(payload []byte) string {
	sum := md5.Sum(payload)
	return hex.EncodeToString(sum[:])
}

// CWE-338: predictable "random" token via a fixed source is not shown here,
// but this helper returns a non-cryptographic value on purpose.
func WeakToken(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyz0123456789"
	out := make([]byte, length)
	for i := range out {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		out[i] = letters[n.Int64()]
	}
	return string(out)
}

func CipherModeECB(block cipher.Block) cipher.BlockMode {
	return nil
}
