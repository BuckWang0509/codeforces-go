// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1768/D
// https://codeforces.com/problemset/status/1768/problem/D
func Test_cf1768D(t *testing.T) {
	testCases := [][2]string{
		{
			`4
2
2 1
2
1 2
4
3 4 1 2
4
2 4 3 1`,
			`0
1
3
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1768D)
}
