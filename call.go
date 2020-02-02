package miscellaneous

import (
	"fmt"
	"strings"
)

// Key represents key of mapping.
type Key string

// Value represents value of mapping.
type Value string

// KeyValue represents Key and Value mapping.
type KeyValue interface {
	Key() Key
	Value() Value
}

// ParseKeyValue parses line of string into key and value.
func ParseKeyValue(line string) (KeyValue, error) {
	count := strings.Count(line, "=")
	if count != 1 {
		return nil, fmt.Errorf("invalid input: %s", line)
	}
	items := strings.Split(line, "=")
	if len(items) != 2 {
		return nil, fmt.Errorf("invalid state expect 2 items, got %d from: %s", len(items), line)
	}
	return &pair{
		first:  strings.Trim(items[0], " \t"),
		second: strings.Trim(items[1], " \t"),
	}, nil
}

type pair struct {
	first  string
	second string
}

func (p *pair) Key() Key {
	return Key(p.first)
}

func (p *pair) Value() Value {
	return Value(p.second)
}
