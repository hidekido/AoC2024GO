package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBaseParse(t *testing.T) {
	line := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := parse(line)
	expected := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	fmt.Println(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parse(\"%s\")=\"%s\"\n", line, result)
	}
}

func TestParse2(t *testing.T) {
	line := "xmul(2,4)%&mul[3,7]!@^don't()_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := parse2(line)
	expected := []string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parse2(\"%s\")=\"%v\"\n", line, result)
	}
}

func TestCalc(t *testing.T) {
	first := "mul(2,4)"
	res := calc(first)
	if 8 != res {
		t.Errorf("calc(\"%s\")=%d\n", first, res)
	}

	second := "mul(123,356)"
	res = calc(second)
	if 43788 != res {
		t.Errorf("calc(\"%s\")=%d\n", second, res)
	}
}

func TestCalcWhole(t *testing.T) {
	base := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	parsed := parse2(base)
	result, _ := calcWhole(parsed, true)
	if 48 != result {
		t.Errorf("calcWhole(\"%s\")=%d\n", parsed, result)
	}
}
