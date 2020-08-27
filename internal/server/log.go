package server

import (
	"fmt"
	"sync"
)

// Log is the physical commit log which stores all our records and is append only
type Log struct {
	mu      sync.Mutex
	records []Record
}

// NewLog returns a pointer to a new Log struct
func NewLog() *Log {
	return &Log{}
}

// Append adds a new record of the end of the Log, returns offset
func (l *Log) Append(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	record.Offset = uint64(len(l.records))
	l.records = append(l.records, record)
	return record.Offset, nil

}

func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if offset >= uint64(len(l.records)) {
		return Record{}, ErrOffsetNotFound
	}

	return l.records[offset], nil
}

// Record struct stores some data
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// ErrOffsetNotFound indicates offset wasn't found
var ErrOffsetNotFound = fmt.Errorf("offset not found")
