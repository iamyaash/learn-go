# Formatting
Formatting issues in `go` are the most contentious and the least consequential thing to worry about. Go can take care of the styling, as it uses `gofmt` program in a package level rather than source file level. 

The `go fmt` reads a `go` program and emits the source in a standard style of **indentation & vertical alignment**.

#### Formatting Details
**Indentation**:
Tabs for indentation and `gofmt` emits them by default.
> **Note**: Use space only if you must!

**Line Length**:
Go has no line length limit. No need to worry about overflowing a line and if a line feels too long, wrap it and indent with extra tabs.
```go
func main() {
	// A long line of code can be wrapped and indented for better readability
	result := calculateSum(
		100, 200, 300, 400, 500, 
		600, 700, 800, 900, 1000,
	)

	fmt.Println("The result is:", result)
}
```

**Parentheses**:
Go needs fewer parentheses than C and Java. The control structures(`if, for, switch`) do not have parentheses in their syntax.
```go
func main() {
    x := 10

    if x%2 == 0 { // No parentheses around the condition
        fmt.Println("x is even")
    } else {
        fmt.Println("x is odd")
    }
}
```

# Commentary
Go provides
- `C`-Style `/*  */` block comments, and
- `C++` -Style `//` line comments.

Line comments are common, but the block comments are mostly for package comments(also useful within an expression or to disable large swaths of code).

Comments that appear before top-level declarations, with no intervening newlines, are considered to document the declaration itself. These “_doc comments_d” are the primary documentation for a given Go package or command.