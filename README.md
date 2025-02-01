# Reverse Polish Notation Calculator

This is a simple calculator that evaluates Reverse Polish Notation (RPN) expressions.
The goal was to replicate the behavior of a RPN calculator
using the Stack data structure.

This project is made for learning purposes only. Both for learning Golang and improve my knowledge about the Stack data structure.

## Usage

Clone the repository:
```
$ git clone https://github.com/eoBattisti/proj-stack-machine/
```

Into the project directory, run the following command:
```bash
$ go build -o rpn
```

Or

```bash
go run main.go "<expression>"
```

## Examples

### Basic operations
This calculator supports the following operations:

1. Addition: `+`
```bash
$ ./rpn "1 2 +"
3
```

2. Subtraction: `-`
```bash
$ ./rpn "3 2 -"
1
```

3. Multiplication: `*`
```bash
$ ./rpn "2 3 *"
6
```

4. Division: `/`
```bash
$ ./rpn "6 3 /"
2
```

### Chain operations

It's possible to chain operations in a single expression:

```bash
$ ./rpn "2 3 * 70 +"
142
```
