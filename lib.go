package gokey

import (
	"strings"

	"github.com/gofrs/uuid/v5"
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
