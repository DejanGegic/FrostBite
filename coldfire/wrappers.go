package coldfire

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// F is a wrapper for the Sprintf function.
func F(str string, arg ...interface{}) string {
	return fmt.Sprintf(str, arg...)
}

func f(s string, arg ...interface{}) string {
	return fmt.Sprintf(s, arg...)
}

// StrToInt converts a string into an integer.
func StrToInt(string_integer string) int {
	// i, _ := strconv.ParseInt(string_integer, 10, 32)
	i, _ := strconv.Atoi(string_integer)
	return i
}

// IntToStr converts an integer into a string.
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// RandomInt returns an integer within a given range.
func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
