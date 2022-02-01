package masker

const (
	mobileStart  = 4
	mobileEnd    = 8
	mobileLength = 4
)

// Mobile mask will mask 4 digits from the middle digit
//
// Example:
//   input: 0987654321
//   output: 0987***321
func Mobile(input string) (output string) {
	if len(input) == 0 {
		return ""
	}

	return Replacer(input, strIteration(defaultMask, mobileLength), mobileStart, mobileEnd)
}
