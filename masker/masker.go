package masker

var (
	defaultMask = "*"
)

func ChangeDefaultMask(mask string) {
	defaultMask = mask
}

func strIteration(str string, length int) string {
	var value string
	for i := 1; i <= length; i++ {
		value += str
	}
	return value
}

func Replacer(str string, overlay string, start int, end int) (overlayed string) {
	rStr := []rune(str)
	count := len([]rune(rStr))

	if count == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if start > count {
		start = count
	}
	if end < 0 {
		end = 0
	}
	if end > count {
		end = count
	}
	if start > end {
		tmp := start
		start = end
		end = tmp
	}

	overlayed = ""
	overlayed += string(rStr[:start])
	overlayed += overlay
	overlayed += string(rStr[end:])
	return overlayed
}
