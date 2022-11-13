package hash

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
)

// Takes a byte slice b and converts its data to an FNV-1 sum as a string.
func BytesToFNV(b []byte) (s string, err error) {
	fnvHash := fnv.New64()
	_, err = fnvHash.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%v", fnvHash.Sum64())
	return s, nil
}

// Takes a byte slice b and converts its data to an md5 sum as a string.
func BytesToMd5(b []byte) string {
	md5Hash := md5.Sum(b)
	m := fmt.Sprintf("%x", md5Hash)

	return m
}
