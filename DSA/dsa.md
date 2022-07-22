## Guidelines for Asymptotic Analysis
### 1. Loops
the runtime of a loop is the runtime of the statements in it multiplied by the number of iterations.
```go
for i := 1; i <= n; i++ {       // executes n times
	m = m + 2                   // constant time, c
}
```
total time = constant c * n = cn = O(n)

### 2. Nested Loops
the runtime is the product of the sizes of all loops.
```go
for i := 1; i <= n; i++ {
	for j:=1;j<=n;j++{
		k = k+1
    }
}
```
total time = c * n * n = cn^2 = O(n^2)

### 3. Consecutive statements:
add the complexities of each statement
```go
x = x + 1
for i := 1; i <= n; i ++ {
	m = m + 2
}
for i := 1; i <= n; i ++ {
	for j := 1; j <=n; j ++ {
		k = k + 1
    }
}
```
total time = c + cn + cn^2 = O(n^2)

### 4. If-Else Statements:
the worst case runtime = the test + then or the test + else (whichever is larger)
```go
if length() == 0 {
	return false
} else {
	for j := 1; j <= len(a); j ++ {
		k = k + 1
}
}
```




## Recursion and Backtracking
A function which calls itself is recursive. A recursive method solves a problem by calling a copy of itself to solve a smaller problem.
Recursion is useful in sort, search, traversal... any task that can be split into similar subtasks.

