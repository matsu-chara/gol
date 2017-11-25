package kvs

import (
	"errors"
	"fmt"
	"strings"
)

// Entry is a record in KVS.
type Entry struct {
	Key   string
	Value Value
}

const entrySeparator = ": "

// NewEntry returns new Entry.
func NewEntry(key string, value Value) (*Entry, error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}
	if len(strings.Split(key, "/")) != 1 {
		return nil, fmt.Errorf("key contains '/'")
	}

	if value.Link == "" {
		return nil, errors.New("link is empty")
	}

	entry := Entry{
		Key:   key,
		Value: value,
	}
	return &entry, nil
}

func (e Entry) String() string {
	return fmt.Sprintf("%s%s%s", e.Key, entrySeparator, e.Value.Link)
}

// Entries for sort and filter
type Entries []Entry

// FilterByPrefix filters entries by entry prefix
func (es Entries) FilterByPrefix(prefix string) Entries {
	filtered := make([]Entry, 0)
	for _, entry := range es {
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
