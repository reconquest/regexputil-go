package regexputil

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubexp(t *testing.T) {
	test := assert.New(t)

	testcases := []struct {
		input  string
		exp    string
		subexp string
		value  string
	}{
		{
			`a b c`,
			`(?P<foo>.)\s+(?P<blah>.*)`,
			`foo`,
			`a`,
		},
		{
			`a b c`,
			`(?P<foo>.)\s+(?P<blah>.*)`,
			`blah`,
			`b c`,
		},
		{
			`a b c`,
			`(?P<foo>.)\s+(?P<blah>.*)`,
			`unknown`,
			``,
		},
		{
			`a b c`,
			`(?P<foo>1)?a`,
			`foo`,
			``,
		},
	}

	for index, testcase := range testcases {
		re, err := regexp.Compile(testcase.exp)
		if err != nil {
			test.NoError(err, fmt.Sprintf("testcase #%d", index))
			continue
		}

		matches := re.FindStringSubmatch(testcase.input)

		value := Subexp(re, matches, testcase.subexp)
		test.EqualValues(testcase.value, value)
	}
}
