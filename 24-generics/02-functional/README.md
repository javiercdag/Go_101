# Generic functions

In this exercise you  will have to implement a generic *reduce* function.
The reduce function receives as arguments a slice of type E, a variable of type E and a function of type (R, E) -> R,
and returns a single value of type R,where E and R could be any type.

Reduce is a function that receives an operation and a list and returns a value which is the accumulation of the operation performed on each list element.
For example:
```python
    multiply := accumulator, b -> accumulator * b
    values := {1,2,3,4}
    reduced := reduce(values,1,multiply) // produces 24
    // 1 is the initial value of the accumulator
```


Place your code into the file exercise.go near the placeholder // INSERT YOUR CODE HERE.