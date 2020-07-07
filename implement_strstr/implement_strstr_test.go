package implement_strstr

import (
	"testing"
)

func TestStrPos(t *testing.T) {

	tests := []map[string]string{
		{
		    "haystack":"hello",
			"needle":"ll",
		},
		{
			"haystack":"a",
			"needle":"a",
		},
		{
			"haystack":"hello",
			"needle":"",
		},
		{
			"haystack":"hello",
			"needle":"xh",
		},
	}
	results := []int{2,0,0,-1}

	caseNums := 4
	for i:=0; i< caseNums; i++ {
		if ret := strPos(tests[i]["haystack"], tests[i]["needle"]); ret != results[i] {
			t.Fatalf("case %d fail, test: %+v\n, actual: %d, expect:%d", i, tests[i], ret, results[i])
		}
	}
}
