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
	count := strings.Count(line, " =")
	if count == 0 {
		return nil, fmt.Errorf("invalid input: %s", line)
	}
	index := strings.Index(line, " =")
	key := line[:index]
	value := strings.Replace(line[index:], " =", "", 1)
	return &pair{
		first:  strings.Trim(key, " \t"),
		second: strings.Trim(value, " \t"),
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
