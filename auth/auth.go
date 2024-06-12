package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argon2Params struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

var argonOpts = &argon2Params{
	time:    1,
	memory:  64 * 1024,
	threads: 4,
	keyLen:  32,
}

func GeneratePasswordHash(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonOpts.time,
		argonOpts.memory,
		argonOpts.threads,
		argonOpts.keyLen,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	finalHash := fmt.Sprintf("%s:%s", b64Salt, b64Hash)
	return finalHash, nil
}

func VerifyPasswordHash(password, hash string) (bool, error) {
	parts := strings.Split(hash, ":")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	hashToVerify := argon2.IDKey(
		[]byte(password),
		salt,
		argonOpts.time,
		argonOpts.memory,
		argonOpts.threads,
		argonOpts.keyLen,
	)

	if subtle.ConstantTimeCompare(decodedHash, hashToVerify) == 1 {
		return true, nil
	}

	return false, nil
}
