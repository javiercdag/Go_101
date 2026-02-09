# Search

Write a function that reads the below text and returns `true` or `false` if "ONE" appears in the text *more than once*, without respect to letter case (i.e., the each of words `Hello`, `heLLo`, and `HELlo` match for the search strings `HELLO` and `hello`).

> "One by one he draws us out of the world. The Empire never ended. The Head Apollo is about to return. St. Sophia is going to be born again; she was not acceptable before. The Buddha is in the park. Siddhartha sleeps (but is going to awaken). The time you have waited for has come."

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

HINT:

- Make sure to remove punctuation from the input (like `.`, `;`, or `?`) and trim the text to alphanumeric characters. For instance, you can use [`regexp.ReplaceAllString`](https://pkg.go.dev/regexp#Regexp.ReplaceAllString) with the regular expression `[^a-zA-Z0-9 ]+` for this purpose.
- You can use a `map[string]int` to remember the number of times a particular word is found while iterating through the text word by word.
