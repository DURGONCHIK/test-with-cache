package hashutil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func MD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func SHA256Hash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
