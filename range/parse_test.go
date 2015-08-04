package pgrange

import "testing"

func TestParse(t *testing.T) {
	//input :=
	inputs := []string{`["2014-01-01","2015-01-01")`, `[42,44)`, `[,12]`, `[12,)`}

	for _, input := range inputs {
		min, max, bounds := parse([]byte(input))

		t.Logf("Input %s : Min %s - Max %s - Bounds %s", input, min.Bytes, max.Bytes, bounds)
	}
}
