package logic

import (
	"strings"
)

const symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func decimalToSixtytwo(decimalNum int64) string {
	if decimalNum == 0 {
		return "0"
	}

	result := ""

	for decimalNum > 0 {
		remainder := decimalNum % 62
		result = string(symbols[remainder]) + result
		decimalNum = decimalNum / 62
	}

	return result
}

func sixtytwoToDecimal(sixtyTwoNum string) int64 {
	var decimalNum int64 = 0
	var power int64 = 1

	for i := len(sixtyTwoNum) - 1; i >= 0; i-- {
		index := strings.IndexByte(symbols, sixtyTwoNum[i])
		decimalNum += int64(index) * power
		power *= 62
	}

	return decimalNum
}
