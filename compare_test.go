package main

import (
	"fmt"
	"testing"
)

func Test_compare_words_1(t *testing.T) {
	first := "Eko"
	second := "Ekom"
	result := Compare([]byte(first), []byte(second))
	if result != 1 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}

func Test_compare_words_2(t *testing.T) {
	first := "feyza"
	second := "Araç"
	result := Compare([]byte(first), []byte(second))
	if result != 2 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}

func Test_compare_words_3(t *testing.T) {
	first := "Traffic"
	second := "Comparison"
	result := Compare([]byte(first), []byte(second))
	if result != 2 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}

func Test_compare_words_4(t *testing.T) {
	first := "Attention"
	second := "attention"
	result := Compare([]byte(first), []byte(second))
	if result != 1 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}

func Test_compare_words_5(t *testing.T) {
	first := "Attention"
	second := "Attention"
	result := Compare([]byte(first), []byte(second))
	if result != 0 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}

func Test_compare_words_6(t *testing.T) {
	first := "Ekomatekise"
	second := "Ekom"
	result := Compare([]byte(first), []byte(second))
	if result != 2 {
		panic(fmt.Sprintf("%s and %s comparison is wrong.", first, second))
	}
}
