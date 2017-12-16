package kvs

import (
	"fmt"
	"sort"
)

// Data in kvs
type Data struct {
	data map[string]Value
}

// Get key from data
// return entry and isExists
func (data *Data) Get(key string) (*Entry, bool) {
	value, ok := data.data[key]
	if !ok {
		return nil, false
	}

	entry := Entry{
		Key:   key,
		Value: value,
	}
	return &entry, true
}

// Put entry to data.
func (data *Data) Put(entry *Entry) error {
	oldValue, ok := data.Get(entry.Key)
	if ok {
		return fmt.Errorf("%s is already registered as %s by %s", entry.Key, oldValue.Value.Link, oldValue.Value.RegisteredBy)
	}

	data.data[entry.Key] = entry.Value
	return nil
}

// Remove key from data.
func (data *Data) Remove(key string, registeredBy string) error {
	if _, ok := data.data[key]; !ok {
		return nil
	}

	if !data.data[key].IsRegisteredBy(registeredBy) {
		return fmt.Errorf("registeredBy was not equal to \"%s\"", data.data[key].RegisteredBy)
	}

	delete(data.data, key)
	return nil
}

// ListKeys return all keys in data.
func (data *Data) ListKeys() []string {
	keys := make([]string, 0, len(data.data))
	for k := range data.data {
		keys = append(keys, k)
	}
	return keys
}

// List returns all entries in data.
func (data *Data) List() []Entry {
	entries := make([]Entry, 0, len(data.data))
	for k, v := range data.data {
		entries = append(entries, Entry{k, v})
	}
	sort.Sort(Entries(entries))
	return entries
}

// Dump data.
func (data *Data) Dump() map[string]Value {
	return data.data
}
