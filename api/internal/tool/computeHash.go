package tool

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// ComputeFileHash computes the SHA-256 hash of the given file.
func ComputeFileHash(file io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// VerifyFileIntegrity verifies if the given file's hash matches the stored hash.
func VerifyFileIntegrity(file io.Reader, storedHash string) (bool, error) {
	// Compute the current file hash
	currentHash, err := ComputeFileHash(file)
	if err != nil {
		return false, err
	}

	// Compare the current hash with the stored hash
	if currentHash == storedHash {
		return true, nil
	}
	return false, nil
}
