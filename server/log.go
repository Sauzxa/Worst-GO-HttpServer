package server

import "sync"

type Recored struct {
	Offset uint64 `json:"offset"` // filed will be returned as json
	Value  []byte `json:"value"`  // same here we should pute it with capital to let it public
}

type Log struct {
	mu       sync.Mutex // to lock the loh when we append record
	Recoreds []Recored
}

func NewLog() *Log { // instance ml log struct to get access with pointer
	return &Log{}
}
func (l *Log) Append(record Recored) (uint64, error) { // append record to log
	defer l.mu.Unlock() // ki tkamal appending unlockiha
	l.mu.Lock()         // lock the recoreds when we append record
	// set the record offset
	record.Offset = uint64(len(l.Recoreds))
	//append

	l.Recoreds = append(l.Recoreds, record)

	return record.Offset, nil
}
func (l *Log) Read(offset uint64) (Recored, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Recoreds[offset], nil
}
