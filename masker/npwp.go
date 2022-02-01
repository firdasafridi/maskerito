package masker

const (
	npwpStart  = 6
	npwpEnd    = 15
	npwpLength = 9
)

// NPWP mask last 9 digits of ID number
//
// Example:
//   input: 08.123.123.12-123.322
//   output: B12345****
func NPWP(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	start := npwpStart
	if len(input) <= npwpStart {
		start = 1
	}

	return Replacer(input, strIteration(defaultMask, npwpLength), start, npwpEnd)
}
