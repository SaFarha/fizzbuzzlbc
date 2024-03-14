package helper

import (
	"errors"
	"strconv"
)

// FizzBuzzHelperParams helper parameters
type FizzBuzzHelperParams struct {
	Int1  uint64
	Int2  uint64
	Limit uint64
	Str1  string
	Str2  string
}

// FizzBuzzHelper is the function that contain the use case of the fizzbuzz handler.
func FizzBuzzHelper(params FizzBuzzHelperParams) ([]string, error) {

	if params.Int1 == 0 {
		return nil, errors.New("Int1 need to be greater than 0")
	}

	if params.Int2 == 0 {
		return nil, errors.New("Int2 need to be greater than 0")
	}

	if params.Limit == 0 {
		return nil, errors.New("Limit need to be greater than 0")
	}

	if params.Str1 == "" {
		return nil, errors.New("Str1 need to be filled")
	}

	if params.Str2 == "" {
		return nil, errors.New("Str2 need to be filled")
	}

	var i uint64
	var fizzbuzzIndex = params.Int1 * params.Int2
	ret := make([]string, params.Limit)
	fizzbuzzStr := params.Str1 + params.Str2

	for i = 1; i <= params.Limit; i++ {
		switch {
		case i%fizzbuzzIndex == 0:
			ret[i-1] = fizzbuzzStr
		case i%params.Int1 == 0:
			ret[i-1] = params.Str1
		case i%params.Int2 == 0:
			ret[i-1] = params.Str2
		default:
			ret[i-1] = strconv.FormatUint(i, 10)
		}
	}

	return ret, nil
}
