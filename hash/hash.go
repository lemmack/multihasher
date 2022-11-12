package hash

import (
	"fmt"
	"hash/fnv"
	"io"
)

// Takes any io.Reader (such as a multipart.File), reads its data and converts the data to an FNV-1 sum in string format.
func ReaderToFNV(f io.Reader) (s string) {
	h := fnv.New64()
	io.Copy(h, f)
	s = fmt.Sprintf("%v", h.Sum64())
	return s
}
