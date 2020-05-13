package valid_parentheses

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	cases := map[string]bool{
		"()": true,
		"()[]{}": true,
		"(]":     false,
		"([)]": false,
		"{[]}": true,
	}

	for str, target := range cases {
		//validRes := Valid(s)
		validRes := Valid2(str)

		fmt.Println(str,target,validRes)

		if validRes != target {
			t.Fail()
		}
	}
}
