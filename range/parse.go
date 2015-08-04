package pgrange

// https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown
import "regexp"

var r *regexp.Regexp

type NullBytes struct {
	Valid bool
	Bytes []byte
}

func parse(input []byte) (lower NullBytes, upper NullBytes, bounds Bounds) {

	if r == nil {
		// Taken from psycopg2
		// https://github.com/psycopg/psycopg2/blob/master/lib/_range.py

		// First bound
		reg := `(\(|\[)`
		// First value
		reg += `("(([^"]|"")*)"|([^",]+))?`
		// Separator
		reg += `,`
		// Second value
		reg += `("(([^"]|"")*)"|([^"\)\]]+))?`
		// Last bound
		reg += `(\)|\])`

		r, _ = regexp.Compile(reg)
	}

	res := r.FindSubmatch(input)

	// Bounds are on 1 and 10
	bounds = getBounds([]byte{res[1][0], res[10][0]})

	// lower value is on 3 or 2
	lower = NullBytes{Valid: false}
	if len(res[3]) > 0 {
		lower = NullBytes{Valid: true, Bytes: res[3]}
	} else if len(res[2]) > 0 {
		lower = NullBytes{Valid: true, Bytes: res[2]}
	}

	// upper value is on 7 or 6
	upper = NullBytes{Valid: false}
	if len(res[7]) > 0 {
		upper = NullBytes{Valid: true, Bytes: res[7]}
	} else if len(res[6]) > 0 {
		upper = NullBytes{Valid: true, Bytes: res[6]}
	}

	return lower, upper, bounds
}
