package pgrange

import "database/sql"

type IntRange struct {
	Lower  sql.NullInt64
	Upper  sql.NullInt64
	Bounds Bounds
	Valid  bool
}

func (i *IntRange) Scan(value interface{}) error {
	if value == nil {
		i.Valid = false
	}

	lower, upper, bounds := parse(value.([]byte))

	if lower.Valid {
		err := i.Lower.Scan(lower.Bytes)
		if err != nil {
			return err
		}
	} else {
		i.Lower = sql.NullInt64{Valid: false}
	}

	if upper.Valid {
		err := i.Upper.Scan(upper.Bytes)
		if err != nil {
			return err
		}
	} else {
		i.Upper = sql.NullInt64{Valid: false}
	}

	i.Bounds = bounds

	return nil
}
