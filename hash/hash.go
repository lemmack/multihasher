package hash

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
)

// Takes any io.Reader (such as a multipart.File), reads its data and converts the data to an FNV-1 sum in string format.
func BytesToFNV(b []byte) (s string, err error) {

	h := fnv.New64()
	_, err = h.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%v", h.Sum64())
	return s, nil
}

// Takes any io.Reader (such as a multipart.File), reads its data and converts the data to an md5 sum in string format.
func BytesToMd5(b []byte) string {
	sum := md5.Sum(b)
	m := fmt.Sprintf("%x", sum)

	return m
}
