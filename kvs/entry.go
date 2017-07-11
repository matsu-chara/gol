package kvs

import (
	"errors"
	"fmt"
	"strings"
)

// Entry is a record in KVS.
type Entry struct {
	Key   string
	Value string
}

const entrySeparator = ": "

// NewEntry returns new Entry.
func NewEntry(key string, value string) (*Entry, error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}

	if value == "" {
		return nil, errors.New("value is empty")
	}

	entry := Entry{
		Key:   key,
		Value: value,
	}
	return &entry, nil
}

func (e Entry) String() string {
	return e.ToPeco()
}

// ToPeco serialize entry for peco
func (e Entry) ToPeco() string {
	return fmt.Sprintf("%s%s%s", e.Key, entrySeparator, e.Value)
}

// EntryFromPeco deserialize entry from ToPeco format
func EntryFromPeco(str string) (*Entry, error) {
	keyval := strings.Split(str, entrySeparator)
	if len(keyval) != 2 {
		return nil, fmt.Errorf("fromPeco failed. %s", str)
	}
	entry := Entry{
		Key:   strings.Trim(keyval[0], "\n"),
		Value: strings.Trim(keyval[1], "\n"),
	}
	return &entry, nil
}

// Entries for sort and filter
type Entries []Entry

func (entries Entries) FilterByPrefix(prefix string) Entries {
	filtered := make([]Entry, 0)
	for _, entry := range entries {
		if strings.HasPrefix(entry.Key, prefix) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// Len for sort
func (es Entries) Len() int {
	return len(es)
}

// Less for sort
func (es Entries) Less(i, j int) bool {
	return es[i].Key < es[j].Key
}

// Swap for sort
func (es Entries) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}
