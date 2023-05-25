package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	private_key := GenratePrivateKey()

	assert.Nil(t, private_key.GenratePublicKey().key)
}
