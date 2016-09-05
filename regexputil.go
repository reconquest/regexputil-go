package regexputil

import "regexp"

func Subexp(exp *regexp.Regexp, matches []string, subexp string) string {
	for index, name := range exp.SubexpNames() {
		if index >= len(matches) {
			continue
		}

		if name == subexp {
			return matches[index]
		}
	}

	return ""
}

func MatchString(exp string, data string) ([]string, error) {
	re, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}

	return re.FindStringSubmatch(data), nil
}

func Match(exp string, data []byte) ([][]byte, error) {
	re, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}

	return re.FindSubmatch(data), nil
}
