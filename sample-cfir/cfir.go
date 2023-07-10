package cfir

const (
	shouldBeUsed = iota
	shouldBeUnused
	shouldBeQuiet
)

func Name() string {
	return "cfir"
}

