package server

type Recored struct {
	Offset uint64 `json:"offset"` // filed will be returned as json
	Value  []byte `json:"value"`  // same here we should pute it with capital to let it public
}

type Log struct {
	Recoreds []Recored
}

func NewLog() *Log { // instance ml log struct to get access with pointer
	return &Log{}
}
func (l *Log) Append(record Recored) (uint64, error) { // append record to log
	// set the record offset
	record.Offset = uint64(len(l.Recoreds))
	l.Recoreds = append(l.Recoreds, record)

	return record.Offset, nil
}
