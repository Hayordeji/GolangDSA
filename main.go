package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	//DECLARE STRING
	var name string = "Ayodeji Shoga"
	topic := "DSA"

	//CREATES NEW STRING AND REASSIGN
	name = "Shoga Ayodeji"
	fmt.Println(name)

	//CONCATENATION
	message := name + " is learning " + topic + " in golang"
	fmt.Println(message)

	//SUBSTRINGS
	firstName := name[:6]
	lastName := name[6:]
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName)

	//REPLACE STRING
	fullMeaning := strings.Replace(topic, "DSA", "Data Structure and Algorithm", 1)
	fmt.Println(fullMeaning)

	//JOIN STRINGS
	fullName := strings.Join([]string{firstName, lastName}, " ")
	fmt.Println(fullName)

}

// CHECK IF STRING READS THE SAME BACKWARDS AS FORWARDS
func checkPalindrome(s string) bool {
	var filtered []rune

	//REMOVE SPECIAL CHARACTERS AND CONVERT CHARACTERS TO LOWERCASE
	for _, value := range s {
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			filtered = append(filtered, unicode.ToLower(value))
		}
	}

	//TWO POINTER TO CHECK IF PALINDROME
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--

	}
	return true
}

func findUniqueCharacter(s string) int {
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
		}
		filtered = append(filtered, unicode.ToLower(r))
	}

	index := 0
	for i := 0; i < len(filtered); i++ {
		if filtered[i] == filtered[index] {
			index++
		} else {
			return index
		}
	}
	return 0
}

func reverseWords(s string) string {

	substrings := strings.Fields(s)
	left, right := 0, len(substrings)-1

	for left < right {
		substrings[left], substrings[right] = substrings[right], substrings[left]
		left++
		right--
	}
	result := strings.Join(substrings, " ")

	return result
}
