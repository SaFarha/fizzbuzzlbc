package utils

import (
	"encoding/base64"
	"fizzbuzzlbc/database/models"
	"strconv"
)

func CreateRequestUid(request models.FzParamRequest) string {

	limitStr := strconv.Itoa(int(request.Limit))

	int1Str := strconv.Itoa(int(request.Int1))

	int2Str := strconv.Itoa(int(request.Int2))

	convertStr := limitStr + ";" + int1Str + ";" + int2Str + ";" + request.Str1 + ";" + request.Str2

	return base64.StdEncoding.EncodeToString([]byte(convertStr))
}
