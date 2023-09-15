package util

import "testing"

func TestLength(t *testing.T) {
	if Length("hello") != 5 {
		t.Error("Length of English string is wrong")
	}
	if Length("你好") != 4 {
		t.Error("Length of Chinese string is wrong")
	}
	if Length("你好hello") != 9 {
		t.Error("Length of Chinese and English string is wrong")
	}
	if Length("\033[31m你好\033[0m") != 4 {
		t.Error("Length of colored Chinese string is wrong")
	}
}
