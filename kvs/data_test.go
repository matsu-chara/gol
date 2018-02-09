package kvs

import (
	"reflect"
	"sort"
	"testing"
)

func TestGet(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data: values,
	}

	result, isExists := data.Get("k1")
	if !isExists {
		t.Error("k1 was not found.")
	}

	if result.Key != "k1" {
		t.Errorf("%s is not expected key `k1`", result)
	}

	if result.Value.Link != "v1" {
		t.Errorf("%s is not expected value `v1`", result)
	}

	result2, isExists2 := data.Get("notFound")
	if isExists2 {
		t.Error("key notFound was found.")
	}

	if result2 != nil {
		t.Errorf("%s is not expected value nil", result2)
	}
}

func TestPut(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data:       values,
		permission: ReadAndWrite,
	}

	err := data.Put(&Entry{
		Key:   "k3",
		Value: Value{Link: "v3", RegisteredBy: "r3"},
	})
	if err != nil {
		t.Error("put was failed")
	}
	result, isExists := data.Get("k3")
	if !isExists {
		t.Error("k3 was not found.")
	}
	if result.Value.Link != "v3" {
		t.Error("k3 link was wrong")
	}
	if result.Value.RegisteredBy != "r3" {
		t.Error("k3 registeredBy was wrong")
	}
	err2 := data.Put(&Entry{
		Key:   "k1",
		Value: Value{Link: "v0"},
	})
	if err2 == nil {
		t.Error("k1 was already registered but not erro")
	}
}

func TestRemove(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data:       values,
		permission: ReadAndWrite,
	}

	data.Remove("k1", "")
	_, isExists := data.Get("k1")
	if isExists {
		t.Error("k1 was found.")
	}

	err := data.Put(&Entry{
		Key:   "k3",
		Value: Value{Link: "v3", RegisteredBy: "r3"},
	})
	if err != nil {
		t.Error("put was failed")
	}
	err = data.Remove("k3", "")
	if err == nil {
		t.Error("registeredBy was not equal but no error")
	}
	data.Remove("k3", "r3")
	_, isExists = data.Get("k3")
	if isExists {
		t.Error("k3 was found.")
	}
}

func TestListKeys(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data: values,
	}
	keys := data.ListKeys()
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"k1", "k2"}) {
		t.Errorf("keys are not equal %s", keys)
	}
}

func TestList(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data: values,
	}
	keys := data.List()
	if !reflect.DeepEqual(keys, []Entry{
		{Key: "k1", Value: Value{Link: "v1"}},
		{Key: "k2", Value: Value{Link: "v2"}},
	}) {
		t.Errorf("keys are not equal %s", keys)
	}
}

func TestDump(t *testing.T) {
	values := map[string]Value{"k1": {Link: "v1"}, "k2": {Link: "v2"}}
	data := Data{
		data: values,
	}
	dumped := data.Dump()
	if !reflect.DeepEqual(dumped, values) {
		t.Errorf("keys are not equal %s", dumped)
	}
}
