package leetcode

import (
	"fmt"
	"testing"
)

type question999999999 struct {
	para999999999
	ans999999999
}

// para 是参数
// one 代表第一个参数
type para999999999 struct {
	nums string
}

// ans 是答案
// one 代表第一个答案
type ans999999999 struct {
	one string
}

func Test_Problem999999999(t *testing.T) {

	qs := []question999999999{

		{
			para999999999{"8000"},
			ans999999999{"85000"},
		},

		{
			para999999999{"-8000"},
			ans999999999{"-58000"},
		},

		{
			para999999999{"12090"},
			ans999999999{"512090"},
		},

		{
			para999999999{"555"},
			ans999999999{"5555"},
		},

		{
			para999999999{"27346209830709182346"},
			ans999999999{"527346209830709182346"},
		},

		{
			para999999999{"982346"},
			ans999999999{"9852346"},
		},

		{
			para999999999{"-1"},
			ans999999999{"-15"},
		},

		{
			para999999999{"0"},
			ans999999999{"50"},
		},

		{
			para999999999{"-2010"},
			ans999999999{"-20105"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 999999999------------------------\n")

	for _, q := range qs {
		_, p := q.ans999999999, q.para999999999
		fmt.Printf("【input】:%v       【output】:%v\n", p, addFive(p.nums))
	}
	fmt.Printf("\n\n\n")
}
