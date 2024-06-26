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
	GenerateKeyError      = fmt.Errorf("Error generating the key")
	InvalidKeyFormatError = fmt.Errorf("Given key was not formatted correctly")
	InvalidUUIDError      = fmt.Errorf("Given key uuid was not formatted correctly")
)

// ClearUUID removes all the - from the UUID
//  Example: 550e8400-e29b-41d4-a716-446655440000 => 550e8400e29b41d4a716446655440000
func ClearUUID(val uuid.UUID) string {
	return strings.ReplaceAll(val.String(), "-", "")
}

// GenerateKeyFromUUID generates a key from a UUID
//  Example: GenerateKeyFromUUID("prefix", uuid.NewV4()) => prefix_550e8400e29b41d4a716446655440000
// It will remove all the - from the UUID
// It will also remove any trailing _ from the prefix
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

// GetUUIDFromKey extracts the UUID from the key
//  Example: GetUUIDFromKey("prefix_550e8400e29b41d4a716446655440000") => 550e8400-e29b-41d4-a716-446655440000
// It will return an error if the key is not formatted correctly
//  It will return an error if the UUID is not formatted correctly
func GetUUIDFromKey(key string) (uuid.UUID, error) {
	tokens := strings.Split(key, "_")
	// 32 => UUID without any -
	if len(tokens) != 2 || len(tokens[1]) != 32 {
		return uuid.UUID{}, InvalidKeyFormatError
	}

	dirtyUuid := tokens[1]
	builder := strings.Builder{}
	builder.Grow(36)

	builder.WriteString(dirtyUuid[:8])
	builder.WriteString("-")
	builder.WriteString(dirtyUuid[8:12])
	builder.WriteString("-")
	builder.WriteString(dirtyUuid[12:16])
	builder.WriteString("-")
	builder.WriteString(dirtyUuid[16:20])
	builder.WriteString("-")
	builder.WriteString(dirtyUuid[20:])

	result, err := uuid.FromString(builder.String())
	if err != nil {
		return uuid.UUID{}, InvalidUUIDError
	}

	return result, nil
}

// GenerateKey generates a key with a prefix and a size
//  Example: GenerateKey("prefix", 10) => prefix_550e8400e2
// It will remove any trailing _ from the prefix
// It will return an error that if the key could not be generated correctly
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

// MustGenerateKey generates a key with a prefix and a size
//  Example: MustGenerateKey("prefix", 10) => prefix_550e8400e2
// It will remove any trailing _ from the prefix
// It will panic if the key could not be generated correctly
func MustGenerateKey(prefix string, size int) string {
	value, err := GenerateKey(prefix, size)
	if err != nil {
		panic(err)
	}

	return value
}
