package kvs

import (
	"time"
)

// Value in kvs
type Value struct {
	Link         string
	RegisteredBy string
	CreatedAt    time.Time
}

// IsRegisteredBy check this value is registeredBy args
func (v Value) IsRegisteredBy(registeredBy string) bool {
	return v.RegisteredBy == registeredBy
}
