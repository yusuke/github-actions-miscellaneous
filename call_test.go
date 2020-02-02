package miscellaneous

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseKeyValue(t *testing.T) {
	kv, err := ParseKeyValue("foo = bar")
	assert.Nil(t, err)
	assert.Equal(t, Key("foo"), kv.Key())
	assert.Equal(t, Value("bar"), kv.Value())
}

func TestParseKeyValue2(t *testing.T) {
	kv, err := ParseKeyValue("GITHUB = fun")
	assert.Nil(t, err)
	assert.Equal(t, Key("GITHUB"), kv.Key())
	assert.Equal(t, Value("fun"), kv.Value())
}
