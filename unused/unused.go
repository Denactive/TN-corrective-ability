package unused

import (
	"fmt"
	"strconv"
)

func powBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

func shiftLeftPositiveBit(digit, pos int) int {
	bitsNum := getBinaryLength(digit)

	if pos < 0 || pos > bitsNum {
		return digit
	}

	return digit + powBinary(pos)
}

func bitsToString(digit int) string {
	var res string
	for i := powBinary(getBinaryLength(digit) - 1); digit != 0 && i > 0; i /= 2 {
		res += fmt.Sprintf("%d", digit / i)
		digit %= i
	}
	return res
}

func bitsToBytes(digit int) []byte {
	var res []byte
	for i := powBinary(getBinaryLength(digit) - 1); digit != 0 && i > 0; i /= 2 {
		res = append(res, byte(digit / i))
		digit %= i
	}
	return res
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func getBinaryLength(digit int) int {
	bitsNum := 0
	for digitCopy := digit; digitCopy / 2 != 0; digitCopy /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}