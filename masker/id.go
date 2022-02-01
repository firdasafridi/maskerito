package masker

const (
	idStart  = 6
	idEnd    = 10
	idLength = 4
)

// ID mask last 4 digits of ID number
//
// Example:
//   input: B12345CAZ
//   output: B12345****
func ID(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	start := idStart
	if len(input) <= idStart {
		start = 2
	}

	return Replacer(input, strIteration(defaultMask, idLength), start, len(input))
}
