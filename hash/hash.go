package hash

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
)

// Takes a byte slice and converts it to an FNV-1 sum as a string.
func BytesToFNV(b []byte) (s string, err error) {
	h := fnv.New64()
	_, err = h.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%v", h.Sum64())
	return s, nil
}

// Takes a byte slice and converts it to an md5 sum as a string.
func BytesToMd5(b []byte) string {
	sum := md5.Sum(b)
	m := fmt.Sprintf("%x", sum)

	return m
}
