### **Week 1: Fundamentals of Golang**

#### **Day 1: Introduction & Go Syntax**
- **Theory**: Go data types, variables, constants, and basic syntax.
- **Coding**: Write simple programs like a calculator or string manipulator.
- **Tasks**:
  - What are Go’s unique features compared to other languages?
  - Implement a basic program that takes user input and outputs the result.

##### Tasks

**1. Basic Program:**

Write a Go program that takes two numbers as input, performs addition, subtraction, multiplication, and division, and then prints the results.

**2. String Manipulation:**

Create a program that takes a string as input, converts it to uppercase, and prints the number of characters in the string.

**3. Constants and Variables:**

Define two constants, PI and E, for the mathematical constants Pi (3.14) and Euler's number (2.71). Then write a program that calculates the circumference of a circle using the radius input by the user and prints both the circumference and the exponential of the radius.

Formula:
```
Circumference = 2 * PI * radius
Exponential = E^radius (you can use math.Pow to calculate the power)
```

**4. Data Types Practice:**

Write a Go program that declares variables of different types (int, float64, string, bool) and prints them along with their types.

**5. Type Conversion:**

Write a program that converts a float64 value to an int and prints both the original and the converted values.

**6. Challenge:**

Create a program that prompts the user for their age, then calculates and prints how many days old they are (assuming 365 days per year).

```
Enter your age: 25
You are approximately 9125 days old.
```
#### **Day 2: Control Structures**
- **Theory**: Conditionals (if-else), loops (for, range), and switch.
- **Coding**: Create a number guessing game using control structures.
- **Tasks**:
  - What is the difference between a regular `for` loop and `range` in Go?
  - Implement a loop that sums all numbers in an array.

##### Tasks

**1. FizzBuzz with a Twist:**

Write a program that prints the numbers from 1 to 100. For multiples of 3, print "Fizz" instead of the number, and for multiples of 5, print "Buzz". For numbers that are multiples of both 3 and 5, print "FizzBuzz". Additionally, if a number is a prime, print "Prime" instead of the number.
```
1 Prime Fizz Prime Buzz Fizz Prime 8 Fizz Buzz Prime Fizz 14 FizzBuzz
```

**2. Palindrome Checker:**

Create a program that takes a string as input and uses a loop to check whether it is a palindrome (reads the same backward and forward). Ignore spaces and case when determining if it is a palindrome.

```
Enter a string: RaceCar
It is a palindrome.
```

**3. Grade Calculator:**

Write a program that takes a series of grades (0-100) as input and uses a loop to calculate the ```average grade```, the ```highest grade```, and the ```lowest grade```. The program should stop when the user inputs a negative number.

```
Enter grade (negative number to end): 85
Enter grade: 92
Enter grade: 76
Enter grade: -1
Average Grade: 84.3
Highest Grade: 92
Lowest Grade: 76
```

**4. Factorial Calculation with Error Handling:**

Create a program that prompts the user for a non-negative integer and calculates the factorial of that number using a loop. If the user enters a negative number, display an error message and prompt them to enter a non-negative integer.

```
Enter a non-negative integer: 5
Factorial: 120

Enter a non-negative integer: -3
Error: Please enter a non-negative integer.
```

**5. Number Guessing Game with Limited Attempts:**

Write a number guessing game where the program randomly selects a number between 1 and 50. The user has up to 5 attempts to guess the number. After each incorrect guess, provide a hint ("too high" or "too low"). If the user fails to guess within 5 attempts, reveal the correct number.

```
Guess the number between 1 and 50: 20
Too low. Try again.
Guess the number: 30
Too high. Try again.
...
Sorry, you've used all attempts. The correct number was 27.

```

#### **Day 3: Functions and Error Handling**
- **Theory**: Function syntax, multiple return values, and error handling in Go.
- **Coding**: Create functions that perform mathematical operations and handle errors.
- **Tasks**:
  - Explain Go’s approach to error handling.
  - Write a function that divides two numbers and returns both the quotient and any error.

##### Tasks

**Task 1: Multiple Return Values**

Write a function calculate(a, b int) that returns the sum, difference, product, and quotient (if b is not zero) of two integers. Return an error if the second integer is zero.

Example:
```go
sum, diff, product, quotient, err := calculate(10, 0)
```

**Task 2: Variadic Function**

Write a variadic function findMax that takes any number of integer arguments and returns the largest one.

Example:
```go
max := findMax(3, 9, 2, 5, 7)
```

**Task 3: Custom Error**

Create a custom error InsufficientFundsError that includes the amount requested and the amount available. Use it in a withdraw function that checks if there are enough funds before proceeding.

Example:
```go
type InsufficientFundsError struct {
    Requested, Available float64
}

func (e *InsufficientFundsError) Error() string {
    return fmt.Sprintf("requested: %.2f, available: %.2f", e.Requested, e.Available)
}

func withdraw(balance, amount float64) error {
    if amount > balance {
        return &InsufficientFundsError{Requested: amount, Available: balance}
    }
    return nil
}
```

