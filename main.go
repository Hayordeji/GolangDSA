package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	var name, nameError = reader.ReadString('\n')
	if nameError != nil {
		fmt.Println("Error: ", nameError.Error())
		return
	}
	if strings.TrimSpace(name) == "ayodeji" || strings.TrimSpace(name) == "Ayodeji" {
		fmt.Println("Your name can't be Ayodeji because the owner of this app is Ayodeji")
		return
	}

	fmt.Print("Enter your age: ")

	var age, ageErr = reader.ReadString('\n')
	if ageErr != nil {
		fmt.Println("Error: ", ageErr.Error())
	}
	ageNum, ageNumErr := strconv.Atoi(strings.TrimSpace(age))
	if ageNumErr != nil {
		fmt.Println("Error: ", ageNumErr.Error())
		return
	}
	if ageNum < 18 {
		fmt.Println("Your age is too low")
		return
	}
	fmt.Println("Welcome ", name, "you said you are", ageNum, "years old.")
}
