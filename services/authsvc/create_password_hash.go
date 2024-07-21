package authsvc

import (
	"bytes"
	"crypto/sha256"
)

func createPasswordHash(password string, salt []byte) []byte {
	// 组合 password 和 salt
	var buffer bytes.Buffer
	buffer.Write([]byte(password))
	buffer.Write(salt)

	// sha256 hash
	h := sha256.New()
	h.Write(buffer.Bytes())
	return h.Sum(nil)
}
