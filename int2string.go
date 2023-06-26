package int2string

import (
	"fmt"
	"strconv"
)

func ConvertSprintf(number int) string {
	return fmt.Sprintf("%d", number)
}

func ConvertFormatInt(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

func ConvertItoa(number int) string {
	return strconv.Itoa(number)
}