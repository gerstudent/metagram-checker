package main

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		s    string
		t    string
		want bool
	}{
		{"example", "etample", true},
		{"examp", "etample", false},
		{"example", "ethmple", false},
		{"test", "text", true},
		{"tett", "text", true},
		{"tett", "texx", false},
		{"кот", "кит", true},
		{"кот", "ктт", true},
		{"кот", "кот", false},
	}
	functions = []func(s, t string) bool{isMetagramm}
)

func TestIsMetagramm(t *testing.T) {
	for _, fun := range functions {
		for _, test := range tests {
			test := test

			t.Run(fmt.Sprint(test.s, " ", test.t, " ", test.want), func(t *testing.T) {
				t.Parallel()
				if have := fun(test.s, test.t); have != test.want {
					t.Errorf(`
					Ожидаемый результат: %+v
					Полученный результат: %+v`, test.want, have)
				}
			})

		}
	}
}
