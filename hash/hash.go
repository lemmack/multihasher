package hash

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash/fnv"

	"golang.org/x/crypto/ripemd160"
)

type Hashes struct {
	Fnv    string `json:"fnv"`
	Md5    string `json:"md5"`
	Ripemd string `json:"ripemd160"`
}

// Takes a byte slice b and returns a json object (as a byte slice) containing multiple hashes of the data in b
func MakeHashJson(b []byte) ([]byte, error) {
	rawFnv, err := bytesToFNV(b)

	if err != nil {
		return nil, err
	}

	rawMd5 := bytesToMD5(b)
	rawRipemd, err := bytesToRIPEMD(b)

	if err != nil {
		return nil, err
	}

	hashStruct := Hashes{Fnv: rawFnv, Md5: rawMd5, Ripemd: rawRipemd}
	jsonBytes, err := json.Marshal(hashStruct)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// Takes a byte slice and converts its data to an FNV-1 sum as a string.
func bytesToFNV(b []byte) (s string, err error) {
	fnvHash := fnv.New64()
	_, err = fnvHash.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%v", fnvHash.Sum64())
	return s, nil
}

// Takes a byte slice and converts its data to an MD5 sum as a string.
func bytesToMD5(b []byte) string {
	md5Hash := md5.Sum(b)

	s := fmt.Sprintf("%x", md5Hash)
	return s
}

// Takes a byte slice and converts its data to a RIPEMD-160 sum as a string.
func bytesToRIPEMD(b []byte) (s string, err error) {
	rmdHash := ripemd160.New()
	_, err = rmdHash.Write(b)

	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%x", rmdHash.Sum(nil))
	return s, nil
}
