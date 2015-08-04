package pgrange

import (
	"errors"
	"time"

	"github.com/lib/pq/oid"
)

func parseTime(typ oid.Oid, s []byte) (*time.Time, error) {
	str := string(s)

	var f string
	switch typ {
	case oid.T_timestamptz:
		f = "2006-01-02 15:04:05-07"
	case oid.T_timetz:
		f = "15:04:05"
	default:
		return nil, errors.New("invalid datetime type")
	}

	if str[len(str)-3] == ':' {
		f += ":00"
	}

	t, err := time.Parse(f, str)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
