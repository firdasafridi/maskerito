package masker

const (
	bankStart  = 4
	bankEnd    = 8
	bankLength = 4
)

// Bank Account Number mask last 4 digits of ID number
//
// Example:
//   input: 7123123123
//   output: 7123123123
func BankAccountNumber(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	start := bankStart
	if len(input) <= bankStart {
		start = 1
	}

	return Replacer(input, strIteration(defaultMask, bankLength), start, bankEnd)
}
