package txt

import "strings"

func boolToYesNoFlag(value bool) string {
	if value {
		return "Y"
	}
	return "N"
}

func underscore(value string) string {
	if value == "" {
		return "_"
	}
	return value
}

func limit(value string, lim uint16) string {
	if len(value) > int(lim) {
		return value[:lim]
	}
	return value
}

func lPadSp(in string, lim int) string {
	if lim-len(in) <= 0 {
		return in
	}
	return strings.Repeat(" ", lim-len(in)) + in
}

func rPadSp(in string, lim int) string {
	if lim-len(in) <= 0 {
		return in
	}
	return in + strings.Repeat(" ", lim-len(in))
}
