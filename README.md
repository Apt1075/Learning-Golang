# Learning Go (Golang) — Core Concepts & Fundamentals

Welcome to my Go (Golang) learning repository! This project serves as a structured collection of Go programs designed to practice and demonstrate the core fundamentals of the language. It covers basic syntax, variables, user input, custom packages, control flow, data structures, and error handling.

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
*   **`array/`**
    *   [`array/main.go`](file:///d:/Go-Lang/array/main.go) — Basic array syntax, length declaration, indexing, and array type formatting.
*   **`slice/`**
    *   [`slice/main.go`](file:///d:/Go-Lang/slice/main.go) — Working with slices, checking capacity (`cap()`) and length (`len()`), dynamically adding elements via `append()`, using variadic append (`append(s1, s2...)`), and creating slices with `make()`.
*   **`error_handling/`**
    *   [`error_handling/main.go`](file:///d:/Go-Lang/error_handling/main.go) — Custom error creation (`fmt.Errorf`), returning errors alongside results (multiple returns), and handling them using idiomatic `if err != nil` patterns.
*   **`test/`**
    *   [`test/test.go`](file:///d:/Go-Lang/test/test.go) — A basic demonstration of creating and referencing custom packages/modules in Go.
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

### 5. Arrays and Slices
*   **Arrays (`array/`):** Fixed-size collections of homogeneous elements (e.g., `[5]int`).
*   **Slices (`slice/`):** Dynamic, flexible views into arrays. Topics include:
    *   Getting length (`len`) and capacity (`cap`).
    *   Appending elements dynamically.
    *   Concatenating two slices with the variadic `...` operator.
    *   Pre-allocating slice memory using the `make()` function.

### 6. Error Handling
In [`error_handling/main.go`](file:///d:/Go-Lang/error_handling/main.go):
*   Implementing idiomatic Go error handling.
*   Returning an `error` interface as the last return value from a function.
*   Creating formatted errors using `fmt.Errorf`.
*   Writing robust code to prevent runtime panics by validating function execution outputs.

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

### Functions
```bash
go run ./func/main.go
```

### Control Flow Examples
```bash
go run ./ifelse/main.go
go run ./switchCase/main.go
```

### Data Structures (Arrays & Slices)
```bash
go run ./array/main.go
go run ./slice/main.go
```

### Error Handling
```bash
go run ./error_handling/main.go
```
