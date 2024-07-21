package authsvc

import (
	"crypto/rand"
	"fmt"
	"log"
)

// 32 位随机 buffer 作为盐
func createRandSalt() []byte {
	output := make([]byte, 32)
	_, err := rand.Read(output)
	if err != nil {
		log.Panic(fmt.Errorf("create salt error: %w", err))
	}

	return output
}
