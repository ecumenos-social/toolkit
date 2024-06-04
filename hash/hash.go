package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"

	"golang.org/x/crypto/bcrypt"
)

func Hash(raw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(raw), 14)
	return string(bytes), err
}

func CompareHash(raw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}

func SHA256(raw string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(raw)))
}

func SHA512(raw string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(raw)))
}

func SHA1(raw string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(raw)))
}

func MD5(raw string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(raw)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func FNV(raw string) string {
	return fmt.Sprintf("%x", fnv.New128().Sum([]byte(raw)))
}

func Adler32(raw string) string {
	return fmt.Sprintf("%x", adler32.New().Sum([]byte(raw)))
}

func CRC32(raw string) string {
	crc32Table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum([]byte(raw), crc32Table)

	return fmt.Sprintf("%08x", checksum)
}

func CRC64(raw string) string {
	crc64Table := crc64.MakeTable(crc32.IEEE)
	checksum := crc64.Checksum([]byte(raw), crc64Table)

	return fmt.Sprintf("%08x", checksum)
}