**Task 4: Wrapping and Unwrapping Errors**

Create a readFile function that reads a file and returns an error if the file cannot be opened. Wrap the error using fmt.Errorf with %w, and then unwrap it to display the original error.

Example:
```go
func readFile(filename string) error {
    _, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    return nil
}

func main() {
    err := readFile("nonexistent.txt")
    if err != nil {
        fmt.Println(err)
        fmt.Println("Original error:", errors.Unwrap(err))
    }
}
```

#### **Day 4: Go Data Structures**
- **Theory**: Arrays, slices, maps, and structs.
- **Coding**: Implement a contact list using structs and slices.
- **Tasks**:
  - How does Go handle slices differently from arrays?
  - Create a map of items and their prices, and write a function that returns the total price of selected items.

#### Tasks

**Task 1: Structs and Methods – Library System**

Create a library system where you can add books, lend books to users, and return books. Each book should be represented by a struct with attributes like title, author, and availability status. You should implement methods for:

1. Adding a book to the library.
2. Lending a book to a user (change its availability).
3. Returning a book (mark it as available again).

Additionally, handle edge cases like trying to lend a book that’s already lent out or returning a book that wasn’t lent out.

Key Points:
- Use a slice of structs to store multiple books.
- Each method should be able to modify the availability status of a book.

**Task 2: Nested Maps and Slices – User Data Management**

Write a program to manage user data for an application. Each user has a username, and each username has multiple associated attributes like age, email, and a list of favorite movies. You should use a map to store this data where:
- The keys are usernames (strings).
- The values are structs containing the user's information (```age```, ```email```, and a slice of ```favoriteMovies```).

Implement the following features:
1. Add a new user.
2. Update a user's information.
3. Retrieve a user's favorite movies.
4. List all users who have a specific movie in their favorites.

Key Points:
- Use maps for quick lookups and slices for storing lists of favorite movies.
- Ensure you handle updating a user who already exists and avoid duplicate entries in their favorite movies.

**Task 3: Slices and Maps – Most Frequent Words**

Write a program that takes a long string of text and returns the top 5 most frequent words and their respective counts. Use a map to store the frequency of each word and a slice to store the words sorted by their frequency.

Steps:
1. Split the text into words (consider punctuation and case sensitivity).
2. Use a map to count occurrences of each word.
3. Sort the words by their frequency and return the top 5.

Key Points:
- Use slices for sorting.
- Ensure that the map keys are the words, and the values are their frequencies.

**Task 4: Slices – Merging Intervals**

Given a list of intervals represented by pairs of integers (e.g., ```[1, 3]```, ```[2, 6]```), write a function that merges all overlapping intervals and returns a list of merged intervals.

For example, if the input is ```[[1, 3], [2, 6], [8, 10], [15, 18]]```, the output should be ```[[1, 6], [8, 10], [15, 18]]```.

Key Points:
- Use slices to represent intervals.
- You need to sort the intervals by their start time before merging them.
- Ensure your function can handle edge cases, such as when no intervals overlap or when all intervals overlap.

**Task 5: Structs and Maps – Product Inventory System**

Create a product inventory system where each product is represented by a struct with fields like ```ID```, ```name```, ```quantity```, and ```price```. The system should allow:
1. Adding new products.
2. Updating the quantity of existing products.
3. Deleting products.
4. Displaying a summary of the inventory (total value of all products).

Use a map to store the products, where the ```ID``` is the key and the product struct is the value.

Key Points:
- Handle adding, updating, and deleting products efficiently.
- Ensure that the system updates both the quantity and the price when necessary.
- Calculate the total value of the inventory by multiplying the price and quantity for each product.

#### **Day 5: Pointers**
- **Theory**: Go pointers, reference vs value types.
- **Coding**: Write a function that modifies a value using pointers.
- **Tasks**:
  - How do pointers work in Go?
  - Create a program that swaps two numbers using pointers.

#### **Day 6: Review & Practice**
- **Coding**: Build a small project that combines everything from this week (e.g., a to-do list app).
- **Tasks**:
  - Review Go’s error handling practices.
  - Solve common interview questions related to Go basics.

---

### **Week 2: Golang Intermediate Concepts**

#### **Day 1: Interfaces & Polymorphism**
- **Theory**: Go interfaces, polymorphism, and type assertions.
- **Coding**: Create an interface for different shapes and calculate areas.
- **Tasks**:
  - Explain Go’s interface mechanism.
  - Implement a polymorphic function that accepts any shape and returns its area.

#### **Day 2: Concurrency in Go**
- **Theory**: Goroutines, channels, and the Go scheduler.
- **Coding**: Build a program using goroutines to handle multiple tasks concurrently.
- **Tasks**:
  - What are goroutines, and how do they differ from threads?
  - Write a program where multiple goroutines communicate via channels.

