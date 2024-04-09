package gokey_test

import (
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/stneto1/gokey"
	"github.com/stretchr/testify/assert"
)

func TestClearString(t *testing.T) {
	val, err := uuid.FromString("c6b10dd3-1dcf-416c-8ed8-ae561807fcaf")
	assert.NoError(t, err, "Generating UUID from String got an error")

	result := gokey.ClearUUID(val)
	assert.Equal(t, "c6b10dd31dcf416c8ed8ae561807fcaf", result, "Clearing UUID got an error")
}

func TestGenerateKeyFromUuid(t *testing.T) {
	val, err := uuid.FromString("c6b10dd3-1dcf-416c-8ed8-ae561807fcaf")
	assert.NoError(t, err, "Generating UUID from String got an error")

	result := gokey.GenerateKeyFromUUID("go", val)
	assert.Equal(t, "go_c6b10dd31dcf416c8ed8ae561807fcaf", result, "Generating key got an error")
}

func TestGenerateKeyFromUuidWithTrailingUnderscore(t *testing.T) {
	val, err := uuid.FromString("c6b10dd3-1dcf-416c-8ed8-ae561807fcaf")
	assert.NoError(t, err, "Generating UUID from String got an error")

	result := gokey.GenerateKeyFromUUID("go_", val)
	assert.Equal(t, "go_c6b10dd31dcf416c8ed8ae561807fcaf", result, "Generating key got an error")
}

func TestGenerateKeyFromUuidWithMultipleTrailingUnderscore(t *testing.T) {
	val, err := uuid.FromString("c6b10dd3-1dcf-416c-8ed8-ae561807fcaf")
	assert.NoError(t, err, "Generating UUID from String got an error")

	result := gokey.GenerateKeyFromUUID("go___", val)
	assert.Equal(t, "go_c6b10dd31dcf416c8ed8ae561807fcaf", result, "Generating key got an error")
}
