package leetcode

import (
	"fmt"
	"testing"
)

type question99999999 struct {
	para99999999
	ans99999999
}

// para 是参数
// one 代表第一个参数
type para99999999 struct {
	nums int
}

// ans 是答案
// one 代表第一个答案
type ans99999999 struct {
	one int
}

func Test_Problem99999999(t *testing.T) {

	qs := []question99999999{

		{
			para99999999{8000},
			ans99999999{85000},
		},

		{
			para99999999{-8000},
			ans99999999{-58000},
		},

		{
			para99999999{12090},
			ans99999999{85000},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 99999999------------------------\n")

	for _, q := range qs {
		_, p := q.ans99999999, q.para99999999
		fmt.Printf("【input】:%v       【output】:%v\n", p, addFive(p.nums))
	}
	fmt.Printf("\n\n\n")
}
