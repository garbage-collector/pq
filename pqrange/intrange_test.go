package pgrange

import "testing"

func TestIntRange(t *testing.T) {
	input := `[42,44)`

	val := &IntRange{}
	val.Scan([]byte(input))

	t.Logf("Input %s : Min %v - Max %v - Bounds %s", input, val.Lower, val.Upper, val.Bounds)
}
