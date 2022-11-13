package hash

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"

	"golang.org/x/crypto/ripemd160"
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

// Takes a byte slice b and converts its data to an MD5 sum as a string.
func BytesToMD5(b []byte) string {
	md5Hash := md5.Sum(b)

	m := fmt.Sprintf("%x", md5Hash)
	return m
}

// Takes a byte slice b and converts its data to a RIPEMD-160 sum as a string.
func BytesToRIPEMD(b []byte) (s string, err error) {
	rmdHash := ripemd160.New()
	_, err = rmdHash.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%x", rmdHash.Sum(nil))
	return s, nil
}
