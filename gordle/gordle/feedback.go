package gordle

type hint struct {
	status    hintStatus
	character rune
}

type hintStatus byte

const (
	Undefined = iota
	Absent
	WrongPosition
	CorrectPosition
)

func (h hint) String() string {

	if len(string(h.character)) == 0 || h.status == Undefined {
		panic("Hint status cannot be undefined")
	}

	switch h.status {
	case Absent:
		return "[" + string(h.character) + "]"
	case WrongPosition:
		return TTYYellow + "[" + string(h.character) + "]" + TTYReset
	case CorrectPosition:
		return TTYGreen + "[" + string(h.character) + "]" + TTYReset
	default:
		panic("Hint status cannot be undefined")
	}
}
