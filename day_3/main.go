package main

import (
	"errors"
	"fmt"
	"os"
)

func TaskOne(a, b int) (int, int, int, int, error) {
	if b == 0 {
		return 0, 0, 0, 0, errors.New("second value cannot be 0")
	}
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	return sum, diff, product, quotient, nil
}

func FindMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxValue := nums[0]
	for _, n := range nums {
		if n > maxValue {
			maxValue = n
		}
	}
	return maxValue
}

type InsufficientFundsError struct {
	Requested, Available float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds, requested: %.2f, available: %.2f", e.Requested, e.Available)
}

func withdraw(balance, withdrawAmount float64) error {
	if balance < withdrawAmount {
		return &InsufficientFundsError{
			Requested: withdrawAmount,
			Available: balance,
		}
	}
	return nil
}

func openFile(path string) error {
	if _, err := os.Open(path); err != nil {
		return fmt.Errorf("failed to open file at %s: %w", path, err)
	}
	return nil
}

func main() {
	// Task 1
	// sum, diff, product, quotient, err := TaskOne(2, 5)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("sum: %d\ndiff: %d\nproduct: %d\nquotient: %d\n", sum, diff, product, quotient)

	// Task 2
	// maxValue := FindMax(10, 2, 3, 4, 5, 6, 8)
	// fmt.Println("Max value is:", maxValue)

	// Task 3
	// if err := withdraw(1000, 100); err != nil {
	// 	fmt.Printf("Error in withdraw: %s\n", err.Error())
	// } else {
	// 	fmt.Println("Successfull withdraw")
	// }

	err := openFile("./sample1.txt")
	if err != nil {
		fmt.Println("openFile error:", err.Error())
		fmt.Println("detailed error:", errors.Unwrap(err))
	}
}
