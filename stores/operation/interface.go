package operation

import "berty.tech/go-ipfs-log/entry"

// Operation Describe an CRDT operation
type Operation interface {
	// GetKey Gets a key if applicable (ie. key value stores)
	GetKey() *string

	// GetOperation Returns an operation name (ie. append, put, remove)
	GetOperation() string

	// GetValue Returns the operation payload
	GetValue() []byte

	// GetEntry Gets the underlying IPFS log Entry
	GetEntry() *entry.Entry

	// Marshal Serializes the operation
	Marshal() ([]byte, error)
}
