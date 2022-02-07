package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	testCases := []struct {
		number1   uint32
		number2   uint32
		sumNumber uint32
		flag      bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
	}

	for _, v := range testCases {
		sum, overflow := AddUint32(v.number1, v.number2)
		assert.Equal(t, v.sumNumber, sum)
		assert.Equal(t, v.flag, overflow)
	}

}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	testCases := []struct {
		got  float64
		want float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
	}

	for _, v := range testCases {
		point := CeilNumber(v.got)
		assert.Equal(t, point, v.want)
	}
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	testCases := []struct {
		got  string
		want string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, v := range testCases {
		sortedString := AlphabetSoup(v.got)
		assert.Equal(t, sortedString, v.want)
	}
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	testCases := []struct {
		got  string
		str  uint
		want string
	}{
		{"!mysecret*", 2, "!m********"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, v := range testCases {
		pasString := StringMask(v.got, v.str)
		assert.Equal(t, pasString, v.want)
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	testCases := []struct {
		got  [2]string
		want string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}

	for _, v := range testCases {
		result := WordSplit(v.got)
		assert.Equal(t, result, v.want)
	}
}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	testCases := []struct {
		params []interface{}
		want   []interface{}
	}{
		{[]interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{[]interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, v := range testCases {
		set := VariadicSet(v.params...)
		assert.Equal(t, v.want, set)
	}
}
