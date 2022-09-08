package services

import (
	"fmt"
	"crypto/sha256"
)

func SHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)

}
