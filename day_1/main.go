package main

import (
	"fmt"
	"math"
	"strings"
)

func TaskOne() {
	var a, b int
	fmt.Print("Enter first number: ")
	if _, err := fmt.Scanf("%d", &a); err != nil {
		panic(err)
	}
	fmt.Print("Enter second number: ")
	if _, err := fmt.Scanf("%d", &b); err != nil {
		panic(err)
	}
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
}

func TaskTwo() {
	var s string
	fmt.Print("Enter a string: ")
	if _, err := fmt.Scanf("%s", &s); err != nil {
		panic(err)
	}
	fmt.Println("Uppercase:", strings.ToUpper(s))
	fmt.Println("Length:", len(s))
}

func TaskThree() {
	var radius float64
	fmt.Print("Enter the radius of a circle: ")
	if _, err := fmt.Scanf("%f", &radius); err != nil {
		panic(err)
	}

	circ := 2 * math.Pi * radius
	fmt.Printf("Circumference: %.2f\n", circ)

	exponen := math.Pow(math.E, radius)
	fmt.Printf("Exponential: %.2f\n", exponen)
}

func TaskFour() {
	var intNumber int
	var floatNumber float64
	var str string
	var flag bool
	fmt.Print("Enter int: ")
	if _, err := fmt.Scanf("%d", &intNumber); err != nil {
		panic(err)
	}
	fmt.Print("Enter float: ")
	if _, err := fmt.Scanf("%f", &floatNumber); err != nil {
		panic(err)
	}
	fmt.Print("Enter string: ")
	if _, err := fmt.Scanf("%s", &str); err != nil {
		panic(err)
	}
	fmt.Print("Enter bool: ")
	if _, err := fmt.Scanf("%t", &flag); err != nil {
		panic(err)
	}

	fmt.Printf("int: %d\n", intNumber)
	fmt.Printf("float64: %.4f\n", floatNumber)
	fmt.Printf("int: %s\n", str)
	fmt.Printf("int: %t\n", flag)
}

func TaskFive() {
	var original float64
	fmt.Print("Write float value: ")
	if _, err := fmt.Scanf("%f", &original); err != nil {
		panic(err)
	}
	fmt.Printf("Original value: %.2f\n", original)
	fmt.Printf("Converted to int: %d\n", int(original))
}

func TaskSix() {
	var age int
	fmt.Print("Enter your age: ")
	if _, err := fmt.Scanf("%d", &age); err != nil {
		panic(err)
	}
	fmt.Printf("You are approximately %d days old.\n", 365*age)
}

func main() {
	// TaskOne()
	// TaskTwo()
	// TaskThree()
	// TaskFour()
	// TaskFive()
	TaskSix()
}
