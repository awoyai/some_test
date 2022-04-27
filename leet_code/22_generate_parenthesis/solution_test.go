package main

import "testing"

func Test_solution(t *testing.T) {
	a := generateParenthesis(2)
	b := generateParenthesisDFS(2)
	print("%v", a)
	print("%v", b)
}