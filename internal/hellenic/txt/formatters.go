package txt

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
