# Testing

In this task, you need to test the function `strncat` in Go using **`testify/assert`**.

## Function Description

The implementation is available in [exercise_test.go](exercise_test.go), if needed.
Below we show some examples of the expected behaviors.

### `strncat(dest *string, src string, n int)`

- Appends at most `n` characters from `src` to `dest`.
- Modifies `dest` in place

**Example:**

```go
dest := "Hello"
strncat(&dest, "World", 3) // dest == "HelloWor"

dest := "Hi"
strncat(&dest, "All", 10)  // dest == "HiAll"

dest := ""
strncat(&dest, "Go", 1)    // dest == "G"
```

## Requirements
1. Use only assert statements from the [testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert) package.
2. Do not modify the provided function signature.
3. Write exactly three test functions for the function `strncat`.

## Needed test functions:

- `StrNCatInbounds(t *testing.T)` — append a part of the source within index bounds
- `StrNCatOutOfBounds(t *testing.T)` — append more than the source length
- `StrNCatEmptySource(t *testing.T)` — append from an empty source

Insert the code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.
