package kvs

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewEntry(t *testing.T) {
	_, err1 := NewEntry("", Value{Link: "v"})
	if err1 == nil {
		t.Error("no error for empty key")
	}

	_, err2 := NewEntry("k", Value{Link: ""})
	if err2 == nil {
		t.Error("no error for empty value")
	}
}

func TestNewEntryWithKeyWhichHasSlash(t *testing.T) {
	_, err := NewEntry("a/b", Value{Link: "foo"})
	if err == nil {
		t.Error("no error for key with slash")
	}
}

func TestFilterByPrefix(t *testing.T) {
	entry1 := Entry{Key: "key1", Value: Value{Link: "v9"}}
	entry2 := Entry{Key: "kay2", Value: Value{Link: "v8"}}
	entry3 := Entry{Key: "key3", Value: Value{Link: "v7"}}
	entries := Entries([]Entry{entry1, entry2, entry3})

	result := entries.FilterByPrefix("ke")
	if !reflect.DeepEqual(result, Entries([]Entry{entry1, entry3})) {
		t.Errorf("unexpected result %s", result)
	}
}

func TestSortable(t *testing.T) {
	entry1 := Entry{Key: "k1", Value: Value{Link: "v9"}}
	entry2 := Entry{Key: "k2", Value: Value{Link: "v8"}}
	entry3 := Entry{Key: "k3", Value: Value{Link: "v7"}}
	entries := Entries([]Entry{entry3, entry1, entry2})
	sort.Sort(entries)

	if !reflect.DeepEqual(entries, Entries([]Entry{entry1, entry2, entry3})) {
		t.Errorf("sort failed %s", entries)
	}
}
