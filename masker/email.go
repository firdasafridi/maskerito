package masker

import "strings"

const (
	emailStart  = 3
	emailEnd    = 7
	emailLength = 4
)

// Email keep domain and the first 3 letters
//
// Example:
//   input: sample.saja@gmail.com
//   output: ggw****@gmail.com
func Email(input string) (output string) {
	l := len([]rune(input))
	if l == 0 {
		return ""
	}

	listDomain := strings.Split(input, "@")

	start := emailStart
	if emailStart <= len(input) {
		start = 1
	}

	if len(listDomain) == 1 {
		return Replacer(input, strIteration(defaultMask, emailLength), start, emailEnd)
	}

	addr := listDomain[0]
	domain := listDomain[1]

	addr = Replacer(addr, strIteration(defaultMask, emailLength), start, emailEnd)

	return addr + "@" + domain
}
