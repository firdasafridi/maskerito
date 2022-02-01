package masker

import "strings"

const (
	nameStart  = 1
	nameEndLow = 2
	nameEnd    = 3
	nameLength = 2
)

// Name mask the second letter and the third letter
//
// Example:
//   input: ABCD
//   output: A**D
func Name(input string) (output string) {
	count := len([]rune(input))

	if count == 0 {
		return ""
	}

	// if has space
	if strs := strings.Split(input, " "); len(strs) > 1 {
		tmp := make([]string, len(strs))
		for idx, str := range strs {
			tmp[idx] = Name(str)
		}
		return strings.Join(tmp, " ")
	}

	if count == 2 || count == 3 {
		return Replacer(input, strIteration(defaultMask, nameLength), nameStart, nameEndLow)
	}

	if count > 3 {
		return Replacer(input, strIteration(defaultMask, nameLength), nameStart, nameEnd)
	}

	return strIteration(defaultMask, nameLength)
}
