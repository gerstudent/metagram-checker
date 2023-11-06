package main

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		s    string
		t    string
		n    int
		want bool
	}{
		{"example", "etample", 1, true},
		{"example", "ethmple", 2, true},
		{"example", "eXaMplE", 3, true},
		{"examp", "etample", 5, false},
		{"example", "etample", 4, false},
		{"example", "ejjjjjj", 5, false},
		{"test", "text", 1, true},
		{"tett", "tttT", 2, true},
		{"tett", "text", 2, false},
		{"tett", "text", 0, false},
		{"кот", "кто", 2, true},
		{"кот", "кот", 0, true},
		{"кот", "кот", 1, false},
		{"кот", "кип", 1, false},
	}
	functions = []func(s, t string, n int) bool{isMetagrammN}
)

func TestIsMetagrammN(t *testing.T) {
	for _, fun := range functions {
		for _, test := range tests {
			test := test

			t.Run(fmt.Sprint(test.s, " ", test.t, " ", test.n, " ", test.want), func(t *testing.T) {
				t.Parallel()
				if have := fun(test.s, test.t, test.n); have != test.want {
					t.Errorf(`
					Ожидаемый результат: %+v
					Полученный результат: %+v`, test.want, have)
				}
			})

		}
	}
}
