package kvs

import (
	"fmt"
	"sort"
)

// Data in kvs
type Data struct {
	data map[string]string
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
		return fmt.Errorf("%s is already registered as %s", entry.Key, oldValue)
	}

	data.data[entry.Key] = entry.Value
	return nil
}

// Remove key from data.
func (data *Data) Remove(key string) {
	delete(data.data, key)
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
func (data *Data) Dump() map[string]string {
	return data.data
}
