package pgrange

import "testing"

func TestTsTzRange(t *testing.T) {
	input := `["2010-01-01 11:30:00+02","2010-01-01 15:00:00+02")`

	val := &TsTzRange{}
	val.Scan([]byte(input))

	t.Logf("Input %s : Min %v - Max %v - Bounds %s", input, val.Lower, val.Upper, val.Bounds)

	v, err := val.Value()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Value = %s", v)
}
