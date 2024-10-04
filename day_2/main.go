package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TaskOne(n int) {
	fmt.Print("1 ")
	for i := 2; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz ")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz ")
			continue
		}
		if isPrime(i) {
			fmt.Print("Prime ")
			continue
		}
		fmt.Print(i, " ")
	}
}

func IsPolindrome(s string) {
	i := 0
	j := len(s) - 1
	isPol := true
	lowS := strings.ToLower(s)
	for i < j {
		if lowS[i] != lowS[j] {
			isPol = false
			break
		}
		i++
		j--
	}
	fmt.Print("It is ")
	if !isPol {
		fmt.Print("not ")
	}
	fmt.Println("a palindrome.")
}

func TaskThree() {
	var nums []int
	var inp int
	for {
		fmt.Print("Enter grade: ")
		if _, err := fmt.Scanf("%d", &inp); err != nil {
			panic(err)
		}
		if inp < 0 {
			break
		}
		nums = append(nums, inp)
	}
	if len(nums) == 0 {
		fmt.Printf("Average Grade: none\nHighest Grade: none\nLowest Grade: none\n")
		return
	}
	maxNum := nums[0]
	minNum := nums[0]
	sum := 0
	for _, val := range nums {
		if maxNum < val {
			maxNum = val
		}
		if minNum > val {
			minNum = val
		}
		sum += val
	}

	avgValue := float64(sum) / float64(len(nums))

	fmt.Printf("Average Grade: %.2f\nHighest Grade: %d\nLowest Grade: %d\n",
		avgValue,
		maxNum,
		minNum,
	)
}

func TaskFour() {
	fmt.Print("Enter a non-negative integer: ")
	var num int
	if _, err := fmt.Scanf("%d", &num); err != nil {
		panic(err)
	}
	if num < 0 {
		fmt.Println("Error: Please enter a non-negative integer.")
		return
	}
	var mult int64
	mult = 1
	for i := 2; i <= num; i++ {
		mult *= int64(i)
	}
	fmt.Println("Factorial:", mult)
}

func randNumber(maxVal int) int {
	return 1 + rand.Intn(maxVal+1)
}

func GuessGame() {
	rand.Seed(time.Now().UnixNano())

	fmt.Print("Guess the number between 1 and 50: ")
	var guess int
	num := randNumber(50)
	attemts := 0
	for attemts < 5 {
		if _, err := fmt.Scanf("%d", &guess); err != nil {
			panic(err)
		}
		if guess == num {
			fmt.Println("Hooray, it is correct!")
			return
		}
		if guess < num {
			fmt.Println("Too low. Try again.")
		} else {
			fmt.Println("Too high. Try again.")
		}
		attemts++
		if attemts == 5 {
			break
		}
		fmt.Print("Guess the number: ")
	}
	fmt.Println("Sorry, you've used all attempts. The correct number was", num)
}

func main() {
	// TaskOne(100)
	// IsPolindrome("ABCdcba")
	// TaskThree()
	// TaskFour()
	GuessGame()
}
