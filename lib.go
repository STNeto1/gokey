package gokey

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	GenerateKeyError = fmt.Errorf("Error generating the key")
)

func ClearUUID(val uuid.UUID) string {
	return strings.ReplaceAll(val.String(), "-", "")
}

func GenerateKeyFromUUID(prefix string, val uuid.UUID) string {
	cleanPrefix := strings.TrimSuffix(prefix, "_")
	for {
		if strings.HasSuffix(cleanPrefix, "_") {
			cleanPrefix = strings.TrimSuffix(cleanPrefix, "_")
			continue
		}

		break
	}

	return cleanPrefix + "_" + ClearUUID(val)
}

func GenerateKey(prefix string, size int) (string, error) {
	cleanPrefix := strings.TrimSuffix(prefix, "_")
	for {
		if strings.HasSuffix(cleanPrefix, "_") {
			cleanPrefix = strings.TrimSuffix(cleanPrefix, "_")
			continue
		}

		break
	}

	nanoid, err := gonanoid.Generate(alphabet, size)
	if err != nil {
		// log.Println(err) // Enable as log flag later?
		return "", GenerateKeyError
	}

	return cleanPrefix + "_" + nanoid, nil
}

func MustGenerateKey(prefix string, size int) string {
	value, err := GenerateKey(prefix, size)
	if err != nil {
		panic(err)
	}

	return value
}
