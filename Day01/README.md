# Go Language Basics

A beginner-friendly Go program covering the fundamental concepts of the Go programming language.

---

## 📁 File Structure

```
main.go
```

---

## 📚 Topics Covered

### 1. 🖨️ Printing Statements
Basic output using `fmt.Println`.

```go
fmt.Println("Hello! World")
```

---

### 2. 🔢 Data Types & Variables

Go supports the following core data types:

| Type | Description |
|------|-------------|
| `uint32`, `uint64` | Unsigned integers (no negative values) |
| `int8`, `int16`, `int32`, `int64` | Signed integers (positive & negative) |
| `string` | A sequence of characters |
| `bool` | Boolean value (`true` or `false`) |
| `float32`, `float64` | Floating-point numbers (decimals) |
| `byte` | Single byte of data (`int8`), used for characters or small integers |
| `rune` | Single Unicode code point (`int32`) |

**Constant declaration:**
```go
const number int = 32
```

**Multiple variable declaration:**
```go
var (
    first_name  string = "Kavishka"
    middle_name string = "Sasindu"
    last_name   string = "Mudithananda"
)
```

---

### 3. 📝 Variable Declaration

Go supports both **explicit** and **implicit** variable declaration:

```go
// Explicit
var name string = "Kavishka"

// Implicit (short declaration)
name2 := "Sasindu"
```

> **Rule of thumb:** Use implicit (`:=`) when declaring and assigning at the same time. Use explicit (`var`) when you need to declare first and assign later.

---

### 4. 🔄 Type Casting

Convert between types using explicit casting:

```go
number := 3.2
integer_number := int(number) // → 3
```

---

### 5. 🖥️ Console Output

| Verb | Description |
|------|-------------|
| `%v` | Default format (value) |
| `%T` | Type of the variable |
| `%s` | String format |
| `%b` | Binary format |
| `%c` | Character format |

```go
fmt.Println("Hello! World")           // prints with newline
fmt.Printf("My age is %v", age)       // prints with format
fmt.Printf("Age type is %T\n", age)   // prints type
```

---

### 6. 🔀 Conditions & Conditionals

**if / else if / else:**
```go
if number < 10 {
    // ...
} else if number > 10 {
    // ...
} else {
    // ...
}
```

> Note: Go does **not** use parentheses `()` around conditions.

**switch statement:**
```go
switch number {
case 1:
    // ...
case 2:
    // ...
default:
    // ...
}
```

---

### 7. 🔁 Loops

Go only has one loop keyword: `for`.

**Standard for loop:**
```go
for i := 1; i < 10; i++ {
    fmt.Printf("i is %v\n", i)
}
```

**While-style loop:**
```go
idx := 10
for idx < 15 {
    idx++
}
```

**Infinite loop:**
```go
for true {
    // runs forever
}
```

---

### 8. 🔤 Strings

- Strings in Go are stored as `uint8` (1 byte per character).
- Use `%c` to print a character, `%v` to print its integer value.
- Use `range` to iterate over characters:

```go
str := "hello world"
for _, char := range str {
    fmt.Printf("%c\n", char)
}
```

> The `_` (blank identifier) is used to ignore the index when it's not needed.

---

### 9. 📦 Arrays

Go arrays have a **fixed size** defined at declaration.

```go
// Method 1 — explicit
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

// Method 2 — implicit
numbers2 := [5]int{1, 2, 3, 4, 5}

// Method 3 — size inferred from values
numbers3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

**Iterating over an array:**
```go
for i, number := range numbers3 {
    fmt.Printf("Index: %v, Number: %v\n", i, number)
}
```

---

## ▶️ How to Run

Make sure you have [Go installed](https://go.dev/dl/), then run:

```bash
go run main.go
```

---

## 🛠️ Prerequisites

- Go 1.18 or higher

---

## 👤 Author

**Kavishka Sasindu Mudithananda**
