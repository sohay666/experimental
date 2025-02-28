package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Hash(data string) string {
	// Create an MD5 hash object
	hash := md5.New()

	// Write the data to the hash object
	io.WriteString(hash, data)

	// Get the MD5 hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	return fmt.Sprintf("%x", hashBytes)
}
