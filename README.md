# Learning Go (Golang) — Core Concepts & Fundamentals

Welcome to my Go (Golang) learning repository! This project serves as a structured collection of Go programs designed to practice and demonstrate the core fundamentals of the language. It covers basic syntax, variables, user input, custom packages, and function usage.

---

## 📂 Repository Structure

The project is structured as follows:

*   **`main.go`** — Entry point demonstrating basic variables, constants, formatted output, and standard input scanning (`bufio`).
*   **`func/`**
    *   [`func/main.go`](file:///d:/Go-Lang/func/main.go) — Demonstration of functions, including parameter passing, return values, basic arithmetic logic, and simple conditional validation (e.g., prevention of division by zero).
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

---

## 🚀 How to Run the Programs

Make sure you have Go installed on your machine. You can verify it by running:
```bash
go version
```

### Running the root entry point:
```bash
go run main.go
```

### Running the functions exercise:
Navigate to the `func` directory and execute it:
```bash
go run ./func/main.go
```
