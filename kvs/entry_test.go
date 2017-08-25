package kvs

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewEntry(t *testing.T) {
	_, err1 := NewEntry("", "v")
	if err1 == nil {
		t.Error("no error for empty key")
	}

	_, err2 := NewEntry("k", "")
	if err2 == nil {
		t.Error("no error for empty value")
	}
}

func TestNewEntryWithKeyWhichHasSlash(t *testing.T) {
	_, err := NewEntry("a/b", "foo")
	if err == nil {
		t.Error("no error for key with slash")
	}
}

func TestString(t *testing.T) {
	entry := Entry{Key: "k1", Value: "v1"}
	if entry.ToPeco() != entry.String() {
		t.Error("Serializable format is not same between ToPeco and String")
	}
}

func TestToPeco(t *testing.T) {
	entry := Entry{Key: "k1", Value: "v1"}
	result := entry.ToPeco()
	if result != "k1: v1" {
		t.Errorf("Serialized %s is not expected value", result)
	}
}

func TestEntryFromPeco(t *testing.T) {
	entry := Entry{Key: "k1", Value: "v1"}
	toPeco := entry.ToPeco()
	fromPeco, err := EntryFromPeco(toPeco)
	if err != nil {
		t.Errorf("deserialize failure %s", err)
	}

	if entry != *fromPeco {
		t.Errorf("Serialized and Deserialized is not same value %s", fromPeco)
	}

	fromPeco2, err := EntryFromPeco("test")
	if fromPeco2 != nil {
		t.Errorf("deserialize failed, but value returend %s", fromPeco2)
	}
	if err == nil {
		t.Error("expect deserialize failed, but no error")
	}
}

func TestFilterByPrefix(t *testing.T) {
	entry1 := Entry{Key: "key1", Value: "v9"}
	entry2 := Entry{Key: "kay2", Value: "v8"}
	entry3 := Entry{Key: "key3", Value: "v7"}
	entries := Entries([]Entry{entry1, entry2, entry3})

	result := entries.FilterByPrefix("ke")
	if !reflect.DeepEqual(result, Entries([]Entry{entry1, entry3})) {
		t.Errorf("unexpected result %s", result)
	}
}

func TestSortable(t *testing.T) {
	entry1 := Entry{Key: "k1", Value: "v9"}
	entry2 := Entry{Key: "k2", Value: "v8"}
	entry3 := Entry{Key: "k3", Value: "v7"}
	entries := Entries([]Entry{entry3, entry1, entry2})
	sort.Sort(entries)

	if !reflect.DeepEqual(entries, Entries([]Entry{entry1, entry2, entry3})) {
		t.Errorf("sort failed %s", entries)
	}
}
