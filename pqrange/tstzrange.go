package pgrange

import (
	"database/sql/driver"
	"fmt"

	"github.com/lib/pq"
	"github.com/lib/pq/oid"
)

type TsTzRange struct {
	Lower  pq.NullTime
	Upper  pq.NullTime
	Bounds Bounds
	Valid  bool
}

func (i *TsTzRange) Scan(value interface{}) error {
	if value == nil {
		i.Valid = false
	}

	lower, upper, bounds := parse(value.([]byte))

	if lower.Valid {
		t, err := parseTime(oid.T_timestamptz, lower.Bytes)
		if err != nil {
			return err
		}
		i.Lower = pq.NullTime{Valid: true, Time: *t}
	} else {
		i.Lower = pq.NullTime{Valid: false}
	}

	if upper.Valid {
		t, err := parseTime(oid.T_timestamptz, upper.Bytes)
		if err != nil {
			return err
		}
		i.Upper = pq.NullTime{Valid: true, Time: *t}
	} else {
		i.Upper = pq.NullTime{Valid: false}
	}

	i.Bounds = bounds

	return nil
}

func (i TsTzRange) Value() (driver.Value, error) {
	str := i.Bounds.Lower()

	if i.Lower.Valid {
		str += fmt.Sprintf("\"%s\"", i.Lower.Format())
	}

	str += ","

	if i.Upper.Valid {
		str += fmt.Sprintf("\"%s\"", i.Upper.Format())
	}

	str += i.Bounds.Upper()

	return str, nil
}
