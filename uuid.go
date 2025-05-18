package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

const (
	uuidLength   = 16
	randPoolSize = 4096
)

var (
	randPoolEnabled bool
	randPool         []byte
	randPoolLock     sync.Mutex
)

func EnableRandPool() {
	randPoolEnabled = true
}

func DisableRandPool() {
	randPoolEnabled = false
}

func fillRandPool() error {
	buf := make([]byte, randPoolSize)
	_, err := rand.Read(buf)
	if err != nil {
		return err
	}
	randPool = buf
	return nil
}

func Generate() (string, error) {
	bytes := make([]byte, uuidLength)

	if randPoolEnabled {
		randPoolLock.Lock()
		defer randPoolLock.Unlock()

		if len(randPool) < uuidLength {
			if err := fillRandPool(); err != nil {
				return "", err
			}
		}
		copy(bytes, randPool[:uuidLength])
		randPool = randPool[uuidLength:]
	} else {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}
	}

	bytes[6] = (bytes[6] & 0x0F) | 0x40
	bytes[8] = (bytes[8] & 0x3F) | 0x80

	return formatUUID(bytes), nil
}

func formatUUID(bytes []byte) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(bytes[0:4]),
		hex.EncodeToString(bytes[4:6]),
		hex.EncodeToString(bytes[6:8]),
		hex.EncodeToString(bytes[8:10]),
		hex.EncodeToString(bytes[10:16]),
	)
}

func Validate(uuid string) error {
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		return errors.New("invalid UUID format")
	}
	for _, part := range parts {
		if _, err := hex.DecodeString(part); err != nil {
			return errors.New("invalid UUID characters")
		}
	}
	return nil
}

func MustGenerate() string {
	uuid, err := Generate()
	if err != nil {
		panic("failed to generate UUID: " + err.Error())
	}
	return uuid
}

func ToBinary(uuid string) ([]byte, error) {
	uuid = strings.ReplaceAll(uuid, "-", "")
	if len(uuid) != 32 {
		return nil, errors.New("invalid UUID length")
	}
	return hex.DecodeString(uuid)
}

func GetParts(uuid string) ([]string, error) {
	if err := Validate(uuid); err != nil {
		return nil, err
	}
	return strings.Split(uuid, "-"), nil
}
