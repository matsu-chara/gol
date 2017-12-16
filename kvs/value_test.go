package kvs

import (
	"testing"
)

func TestIsRegistered(t *testing.T) {
	value1 := Value{Link: "v1"}
	value2 := Value{Link: "v1", RegisteredBy: "foo"}

	if !value1.IsRegisteredBy("") {
		t.Error("register check error")
	}
	if value1.IsRegisteredBy("bar") {
		t.Error("register check error")
	}
	if !value2.IsRegisteredBy("foo") {
		t.Error("register check error")
	}
	if value2.IsRegisteredBy("bar") {
		t.Error("register check error")
	}
}
