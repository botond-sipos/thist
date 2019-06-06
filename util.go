package thist

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Max(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	max := s[0]
	for _, x := range s {
		if x > max {
			max = x
		}
	}
	return max
}

func Min(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	max := s[0]
	for _, x := range s {
		if x < max {
			max = x
		}
	}
	return max
}

func Mean(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	var sum float64
	for _, x := range s {
		sum += x
	}
	return sum / float64(len(s))
}

func AbsFloats(s []float64) []float64 {
	res := make([]float64, len(s))
	for i, x := range s {
		res[i] = math.Abs(x)
	}
	return res
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ClearScreen() {
	fmt.Printf("\033[2J")
}

func StringsMaxLen(s []string) int {
	if len(s) == 0 {
		return 0
	}

	max := len(s[0])
	for _, x := range s {
		if len(x) > max {
			max = len(x)
		}
	}
	return max

}

func AutoLabel(s []float64, m float64) []string {
	res := make([]string, len(s))
	digits := -int(math.Log10(math.Abs(m) / 10))
	nf := false
	if Min(s) < 0 {
		nf = true
	}
	if math.Abs(m) == 0 {
		digits = 5
	}
	if math.Abs(m) < 1 && digits < 2 {
		digits = 3
	}
	if digits <= 0 {
		digits = 1
	}
	if digits > 8 {
		digits = 8
	}
	for i, x := range s {
		f := "%." + strconv.Itoa(digits-int(math.Log10(math.Abs(x)+1))) + "f"
		ff := f
		if nf && !(x < 0) {
			ff = " " + f
		}
		res[i] = fmt.Sprintf(ff, x)
	}
	return res
}

// LeftPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// RightPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// https://www.socketloop.com/tutorials/golang-aligning-strings-to-right-left-and-center-with-fill-example
func CenterPad2Len(s string, fill string, n int) string {
	div := n / 2

	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}