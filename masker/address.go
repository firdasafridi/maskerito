package masker

const (
	addressStart  = 6
	addressLength = 8
)

// Address keep first 6 letters, mask the rest
//
// Example:
//   input: B12345CAZ
//   output: Jln Ja********
func Address(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	start := addressStart
	if len(input) <= addressStart {
		start = 2
	}

	return Replacer(input, strIteration(defaultMask, addressLength), start, len(input))
}
