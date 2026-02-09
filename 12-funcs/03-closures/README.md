# Closures as proxy

In this exercise, you will create a **limiter proxy** function.
The function should be called *proxy* and should satisfy the following criteria:
- it should receive and int and a function that has no parameters and returns an int
- it should return a function that has no parameters and returns an int and an error
- if the returned function was called less times than the initial limit value it propagates the call to the received function and returns its value and nil, otherwise it should return 0 and an error


It is important that you do not externalize the state of the function, instead use closures to accomplish a similar effect.
Just for clarity, here is an example usage of the proxy function.
```python
  limited := proxy(2,() -> 2)
  limited() // returns (2, nil)
  limited() // returns (2, nil)
  limited() // returns (0, Error)
```


Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.