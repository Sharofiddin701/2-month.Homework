1. package main

import "fmt"

func isPalindrome(x int) bool {
	var b int
	var t int = x
	if x < 0 {
		return false
	}
	for x > 0 {
		b = b*10 + x%10
		x = x / 10
	}
	if t == b {
		return true
	} else {
		return false
	}
}

func main() {
	num := 121
	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome.\n", num)
	} else {
		fmt.Printf("%d is not a palindrome.\n", num)
	}
}


2. package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.Replace(s, "CM", "CCCCCCCCC", -1) // 900
	s = strings.Replace(s, "CD", "CCCC", -1)      // 400
	s = strings.Replace(s, "XC", "XXXXXXXXX", -1) // 90
	s = strings.Replace(s, "XL", "XXXX", -1)      // 40
	s = strings.Replace(s, "IX", "IIIIIIIII", -1) // 9
	s = strings.Replace(s, "IV", "IIII", -1)      // 4

	var sum int
	for _, roman := range s {
		sum += romans[roman]
	}
	return sum
}

func main() {

	romanNumeral := "MCMXCIV"
	value := romanToInt(romanNumeral)
	fmt.Printf("%s in Roman numerals is equal to %d in integer.\n", romanNumeral, value)
}

3. package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    p := strs[0]
    for _, s := range strs {
        i := 0
        for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {}
        p = p[:i]
    }
    return p
}

func main() {
    strings := []string{"flower", "flow", "flight"}
    prefix := longestCommonPrefix(strings)
    fmt.Printf("The longest common prefix of %v is: %s\n", strings, prefix)
}