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
