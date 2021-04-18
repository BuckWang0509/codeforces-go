// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[0,1,1,3]`, `2`, 
			`[0,3,2,3]`,
		},
		{
			`[2,3,4,7]`, `3`, 
			`[5,2,6,5]`,
		},
		{
			`[0,1,2,2,5,7]`, `3`, 
			`[4,3,6,4,6,7]`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, getMaximumXor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-50/problems/maximum-xor-for-each-query/
