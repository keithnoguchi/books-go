package log

import (
	api "distributed/ch04/api/v1"
)

type CommitLog struct{}

func (l *CommitLog) Append(record *api.Record) (uint64, error) {
	return 0, nil
}

func (l *CommitLog) Read(offset uint64) (*api.Record, error) {
	return nil, nil
}
