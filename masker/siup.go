package masker

const (
	siupStart  = 8
	siupEnd    = 12
	siupLength = 4
)

// SIUP mask last 4 digits of ID number
//
// Example:
//   input: 510/20/ABC/1888
//   output: 510/20/A****888
func SIUP(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	start := siupStart
	if len(input) <= siupStart {
		start = 1
	}

	return Replacer(input, strIteration(defaultMask, siupLength), start, siupEnd)
}
