package yaml

import (
	"fmt"
	"strings"
)

// TODO(maia): store yaml strings as []string so we don't need to parse strings for separators?
const separator = "---"

func ConcatYAML(yaml ...string) string {
	if len(yaml) == 0 {
		return ""
	} else if len(yaml) == 1 {
		return yaml[0]
	}

	result := yaml[0]
	for _, y := range yaml[1:] {
		result = concatYAML(result, y)
	}
	return result
}

func concatYAML(y1, y2 string) string {
	y1 = strings.TrimSpace(y1)
	y2 = strings.TrimSpace(y2)
	if !hasEndingSeparator(y1) && !hasStartingSeparator(y2) {
		return fmt.Sprintf("%s\n%s\n%s", y1, separator, y2)
	} else if hasEndingSeparator(y1) && hasStartingSeparator(y2) {
		y1 = trimEndingSeparator(y1)
	}
	return fmt.Sprintf("%s\n%s", y1, y2)
}

func hasStartingSeparator(y string) bool {
	line, _, _ := strings.Cut(y, "\n")
	return isSeparatorLine(line)
}

func hasEndingSeparator(y string) bool {
	line := y[strings.LastIndex(y, "\n")+1:]
	return isBareSeparatorLine(line)
}

func trimEndingSeparator(y string) string {
	return strings.TrimSpace(y[:strings.LastIndex(y, "\n")+1])
}

func isSeparatorLine(line string) bool {
	line = strings.TrimRight(line, " \t\r")
	return line == separator || strings.HasPrefix(line, separator+" ") || strings.HasPrefix(line, separator+"\t")
}

func isBareSeparatorLine(line string) bool {
	line = strings.TrimRight(line, " \t\r")
	if line == separator {
		return true
	}
	if !isSeparatorLine(line) {
		return false
	}
	return strings.HasPrefix(strings.TrimSpace(line[len(separator):]), "#")
}