#### **Day 3: Synchronization & Mutexes**
- **Theory**: Synchronization using channels, mutexes, and WaitGroups.
- **Coding**: Implement a counter using mutexes to avoid race conditions.
- **Tasks**:
  - When would you use a mutex over channels?
  - Create a producer-consumer problem using channels.

#### **Day 4: Packages & Modules**
- **Theory**: Organizing code with packages, modules, and third-party libraries.
- **Coding**: Build a simple package and import it into a main application.
- **Tasks**:
  - How do Go modules differ from packages?
  - Write a package that handles basic arithmetic operations.

#### **Day 5: Go Testing & Benchmarking**
- **Theory**: Testing in Go using `testing` package, writing unit tests, and benchmarks.
- **Coding**: Write unit tests for existing functions and benchmark them.
- **Tasks**:
  - How does Go’s `testing` package work?
  - Write test cases for a program that calculates discounts on items.

#### **Day 6: Review & Practice**
- **Coding**: Build a more complex project involving interfaces, concurrency, and error handling.
- **Tasks**:
  - Review concurrency and testing practices.
  - Solve concurrency-related interview problems.

---

### **Week 3: PostgreSQL Fundamentals**

#### **Day 1: Introduction to PostgreSQL**
- **Theory**: Basic PostgreSQL architecture, installation, and database setup.
- **Coding**: Create a database and add a few tables (e.g., users and orders).
- **Tasks**:
  - What are the key features of PostgreSQL compared to other databases?
  - Write queries to insert and fetch data from the database.

#### **Day 2: SQL Basics**
- **Theory**: Basic SQL queries (SELECT, INSERT, UPDATE, DELETE).
- **Coding**: Write SQL queries for CRUD operations on the database.
- **Tasks**:
  - What is the difference between an `INNER JOIN` and `LEFT JOIN`?
  - Implement a query that returns orders and their associated users.

#### **Day 3: Table Relationships & Joins**
- **Theory**: One-to-one, one-to-many, and many-to-many relationships.
- **Coding**: Create tables with foreign key constraints and perform JOIN operations.
- **Tasks**:
  - Explain different types of joins in PostgreSQL.
  - Write queries to retrieve data from related tables.

#### **Day 4: Indexes & Query Optimization**
- **Theory**: Indexes, query planning, and performance optimization.
- **Coding**: Create indexes and measure their effect on query performance.
- **Tasks**:
  - How do indexes improve performance in PostgreSQL?
  - Write queries and add indexes to optimize a slow query.

#### **Day 5: Transactions & ACID Properties**
- **Theory**: Transactions, isolation levels, and the ACID model.
- **Coding**: Write a transaction that updates multiple tables.
- **Tasks**:
  - Explain the ACID properties of a transaction.
  - Implement a transaction that ensures data consistency across multiple operations.

#### **Day 6: Review & Practice**
- **Coding**: Build a small application that interacts with PostgreSQL (e.g., a user management system).
- **Tasks**:
  - Review PostgreSQL indexing and transaction practices.
  - Solve SQL query problems commonly asked in interviews.

---

### **Week 4: Advanced PostgreSQL & Integration with Golang**

#### **Day 1: Stored Procedures & Functions**
- **Theory**: Writing stored procedures and functions in PostgreSQL.
- **Coding**: Create a stored procedure that calculates discounts based on purchase history.
- **Tasks**:
  - What are the advantages of using stored procedures?
  - Write a function that calculates totals for different orders.

#### **Day 2: Triggers & Event-Driven Systems**
- **Theory**: PostgreSQL triggers and their use cases.
- **Coding**: Implement a trigger that logs changes to a table.
- **Tasks**:
  - When should you use triggers in a database?
  - Create a trigger that tracks updates to the user’s table.

#### **Day 3: Connecting Go with PostgreSQL**
- **Theory**: Using the `pgx` and `database/sql` packages to interact with PostgreSQL in Go.
- **Coding**: Build a Go application that performs CRUD operations on a PostgreSQL database.
- **Tasks**:
  - How does Go’s `pgx` package differ from `database/sql`?
  - Write Go code that connects to PostgreSQL and performs database operations.

#### **Day 4: Optimizing Go-PostgreSQL Performance**
- **Theory**: Optimizing queries in Go, connection pooling, and handling large datasets.
- **Coding**: Implement connection pooling and handle a large data query efficiently.
- **Tasks**:
  - What is connection pooling, and how does it help performance?
  - Create a Go program that handles large dataset queries efficiently.

#### **Day 5: Security Best Practices**
- **Theory**: Securing PostgreSQL databases (encryption, user roles, access control).
- **Coding**: Implement role-based access control for a database.
- **Tasks**:
  - How can you secure a PostgreSQL database in production?
  - Write a query that grants permissions to different user roles.

#### **Day 6: Final Project & Review**
- **Coding**: Build a full-stack project integrating Go and PostgreSQL (e.g., a product catalog or a blog platform).
- **Tasks**:
  - Review integration strategies between Go and PostgreSQL.
  - Practice interview questions involving database design and Go-based APIs.
