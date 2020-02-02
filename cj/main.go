package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	miscellaneous "github.com/mike-neck/github-actions-miscellaneous"
)

func main() {
	var entries Entries
	flag.Var(&entries, "prop", "key and value, separated by ' = '")
	flag.Parse()
	m := entries.ToMap()
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("error to make json", err)
	}
	fmt.Println(string(bytes))
}

// Entries is key value items
type Entries []miscellaneous.KeyValue

func (e *Entries) String() string {
	return fmt.Sprintf("%v", *e)
}

// Set inherit doc
func (e *Entries) Set(item string) error {
	kv, err := miscellaneous.ParseKeyValue(item)
	if err != nil {
		return err
	}
	*e = append(*e, kv)
	return nil
}

// ToMap makes map
func (e *Entries) ToMap() map[miscellaneous.Key]miscellaneous.Value {
	size := len(*e)
	m := make(map[miscellaneous.Key]miscellaneous.Value, size)
	for _, kv := range *e {
		m[kv.Key()] = kv.Value()
	}
	return m
}
