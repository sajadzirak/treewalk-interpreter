# TreeWalk Interpreter

A tree-walking interpreter for a dynamically-typed programming language, implemented in Go. This project was developed following Thorsten Ball's book "Writing An Interpreter In Go" as a learning exercise in compiler design and programming language implementation.

## Overview

This interpreter implements a functional programming language featuring first-class functions, closures, and a clean expression-oriented syntax. The interpreter implements a complete pipeline from lexical analysis through evaluation, demonstrating fundamental concepts in language design and implementation.

### Pipeline

```
┌─────────────┐    ┌───────┐    ┌────────┐    ┌────────┐    ┌───────────┐    ┌────────┐
│ Source Code │ -> │ Lexer │ -> │ Tokens │ -> │ Parser │ -> │    AST    │ -> │  Eval  │ -> Result
└─────────────┘    └───────┘    └────────┘    └────────┘    └───────────┘    └────────┘
```

## Features

- **Integers and arithmetic operations**: `+`, `-`, `*`, `/`
- **Boolean values and comparison operators**: `!`, `==`, `!=`, `<`, `>`
- **Variable bindings**: `let x = 5;`
- **First-class functions**: Functions as values with lexical scoping
- **Closures**: Functions can capture variables from their defining scope
- **Conditionals**: `if`/`else` expressions
- **Return statements**: Early return from functions
- **REPL**: Interactive Read-Eval-Print Loop for experimentation

## Example Code

```
let add = fn(x, y) { x + y };
add(5, 10);  // returns 15

let fibonacci = fn(n) {
  if (n < 2) {
    n
  } else {
    fibonacci(n - 1) + fibonacci(n - 2)
  }
};
fibonacci(10);  // returns 55

let makeAdder = fn(x) {
  fn(y) { x + y }
};
let addFive = makeAdder(5);
addFive(10);  // returns 15 (closure captures x=5)
```

## Architecture

The interpreter follows a traditional architecture with clearly separated stages:

1. **Lexer** - Tokenizes source code into a stream of tokens
2. **Parser** - Builds an Abstract Syntax Tree (AST) using a Pratt parser
3. **Evaluator** - Walks the AST and evaluates expressions

### Project Structure

```
treewalk-interpreter/
├── main.go           # Entry point, starts the REPL
├── token/            # Token type definitions
├── lexer/            # Lexical analysis
├── ast/              # Abstract Syntax Tree node definitions
├── parser/           # Pratt parser implementation
├── object/           # Runtime value representations and environment
├── evaluator/        # Tree-walking interpreter
└── repl/             # Read-Eval-Print Loop
```

## Installation and Usage

### Prerequisites

- Go 1.23.1 or higher

### Running the REPL

```bash
go run main.go
```

This will start an interactive session:

```
This is the TreeWalk Interpreter.
>> let x = 5;
5
>> x + 10
15
>>
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./lexer
go test ./parser
go test ./evaluator
```

## Implementation Details

### Lexer
The lexer performs tokenization by scanning input character by character, recognizing keywords, identifiers, integers, and operators. It handles two-character operators like `==` and `!=` through lookahead.

### Parser
A Pratt parser (top-down operator precedence parser) constructs the AST. It handles operator precedence naturally through the parsing algorithm, supporting both prefix and infix expressions. The parser includes comprehensive error reporting.

### Evaluator
The evaluator is a tree-walking interpreter that recursively evaluates AST nodes. It uses an environment structure for variable binding, with support for nested scopes to implement closures correctly.

### Object System
Runtime values are represented as Go interfaces with concrete types for integers, booleans, functions, return values, null, and errors. Functions carry their defining environment to support closures.

## Future Enhancements

Potential areas for expansion:
- Additional numeric types (floats, larger integers)
- String type and operations
- Arrays and hash maps
- Additional operators (modulo, bitwise operations)
- Built-in standard library functions

## References

- **Book**: "Writing An Interpreter In Go" by Thorsten Ball
- **Language**: Go programming language

## License

This is an educational project created for learning purposes.
