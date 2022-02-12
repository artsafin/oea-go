package common

import (
	"fmt"
	"regexp"
)

var markdownLinkRe = regexp.MustCompile(`(?m)\[([^]]+)\]\(([^)]+)\)`)

// ReplaceAllStringSubmatchFunc is taken from https://gist.github.com/elliotchance/d419395aa776d632d897
func replaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			if v[i] == -1 || v[i+1] == -1 {
				groups = append(groups, "")
			} else {
				groups = append(groups, str[v[i]:v[i+1]])
			}
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func MarkdownToHTML(md string) string {
	return replaceAllStringSubmatchFunc(markdownLinkRe, md, func(g []string) string {
		return fmt.Sprintf(`<a href="%s">%s</a>`, g[1], g[2])
	})
}