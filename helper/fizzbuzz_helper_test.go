package helper

import (
	"testing"
)

func TestFizzBuzzHelper(t *testing.T) {

	params := FizzBuzzHelperParams{}

	// Failed Int1 params
	{
		ret, err := FizzBuzzHelper(params)
		if ret != nil {
			t.Errorf("Return must be nil")
		}
		if err == nil {
			t.Errorf("Error must not be nil")
		}
		if err != nil && err.Error() != "Int1 need to be greater than 0" {
			t.Errorf(
				"Error msg not good \nwe need : %s \nand we get: %s\n",
				"Int1 need to be greater than 0",
				err.Error(),
			)
		}
	}

	// Setup for other test
	params.Int1 = 3

	// Failed Int2 params
	{
		ret, err := FizzBuzzHelper(params)
		if ret != nil {
			t.Errorf("Return must be nil")
		}
		if err == nil {
			t.Errorf("Error must not be nil")
		}
		if err != nil && err.Error() != "Int2 need to be greater than 0" {
			t.Errorf(
				"Error msg not good \nwe need : %s \nand we get: %s\n",
				"Int2 need to be greater than 0",
				err.Error(),
			)
		}
	}

	// Setup for other test
	params.Int2 = 5

	// Failed Limit params
	{
		ret, err := FizzBuzzHelper(params)
		if ret != nil {
			t.Errorf("Return must be nil")
		}
		if err == nil {
			t.Errorf("Error must not be nil")
		}
		if err != nil && err.Error() != "Limit need to be greater than 0" {
			t.Errorf(
				"Error msg not good \nwe need : %s \nand we get: %s\n",
				"Limit need to be greater than 0",
				err.Error(),
			)
		}
	}

	// Setup for other test
	params.Limit = 100

	// Failed Str1 params
	{
		ret, err := FizzBuzzHelper(params)
		if ret != nil {
			t.Errorf("Return must be nil")
		}
		if err == nil {
			t.Errorf("Error must not be nil")
		}
		if err != nil && err.Error() != "Str1 need to be filled" {
			t.Errorf(
				"Error msg not good \nwe need : %s \nand we get: %s\n",
				"Str1 need to be filled",
				err.Error(),
			)
		}
	}

	// Setup for other test
	params.Str1 = "fizz"

	// Failed Str1 params
	{
		ret, err := FizzBuzzHelper(params)
		if ret != nil {
			t.Errorf("Return must be nil")
		}
		if err == nil {
			t.Errorf("Error must not be nil")
		}
		if err != nil && err.Error() != "Str2 need to be filled" {
			t.Errorf(
				"Error msg not good \nwe need : %s \nand we get: %s\n",
				"Str2 need to be filled",
				err.Error(),
			)
		}
	}

	// Setup for other test
	params.Str2 = "buzz"

	// Test ok
	{
		ret, err := FizzBuzzHelper(params)
		if err != nil {
			t.Errorf("Error must be nil : error %s\n", err.Error())
		}
		if ret == nil {
			t.Errorf("Return must be not nil")
		}

		want := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14",
			"fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz",
			"31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz",
			"46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz",
			"61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz",
			"76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92",
			"fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"}

		if len(want) != len(ret) {
			t.Errorf("number of result wanted not correct: \n nb Wanted: %d\n nb result: %d\n", len(want), len(ret))
			return
		}

		for i := 0; i < len(want); i++ {
			if want[i] != ret[i] {
				t.Errorf("Error in case %d\n value wanted: %s\n value get: %s\n", i, want[i], ret[i])
			}
		}

	}
}
