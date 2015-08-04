package pgrange

type Bounds int

const (
	IncInc Bounds = iota
	IncExc
	ExcInc
	ExcExc
)

func getBounds(b []byte) Bounds {
	switch string(b) {
	case "[]":
		return IncInc
	case "[)":
		return IncExc
	case "(]":
		return ExcInc
	case "()":
		return ExcExc
	}
	return IncExc
}

func (b Bounds) String() string {
	switch b {
	case IncInc:
		return "[]"
	case IncExc:
		return "[)"
	case ExcInc:
		return "(]"
	case ExcExc:
		return "()"
	}
	return ""
}

func (b Bounds) Lower() string {
	switch b {
	case IncInc:
		return "["
	case IncExc:
		return "["
	case ExcInc:
		return "("
	case ExcExc:
		return "("
	}
	return ""
}

func (b Bounds) Upper() string {
	switch b {
	case IncInc:
		return "]"
	case ExcInc:
		return "]"
	case IncExc:
		return ")"
	case ExcExc:
		return ")"
	}
	return ""
}
