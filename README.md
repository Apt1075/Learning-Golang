# Learning Go (Golang) — Core Concepts & Fundamentals

Welcome to my Go (Golang) learning repository! This project serves as a structured collection of Go programs designed to practice and demonstrate the core fundamentals of the language. It covers basic syntax, variables, user input, custom packages, control flow, loops, data structures, and error handling.

---

## 📂 Repository Structure

The project is structured as follows:

*   **`main.go`** — Entry point demonstrating basic variables, constants, formatted output, and standard input scanning (`bufio`).
*   **`func/`**
    *   [`func/main.go`](file:///d:/Go-Lang/func/main.go) — Demonstration of functions, including parameter passing, return values, basic arithmetic logic, and simple conditional validation (e.g., prevention of division by zero).
*   **`ifelse/`**
    *   [`ifelse/main.go`](file:///d:/Go-Lang/ifelse/main.go) — Examples of conditional statements (`if`, `else if`, `else`), including inline variable initialization statements.
*   **`switchCase/`**
    *   [`switchCase/main.go`](file:///d:/Go-Lang/switchCase/main.go) — Examples of `switch` statements for control flow.
*   **`forloop/`**
    *   [`forloop/main.go`](file:///d:/Go-Lang/forloop/main.go) — A deep dive into Go's loop control structures (standard `for` loops, while-style, infinite loops with `break`, iterating over arrays/slices with `range`, nested loops, and loop control with `continue`).
*   **`array/`**
    *   [`array/main.go`](file:///d:/Go-Lang/array/main.go) — Basic array syntax, length declaration, indexing, and array type formatting.
*   **`slice/`**
    *   [`slice/main.go`](file:///d:/Go-Lang/slice/main.go) — Working with slices, checking capacity (`cap()`) and length (`len()`), dynamically adding elements via `append()`, using variadic append (`append(s1, s2...)`), and creating slices with `make()`.
*   **`maps/`**
    *   [`maps/main.go`](file:///d:/Go-Lang/maps/main.go) — Basic maps setup using `make(map[keyType]valueType)`, key insertion/updates, and key-existence check using the "comma-ok" idiom (`value, exists := myMap[key]`).
*   **`error_handling/`**
    *   [`error_handling/main.go`](file:///d:/Go-Lang/error_handling/main.go) — Custom error creation (`fmt.Errorf`), returning errors alongside results (multiple returns), and handling them using idiomatic `if err != nil` patterns.
*   **`test/`**
    *   [`test/test.go`](file:///d:/Go-Lang/test/test.go) — A basic demonstration of creating and referencing custom packages/modules in Go.
*   **`EvenOROddNumber/`**
    *   [`EvenOROddNumber/Even_or_Odd_number.go`](file:///d:/Go-Lang/EvenOROddNumber/Even_or_Odd_number.go) — Program to check whether a given number is even or odd.
*   **`PositiveAndNegativeNumber/`**
    *   [`PositiveAndNegativeNumber/Positive_or_Negative_number.go`](file:///d:/Go-Lang/PositiveAndNegativeNumber/Positive_or_Negative_number.go) — Program to check whether a given number is positive, negative, or zero.
*   **`SumofFirstNNaturalNumbers/`**
    *   [`SumofFirstNNaturalNumbers/SumofFirstNNaturalNumbers.go`](file:///d:/Go-Lang/SumofFirstNNaturalNumbers/SumofFirstNNaturalNumbers.go) — Program to calculate the sum of the first N natural numbers.
*   **`go.mod`** — The module definition file.

---

## 🛠️ Key Topics Covered

### 1. Variables, Data Types & Formatting
In [`main.go`](file:///d:/Go-Lang/main.go), we practice declaring and using basic Go types:
*   `string`, `int`, `bool`, `float64`
*   Implicit declarations with short variable declaration operator `:=`
*   Constants (`const`)
*   Formatted printing using `fmt.Printf` (e.g., formatting floating-point numbers with `%.3f`)

### 2. Standard Input/Output
We explore different ways to read user input from the console:
*   Using `fmt.Scanf`
*   Using a buffered reader `bufio.NewReader(os.Stdin)` to read lines containing spaces

### 3. Custom Functions
In [`func/main.go`](file:///d:/Go-Lang/func/main.go), several types of functions are implemented:
*   **Simple Function:** `Helloworld()`
*   **Function with local variables:** `AddtionTwoNumber()`
*   **Function with parameters:** `ReminderTwoNumber(a, b int)`
*   **Function with parameters and returns:** `DivisionTwoNumber(a, b float64) int` (includes check for division by zero)

### 4. Control Flow (Conditionals & Switch-Case)
*   **If-Else (`ifelse/`):** Idiomatic conditional structures including variables initialized directly within the `if` statement block (scope-restricted variables).
*   **Switch-Case (`switchCase/`):** Conditional branching using cleanly structured `switch` blocks and default statements.

### 5. Loop Operations (`forloop/`)
Go features a single looping construct—the `for` loop—which we use to implement:
*   **Standard Counter Loop:** `for i := 1; i <= 5; i++`
*   **While-Style Loop:** `for condition`
*   **Infinite Loop:** `for { ... }` combined with `break`
*   **Range Loops (For-Each):** Iterating over arrays and slices with `for index, value := range collection`
*   **Loop Modifiers:** Skipping iterations with `continue` and breaking loops with `break`

### 6. Arrays and Slices
*   **Arrays (`array/`):** Fixed-size collections of homogeneous elements (e.g., `[5]int`).
*   **Slices (`slice/`):** Dynamic, flexible views into arrays. Topics include:
    *   Getting length (`len`) and capacity (`cap`).
    *   Appending elements dynamically.
    *   Concatenating two slices with the variadic `...` operator.
    *   Pre-allocating slice memory using the `make()` function.

### 7. Maps (`maps/`)
Go's built-in hash table data structure. Topics include:
*   Initializing maps with `make(map[KeyType]ValueType)`.
*   Setting and updating key-value pairs.
*   Checking for the existence of keys using the double-value assignment pattern:
    ```go
    value, key_exists := StudentName[key]
    ```

### 8. Error Handling
In [`error_handling/main.go`](file:///d:/Go-Lang/error_handling/main.go):
*   Implementing idiomatic Go error handling.
*   Returning an `error` interface as the last return value from a function.
*   Creating formatted errors using `fmt.Errorf`.
*   Writing robust code to prevent runtime panics by validating function execution outputs.

### 9. Basic Logic & Math Programs
Practice problems demonstrating conditionals and loops:
*   **Even or Odd:** [`EvenOROddNumber/Even_or_Odd_number.go`](file:///d:/Go-Lang/EvenOROddNumber/Even_or_Odd_number.go)
*   **Positive or Negative:** [`PositiveAndNegativeNumber/Positive_or_Negative_number.go`](file:///d:/Go-Lang/PositiveAndNegativeNumber/Positive_or_Negative_number.go)
*   **Sum of First N Natural Numbers:** [`SumofFirstNNaturalNumbers/SumofFirstNNaturalNumbers.go`](file:///d:/Go-Lang/SumofFirstNNaturalNumbers/SumofFirstNNaturalNumbers.go)

---

## 🚀 How to Run the Programs

Make sure you have Go installed on your machine. You can verify it by running:
```bash
go version
```

You can run any of the specific sections by navigating to their directories or pointing `go run` to the relevant file paths:

### Main Playground
```bash
go run main.go
```

### Functions & Error Handling
```bash
go run ./func/main.go
go run ./error_handling/main.go
```

### Control Flow & Loops
```bash
go run ./ifelse/main.go
go run ./switchCase/main.go
go run ./forloop/main.go
```

### Data Structures (Arrays, Slices & Maps)
```bash
go run ./array/main.go
go run ./slice/main.go
go run ./maps/main.go
```

### Basic Logic & Math Programs
```bash
go run ./EvenOROddNumber/Even_or_Odd_number.go
go run ./PositiveAndNegativeNumber/Positive_or_Negative_number.go
go run ./SumofFirstNNaturalNumbers/SumofFirstNNaturalNumbers.go
```
